package controllers

import (
	"camera-service/config"
)

type WebConsoleController struct {
	BaseWebController
}

func (c *WebConsoleController) ShowLoginForm() {

	if message, ok := c.GetSession("error").(string); ok && "" != message {
		c.Data["Error"] = message
	}

	c.DelSession("error")

	//c.TplName = "parse/console.tpl"
	//return
	c.TplName = "login.tpl"
	logged := c.GetSession("isAdminLogged")
	if nil != logged {
		if loggedBool, ok := logged.(bool); ok && loggedBool {
			c.TplName = "console.tpl"
		}
	}

}

func (c *WebConsoleController) SessionLogin() {

	pass := c.GetString("password", "")
	if "" != pass && config.CheckAdminPassword(pass) {
		c.SetSession("isAdminLogged", true)

	} else {
		c.SetSession("error", "密码错误")
	}
	c.Redirect("/console", 302)
}
func (c *WebConsoleController) SessionLogout() {

	c.DelSession("isAdminLogged")
	c.Redirect("/console", 302)
}
