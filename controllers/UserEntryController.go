package controllers

import "github.com/astaxie/beego"

type UserEntryController struct {
	beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *UserEntryController) Get() {
	c.TplName = "UserEntry.html"
}
