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

var (
	sigs = make(chan os.Signal, 1)
	done = make(chan bool, 1)
)

func init() {
	config.Init()

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.EnableErrorsRender = false
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	go beego.Run()

	service.RecordChecker.SetTrueWithoutCheck(config.RecordWithOutCheck).SetCheckFace(config.RecordCheckHasFace).SetCheckMotion(config.RecordCheckHasMotion).SetDrawMark(config.CheckerDrawMark)

	service.Mp4Recorder.SetRecordSavePath(config.RecordSavePath).SetRecordSaveName(config.RecordSaveName).SetCheckIntervalSeconds(config.RecordStartCheckIntervalSeconds).SetStopByCheckFailedMaxTimes(config.RecordStopByCheckFailMaxTimes).SetWriteNewFileIntervalMinutes(config.RecordWriteNewFileIntervalMinutes)

	var err error
	for {
		err = cameraServiceRun()
		if nil == err {
			return
		}
		fmt.Println(err)
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func cameraServiceRun() (err error) {
	defer cameraServiceSafeStop()

	//fmt.Println(config.CameraDeviceID)
	err = service.MjpegWebcam.OpenCamera(config.CameraDeviceID)
	if err != nil {
		return err
	}

	service.MjpegWebcam.SetFrameWidth(config.CameraFrameWidth)
	service.MjpegWebcam.SetFrameHeight(config.CameraFrameHeight)
	service.MjpegWebcam.SetFPS(config.CameraFPS)
	fmt.Println("camera set ", service.MjpegWebcam.CodecString(), service.MjpegWebcam.GetFrameWidth(), service.MjpegWebcam.GetFrameHeight(), service.MjpegWebcam.GetFrameFPS())

	service.Mp4Recorder.SetFrameFromWebcam(service.MjpegWebcam)

	img := gocv.NewMat()
	defer img.Close()

	colorGray := color.RGBA{200, 200, 200, 0}

	cameraReadFailedTimer := time.NewTimer(time.Duration(10) * time.Second)

	for {

		select {
		case <-done:
			fmt.Println("done")
			cameraServiceSafeStop()
			return nil
		default:
			break

		}

		if ok := service.MjpegWebcam.Read(&img); !ok {
			select {
			case <-cameraReadFailedTimer.C:
				cameraReadFailedTimer.Stop()
				cameraReadFailedTimer = nil
				return fmt.Errorf("camera is closed")
			default:
				break

			}

			//fmt.Printf("Device closed: %v\n", config.CameraDeviceID)
			continue
		}
		cameraReadFailedTimer.Reset(time.Duration(10) * time.Second)
		if img.Empty() {
			//fmt.Println(fmt.Sprintf("img empty %s", time.Now()))
			continue
		}
		//i++
		gocv.PutText(&img, time.Now().Format("2006-01-02 15:04:05"), image.Pt(10, 20), gocv.FontHersheyPlain, 1.2, colorGray, 2)

		//fmt.Println(fmt.Sprintf("main send img  %d %d %d", img.Cols(), img.Rows(), img.Size()))
		err, img = service.Mp4Recorder.Send(img)
		service.MjpegService.Send(img)

	}
	fmt.Println("exited", time.Now().Format("15:04:05"))
	return nil

}

func cameraServiceSafeStop() {
	service.MjpegWebcam.Close()
	service.RecordChecker.Close()
	service.Mp4Recorder.StopRecord()
	fmt.Println("safe exited")
}
