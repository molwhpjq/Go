package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *HomeController) Get() {

	c.TplName = "HomeBlock.html"
}
