package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseWebController struct {
	BaseController
}

func (c *BaseWebController) Prepare() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["pageMessage"] = flash.Data
}
