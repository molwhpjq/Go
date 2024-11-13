package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"testbeegodemo/models"
)

const (
	USERNAME = "login"
	LOGIN    = "loginname"
)

type UserController struct {
	beego.Controller
}
type LoginRequest struct {
	Uname string `json:"uname"`

	Upwd string `json:"upwd"`
}
type LoginResponse struct {
	Success bool `json:"success"`

	Message string `json:"message"`
}

// 设置全局变量
var username string

func (c *UserController) Get() {

	c.TplName = "login.html"
}

func (c *UserController) Post() {
	var loginRequest LoginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginRequest)
	if err != nil {
		c.Data["json"] = &LoginResponse{Success: false, Message: "解析请求失败"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	var user models.Userinfo
	username = loginRequest.Uname
	err = o.QueryTable("userinfo").Filter("Uname", loginRequest.Uname).One(&user)
	if err != nil {
		c.Data["json"] = &LoginResponse{Success: false, Message: "用户名不存在"}
		c.ServeJSON()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Upwd), []byte(loginRequest.Upwd))
	if err != nil {
		c.Data["json"] = &LoginResponse{Success: false, Message: "密码错误"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &LoginResponse{Success: true, Message: "登录成功"}
	c.ServeJSON()
}

/*func (c *UserController) Profile() {
	c.TplName = "profile.html"
}*/

/*func (c *UserController) MyBooks() {
	o := orm.NewOrm()
	var books []*models.Bookinfo
	uname := c.GetSession("login")
	_, err := o.QueryTable("book").Filter("publisher", uname.(string)).All(&books)
	if err != nil {
		beego.Error(err)
	}
	c.Data["books"] = books
	c.TplName = "mybooks.html"
}

func (c *UserController) Logout() {
	c.DelSession("login")
	c.Redirect("/", 302)
}
*/
