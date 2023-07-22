package service

import (
	"fmt"
	"gitee.com/hilaoyu/go-basic-utils/utilFile"
	"gitee.com/hilaoyu/go-basic-utils/utilStr"
	"gocv.io/x/gocv"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Recorder struct {
	writer      *gocv.VideoWriter
	frameHeight int
	frameWidth  int
	fps         float64

	recordSaveName string
	recordSavePath string

	checkIntervalSeconds        int
	writeNewFileIntervalMinutes int64
	writerStartTime             time.Time
	checkTimer                  *time.Timer
	checkFailedTimes            int
	stopByCheckFailedMaxTimes   int
}

type RecordFileInfo struct {
	os.FileInfo
	SubPath       string
	ModTimeFormat string
	SizeFormat    string
}

type RecordFileByModTimeDesc []RecordFileInfo

func (fis RecordFileByModTimeDesc) Len() int {
	return len(fis)
}

func (fis RecordFileByModTimeDesc) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis RecordFileByModTimeDesc) Less(i, j int) bool {
	return fis[i].ModTime().After(fis[j].ModTime())
}

var (
	Mp4Recorder = NewRecorder()
)

func NewRecorder() *Recorder {
	return &Recorder{
		frameHeight: 480,
		frameWidth:  600,
		fps:         20,
	}
}

func (r *Recorder) SetFrameFromWebcam(c *Webcam) *Recorder {
	r.SetFrameWidth(int(c.GetFrameWidth())).SetFrameHeight(int(c.GetFrameHeight())).SetFPS(c.GetFrameFPS())
	return r
}

func (r *Recorder) SetFrameWidth(v int) *Recorder {
	r.frameWidth = v
	return r
}
func (r *Recorder) SetFrameHeight(v int) *Recorder {
	r.frameHeight = v
	return r
}

func (r *Recorder) SetFPS(v float64) *Recorder {
	r.fps = v
	return r
}

func (r *Recorder) SetRecordSavePath(v string) *Recorder {
	r.recordSavePath = v
	return r
}

func (r *Recorder) SetRecordSaveName(v string) *Recorder {
	r.recordSaveName = v
	return r
}
func (r *Recorder) SetCheckIntervalSeconds(v int) *Recorder {
	r.checkIntervalSeconds = v
	r.checkTimer = time.NewTimer((time.Duration(r.checkIntervalSeconds) * time.Second))
	return r
}
func (r *Recorder) SetStopByCheckFailedMaxTimes(v int) *Recorder {
	r.stopByCheckFailedMaxTimes = v
	return r
}
func (r *Recorder) SetWriteNewFileIntervalMinutes(v int64) *Recorder {
	r.writeNewFileIntervalMinutes = v
	return r
}

func (r *Recorder) GetRecordSavePath() string {
	return r.recordSavePath
}

func (r *Recorder) IsCanRecord() bool {
	return "" != r.recordSavePath && utilFile.CheckDir(r.recordSavePath)
}

func (r *Recorder) StartRecord() (err error) {
	if !r.IsCanRecord() {
		r.StopRecord()
		return
	}
	if nil != r.writer && ((time.Now().Unix() - r.writerStartTime.Unix()) > (r.writeNewFileIntervalMinutes * 60)) {
		r.StopRecord()
	}
	if nil == r.writer {
		saveFileName := r.recordSaveName + "-" + time.Now().Format("20060102-150405") + ".mp4"
		saveFile := filepath.Join(r.recordSavePath, saveFileName)
		fmt.Println("recording ", saveFile)
		r.writer, err = gocv.VideoWriterFile(saveFile, "mp4v", r.fps, r.frameWidth, r.frameHeight, true)
		if err != nil {
			err = fmt.Errorf("error opening video writer device: %v\n", saveFile)
		}
		r.writerStartTime = time.Now()
	}

	return
}

func (r *Recorder) WriteRecord(img gocv.Mat) (err error) {
	if nil == r.writer {
		fmt.Println("writer is nil")
		return
	}
	fmt.Println(fmt.Sprintf("record img %d x%d", img.Cols(), img.Rows()))
	return r.writer.Write(img)
}

func (r *Recorder) Send(img gocv.Mat) (err error) {
	select {
	case <-r.checkTimer.C:
		if RecordChecker.Check(img) {
			r.StartRecord()
			r.checkFailedTimes = 0
			//fmt.Println("checker true")
		} else {
			//fmt.Println("checker false")
			r.checkFailedTimes++

			if r.checkFailedTimes >= r.stopByCheckFailedMaxTimes {
				r.StopRecord()
			}
		}
		r.checkTimer.Reset(time.Duration(r.checkIntervalSeconds) * time.Second)
		break
	default:
		break

	}
	if nil == r.writer {
		//fmt.Println("writer is nil")
		return
	}
	//fmt.Println(fmt.Sprintf("record img %d x%d", img.Cols(), img.Rows()))
	return r.writer.Write(img)
}

func (r *Recorder) StopRecord() {
	if nil != r.writer {
		r.writer.Close()
		r.writer = nil
		fmt.Println("record writer closed")
	}
}

func (r *Recorder) ListRecords() (recordFiles RecordFileByModTimeDesc, err error) {
	recordSavePath := r.GetRecordSavePath()

	var files []RecordFileInfo
	err = filepath.Walk(recordSavePath, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		if info.IsDir() {
			return nil
		}
		info.Sys()
		files = append(files, RecordFileInfo{
			FileInfo:      info,
			SubPath:       strings.TrimLeft(utilStr.AfterLast(path, recordSavePath), "/"),
			ModTimeFormat: info.ModTime().Format("2006-01-02 15:04:05"),
			SizeFormat:    utilFile.FormatSize(info.Size()),
		})
		return nil
	})
	recordFiles = RecordFileByModTimeDesc(files)
	sort.Sort(recordFiles)
	return
}
