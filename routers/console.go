package routers

import (
	"camera-service/controllers"
	context2 "context"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/url"
)

var filterCheckCruSession beego.FilterFunc = func(ctx *context.Context) {

	//return
	//session
	if logged, ok := ctx.Input.CruSession.Get(context2.Background(), "isAdminLogged").(bool); logged && ok {
		return
	}

	flashMessage := "\x00" + "error" + "\x23" + beego.BConfig.WebConfig.FlashSeparator + "\x23" + "未登录或登录已过期" + "\x00"
	ctx.SetCookie(beego.BConfig.WebConfig.FlashName, url.QueryEscape(flashMessage), 0, "/")
	ctx.Redirect(302, "/console")
	return
}

func init() {
	go beego.SetStaticPath("/static", "static")
	beego.Router("/console", &controllers.WebConsoleController{}, "get:ShowLoginForm;post:SessionLogin")
	beego.Router("/console/logout", &controllers.WebConsoleController{}, "*:SessionLogout")

	beego.InsertFilter("/console/:path(.+)", beego.BeforeExec, filterCheckCruSession)
}
