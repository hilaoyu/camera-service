package service

import (
	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
	"net/http"
)

type MjpegStreamService struct {
	stream *mjpeg.Stream
}

var (
	MjpegService = NewMjpegStreamService()
)

func NewMjpegStreamService() (server *MjpegStreamService) {

	//fmt.Println("mjpeg listen", addr, path)
	server = &MjpegStreamService{
		stream: mjpeg.NewStream(),
	}

	return
}

func (s *MjpegStreamService) Send(img gocv.Mat) {
	//fmt.Println(fmt.Sprintf("mjpeg send img %d x%d", img.Cols(), img.Rows()))
	mjpegBuf, _ := gocv.IMEncode(".jpg", img)
	defer mjpegBuf.Close()
	s.stream.UpdateJPEG(mjpegBuf.GetBytes())

}

func (s *MjpegStreamService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.stream.ServeHTTP(w, r)
}
