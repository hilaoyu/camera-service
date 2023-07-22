package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	//beegoUtils "github.com/beego/wetalk/modules/utils"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) ParseInput(v interface{}) (err error) {
	if "application/json" == c.Ctx.Request.Header.Get("Content-Type") && len(c.Ctx.Input.RequestBody) > 0 {
		err = json.Unmarshal(c.Ctx.Input.RequestBody, v)
		//fmt.Println("json", err)
	} else {
		err = c.ParseForm(v)
		//fmt.Println("form", err)
	}
	return
}
func (c *BaseController) SetMessage(t string, v string) {

	temp, ok := c.Data["pageMessage"]
	if !ok {
		return
	}
	message, ok := temp.(map[string]string)
	if !ok {
		return
	}
	message[t] = v
	c.Data["pageMessage"] = message

}
func (c *BaseController) WithFlashError(mag string) {

	flash := beego.NewFlash()
	flash.Error(mag)
	flash.Store(&c.Controller)

}

func (c *BaseController) WithFlashSuccess(mag string) {

	flash := beego.NewFlash()
	flash.Success(mag)
	flash.Store(&c.Controller)

}
func (c *BaseController) WithFlashWarning(mag string) {

	flash := beego.NewFlash()
	flash.Warning(mag)
	flash.Store(&c.Controller)

}
func (c *BaseController) WithFlashNotice(mag string) {

	flash := beego.NewFlash()
	flash.Notice(mag)
	flash.Store(&c.Controller)

}

/*func (c *BaseController) GetPager(defaultLimit ...int) (p *beegoUtils.Paginator) {
	limit, err := c.GetInt("limit", defaultLimit...)
	if nil != err {
		limit = 10
	}
	if limit > 1000 {
		limit = 1000
	}
	p = beegoUtils.NewPaginator(c.Ctx.Request, limit, 0)
	return
}*/
