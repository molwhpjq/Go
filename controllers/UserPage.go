package controllers

import "github.com/astaxie/beego"

type UserPageController struct {
	beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *UserPageController) Get() {
	c.TplName = "UserPage.html"
}
