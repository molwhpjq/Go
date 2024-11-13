package controllers

import (
	"github.com/astaxie/beego"
)

// AdminController 是一个结构体，继承自beego.Controller
type AdminController struct {
	beego.Controller // 继承beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *AdminController) Get() {
	c.TplName = "admin.html" // 设置模板文件名为"admin.html"
}
