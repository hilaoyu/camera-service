package controllers

import "camera-service/service"

type WebCameraMjpegController struct {
	BaseWebController
}

func (c *WebCameraMjpegController) MjpegStream() {
	service.MjpegService.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *WebCameraMjpegController) MjpegStreamPlay() {
	c.Data["frameWidth"] = service.MjpegWebcam.GetFrameWidth()
	c.Data["frameHeight"] = service.MjpegWebcam.GetFrameHeight()
	c.TplName = "camera_stream_play.tpl"
}
