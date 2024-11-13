package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"testbeegodemo/models"
)

// AdminEntryController 结构体继承自beego.Controller，用于处理管理员入口的HTTP请求
type AdminEntryController struct {
	beego.Controller
}

// 全局变量result，用于存储生成的验证码
var result string

// Get方法处理HTTP GET请求，用于生成验证码

func (c *AdminEntryController) Get() {
	result = "" // 初始化result为空字符串

	// 定义包含数字和字母的字符集
	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 使用随机数生成器生成4位验证码
	for i := 0; i < 4; i++ {
		// 生成随机索引
		index := rand.Intn(len(charSet))
		// 根据索引获取字符并拼接到结果中
		result += string(charSet[index])
	}

	// 打印生成的验证码
	fmt.Println("==========================")
	c.Data["result"] = result // 将验证码存储到beego的上下文Data中，供后续使用
	fmt.Println(result)
	c.TplName = "AdminEntry.html" // 设置要渲染的模板文件名
}

// Post方法处理HTTP POST请求，用于验证用户输入的验证码是否正确，并根据验证结果进行相应的操作

func (c *AdminController) Post() {
	captcha := c.GetString("captcha") // 获取用户输入的验证码
	fmt.Println(captcha)
	fmt.Println("-----------------")
	fmt.Println(result) // 打印存储的验证码

	// 判断用户输入的验证码是否与存储的验证码一致，如果不一致，重定向到AdminEntry页面
	if captcha != result {
		fmt.Println("验证码不对................")
		c.Redirect("/AdminEntry", 302) // 使用302状态码进行临时重定向
		return
	}

	// 其他逻辑代码，包括获取用户名和密码，查询数据库等操作...
	username := c.GetString("username")   // 获取用户名
	password := c.GetString("password")   // 获取密码
	if username == "" || password == "" { // 如果用户名或密码为空，重定向到admin页面，并打印提示信息
		fmt.Println("有数据为空")
		c.Redirect("/admin", 302) // 使用302状态码进行临时重定向
		return
	}
	o := orm.NewOrm()      // 初始化ORM对象，用于数据库操作
	user := models.Admin{} // 定义Admin结构体变量，用于存储查询结果

	o.QueryTable("Admin").Filter("Aname", username).All(&user) // 根据用户名查询数据库中的Admin表，并将结果存储到user变量中
	if username == user.Aname && user.Apwd == password {       // 判断查询结果中的用户名和密码是否与输入一致，如果一致，重定向到admin页面；否则重定向回AdminEntry页面
		c.Redirect("/admin", 302) // 使用302状态码进行临时重定向
	} else {
		c.Redirect("/AdminEntry", 302) // 使用302状态码进行临时重定向
	}
}
