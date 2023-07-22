package routers

import (
	"camera-service/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/console/camera/stream", &controllers.WebCameraMjpegController{}, "get:MjpegStream")
	beego.Router("/console/camera/play", &controllers.WebCameraMjpegController{}, "get:MjpegStreamPlay")
}
