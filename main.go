package main

//  CGO_ENABLED=1;GOOS=linux;GOARCH=amd64
import (
	"camera-service/config"
	"camera-service/controllers"
	_ "camera-service/routers"
	"camera-service/service"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.EnableErrorsRender = false
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	defer safeExit()
	//fmt.Println(config.CameraDeviceID)
	var err error
	err = service.MjpegWebcam.OpenCamera(config.CameraDeviceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	go beego.Run()

	service.MjpegWebcam.SetFrameWidth(config.CameraFrameWidth)
	service.MjpegWebcam.SetFrameHeight(config.CameraFrameHeight)
	service.MjpegWebcam.SetFPS(config.CameraFPS)

	fmt.Println("camera set ", service.MjpegWebcam.CodecString(), service.MjpegWebcam.GetFrameWidth(), service.MjpegWebcam.GetFrameHeight(), service.MjpegWebcam.GetFrameFPS())
	if !service.MjpegWebcam.IsAvailable() {
		fmt.Printf("Cannot read device %v\n", config.CameraDeviceID)
		return
	}

	service.RecordChecker.SetTrueWithoutCheck(config.RecordWithOutCheck).SetCheckFace(config.RecordCheckHasFace).SetCheckMotion(config.RecordCheckHasMotion)

	service.Mp4Recorder.SetFrameFromWebcam(service.MjpegWebcam)
	service.Mp4Recorder.SetRecordSavePath(config.RecordSavePath).SetRecordSaveName(config.RecordSaveName).SetCheckIntervalSeconds(config.RecordStartCheckIntervalSeconds).SetStopByCheckFailedMaxTimes(config.RecordStopByCheckFailMaxTimes).SetWriteNewFileIntervalMinutes(config.RecordWriteNewFileIntervalMinutes)

	img := gocv.NewMat()
	defer img.Close()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	colorGray := color.RGBA{200, 200, 200, 0}
	for {

		select {
		case <-done:
			fmt.Println("done")
			safeExit()
			return
		default:
			break

		}

		if ok := service.MjpegWebcam.Read(&img); !ok {
			//fmt.Printf("Device closed: %v\n", config.CameraDeviceID)
			continue
		}
		if img.Empty() {
			//fmt.Println(fmt.Sprintf("img empty %s", time.Now()))
			continue
		}
		//i++
		gocv.PutText(&img, time.Now().Format("2006-01-02 15:04:05"), image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, colorGray, 2)

		//fmt.Println(fmt.Sprintf("main send img  %d %d %d", img.Cols(), img.Rows(), img.Size()))
		service.MjpegService.Send(img)
		err = service.Mp4Recorder.Send(img)
		//fmt.Println(err)

	}
	fmt.Println("exited", time.Now().Format("15:04:05"))
	return

}

func safeExit() {
	service.MjpegWebcam.Close()
	service.RecordChecker.Close()
	service.Mp4Recorder.StopRecord()
	fmt.Println("safe exited")
}