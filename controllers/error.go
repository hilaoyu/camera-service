package controllers

type ErrorController struct {
	BaseWebController
}

func (c *ErrorController) customError(content string) {
	c.Data["content"] = content
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error404() {
	c.customError("page not found")
}
func (c *ErrorController) Error403() {
	c.customError("sign error1")
}

func (c *ErrorController) Error501() {
	c.customError("server error")
}
