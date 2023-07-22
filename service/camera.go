package service

import (
	"fmt"
	"gocv.io/x/gocv"
)

type Webcam struct {
	camera *gocv.VideoCapture
}

var (
	MjpegWebcam = NewMjpegWebcam()
)

func NewMjpegWebcam() (mjpegWebcam *Webcam) {
	mjpegWebcam = &Webcam{}
	return
}

func (c *Webcam) OpenCamera(deviceID interface{}) (err error) {
	camera, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		err = fmt.Errorf("Error opening capture device: %v\n , %+v", deviceID, err)
		return
	}
	camera.Set(gocv.VideoCaptureFOURCC, camera.ToCodec("MJPG"))
	c.camera = camera
	return
}

func (c *Webcam) IsAvailable() bool {
	return nil != c.camera
}

func (c *Webcam) SetFrameWidth(v float64) *Webcam {
	if c.IsAvailable() {
		c.camera.Set(gocv.VideoCaptureFrameWidth, v)
	}

	return c
}
func (c *Webcam) SetFrameHeight(v float64) *Webcam {
	if c.IsAvailable() {
		c.camera.Set(gocv.VideoCaptureFrameHeight, v)
	}
	return c
}
func (c *Webcam) SetFPS(v float64) *Webcam {
	if c.IsAvailable() {
		c.camera.Set(gocv.VideoCaptureFPS, v)
	}
	return c
}
func (c *Webcam) GetFrameWidth() float64 {
	if !c.IsAvailable() {
		return 0
	}
	return c.camera.Get(gocv.VideoCaptureFrameWidth)
}
func (c *Webcam) GetFrameHeight() float64 {
	if !c.IsAvailable() {
		return 0
	}
	return c.camera.Get(gocv.VideoCaptureFrameHeight)
}
func (c *Webcam) GetFrameFPS() float64 {
	if !c.IsAvailable() {
		return 0
	}
	return c.camera.Get(gocv.VideoCaptureFPS)
}
func (c *Webcam) CodecString() string {
	if !c.IsAvailable() {
		return ""
	}
	return c.camera.CodecString()
}

func (c *Webcam) Read(img *gocv.Mat) bool {
	if !c.IsAvailable() {
		return false
	}
	return c.camera.Read(img)
}

func (c *Webcam) Close() {
	if !c.IsAvailable() {
		return
	}
	c.camera.Close()
}
