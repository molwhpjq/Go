package controllers

import "github.com/astaxie/beego"

type SelectPage struct {
	beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *SelectPage) Get() {
	c.TplName = "SelectPage.html"
}
