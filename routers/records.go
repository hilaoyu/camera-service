package routers

import (
	"camera-service/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//beego.SetStaticPath("/console/camera/records/mp4", config.RecordSavePath)
	beego.Router("/console/camera/records", &controllers.WebCameraRecordsController{}, "get:ListRecords")
	beego.Router("/console/camera/records/*", &controllers.WebCameraRecordsController{}, "get:DownloadFile")

}
