package controllers

import (
	"encoding/json"
	"fmt"
	"testbeegodemo/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"testbeegodemo/models"
)

type RegisterController struct {
	beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

//Register() 方法是注册控制器中的注册方法

func (c *RegisterController) Register() {
	// 定义一个用户信息变量，用于存储从请求中解析的用户信息
	var user models.Userinfo
	// 从请求的 body 中解析 JSON 数据到 user 变量中
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	// 打印解析后的用户信息
	fmt.Println(user)

	// 使用 bcrypt 库生成密码哈希值
	// 生成密码哈希
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Upwd), bcrypt.DefaultCost)
	if err != nil {
		// 如果密码哈希生成失败，返回错误信息，并以 JSON 格式返回给前端
		c.Data["json"] = map[string]string{"error": "Password hashing failed."}
		c.ServeJSON()
		return
	}

	// 将生成的密码哈希值赋给 user 结构体的 Upwd 字段
	user.Upwd = string(passwordHash)

	// 创建一个 ORM 对象，用于数据库操作
	om := orm.NewOrm()

	// 使用 ORM 的 Insert 方法将 user 插入到数据库中
	_, err = om.Insert(&user)

	// 定义一个 API 响应变量，用于存储返回给前端的响应信息
	var response util.APIResponse

	// 如果插入数据库成功，则返回注册成功的响应信息
	if err == nil {
		response = util.JSONResponse(200, "注册成功")
	} else {
		// 如果插入数据库失败，则返回注册失败的响应信息
		response = util.JSONResponse(502, "注册失败")
	}

	// 将 API 响应信息设置到 c.Data 中，并以 JSON 格式返回给前端
	c.Data["json"] = response
	c.ServeJSON()
}
