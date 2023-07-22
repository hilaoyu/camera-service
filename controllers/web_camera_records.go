package controllers

import (
	"camera-service/service"
	"gitee.com/hilaoyu/go-basic-utils/utilFile"
	"path/filepath"
)

type WebCameraRecordsController struct {
	BaseWebController
}

func (c *WebCameraRecordsController) ListRecords() {
	files, err := service.Mp4Recorder.ListRecords()
	if nil != err {
		c.Data["Error"] = err.Error()
	}
	c.Data["files"] = files
	c.TplName = "camera_records.tpl"

}
func (c *WebCameraRecordsController) DownloadFile() {
	subPath := c.Ctx.Input.Param(":splat")
	if "" == subPath {
		c.CustomAbort(404, "not found")
	}
	path := filepath.Join(service.Mp4Recorder.GetRecordSavePath(), utilFile.SafePath(subPath))
	c.Ctx.Output.Download(path, filepath.Base(path))
	//c.Ctx.Output.Header("Content-Type", "video/mp4")
}
func (c *WebCameraRecordsController) Play() {

}
