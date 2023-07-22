package config

import (
	"camera-service/service"
	"gitee.com/hilaoyu/go-basic-utils/utilFile"
	"gitee.com/hilaoyu/go-basic-utils/utils"
	beego "github.com/beego/beego/v2/server/web"
	"path/filepath"
)

var (
	adminPassword string

	CameraDeviceID                    string
	RecordSavePath                    string
	RecordSaveName                    string
	RecordWithOutCheck                bool
	RecordCheckHasFace                bool
	RecordCheckHasMotion              bool
	RecordStartCheckIntervalSeconds   int
	RecordStopByCheckFailMaxTimes     int
	RecordWriteNewFileIntervalMinutes int64

	CameraFrameHeight float64
	CameraFrameWidth  float64
	CameraFPS         float64

	MjpegService *service.MjpegStreamService
)

func Init() {
	adminPassword = beego.AppConfig.DefaultString("adminPassword", "")

	CameraDeviceID = beego.AppConfig.DefaultString("camera::cameraDeviceID", "0")
	RecordSavePath = beego.AppConfig.DefaultString("camera::recordSavePath", "records")
	RecordSaveName = beego.AppConfig.DefaultString("camera::recordSaveName", "camera")
	RecordWithOutCheck = beego.AppConfig.DefaultBool("camera::recordWithOutCheck", false)
	RecordCheckHasFace = beego.AppConfig.DefaultBool("camera::recordCheckHasFace", false)
	RecordCheckHasMotion = beego.AppConfig.DefaultBool("camera::recordCheckHasMotion", false)
	RecordStartCheckIntervalSeconds = beego.AppConfig.DefaultInt("camera::recordStartCheckIntervalSeconds", 3)
	RecordWriteNewFileIntervalMinutes = beego.AppConfig.DefaultInt64("camera::recordWriteNewFileIntervalMinutes", 10)
	RecordStopByCheckFailMaxTimes = beego.AppConfig.DefaultInt("camera::recordStopByCheckFailMaxTimes", 40)

	CameraFrameHeight = beego.AppConfig.DefaultFloat("camera::cameraFrameHeight", 480)
	CameraFrameWidth = beego.AppConfig.DefaultFloat("camera::cameraFrameWidth", 600)
	CameraFPS = beego.AppConfig.DefaultFloat("camera::cameraFPS", 20)

	if "" != RecordSavePath {
		RecordSavePath = utilFile.SafePath(RecordSavePath)
		if !filepath.IsAbs(RecordSavePath) {
			RecordSavePath = filepath.Join(utils.GetSelfPath(), RecordSavePath)
		}
	}
}

func CheckAdminPassword(pass string) bool {
	//fmt.Println(pass,":",apiAdminPassword)
	return "" != pass && "" != adminPassword && pass == adminPassword
}
