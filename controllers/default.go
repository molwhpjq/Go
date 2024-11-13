package controllers

import (
	"github.com/astaxie/beego"
)

// 定义 MainController 结构体，继承自 beego.Controller

type MainController struct {
	beego.Controller
}

// 定义 MainController 的 Get 方法，处理 HTTP GET 请求
func (c *MainController) Get() {
	// 从会话中获取名为 "user" 的数据，如果为空，则重定向到 "/login" 页面，HTTP状态码为302（临时重定向）
	if c.GetSession("user") == nil {
		c.Redirect("/login", 302)
	} else {
		// 如果会话中存在 "user" 数据，则设置布局文件为 "layout.html"
		c.Layout = "layout.html"
		// 渲染控制器视图，可能使用指定的布局文件渲染页面
		c.Render()
	}
}
