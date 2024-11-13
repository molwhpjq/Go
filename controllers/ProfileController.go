package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/models"
)

type ProfileController struct {
	beego.Controller
}

// Get方法是处理HTTP GET请求的处理程序
func (c *ProfileController) Get() {
	o := orm.NewOrm()

	username := username // 设置默认用户名以进行测试
	var Profile []models.Profile
	_, err := o.QueryTable("Profile").Filter("Name", username).All(&Profile)
	if err != nil {
		fmt.Println("查找失败")
		return
	}
	c.Data["Profiles"] = Profile
	c.Data["username"] = username
	c.TplName = "profile.html"
}

func (c *ProfileController) Post() {
	uname := c.GetString("username")
	bio := c.GetString("bio")

	o := orm.NewOrm()
	profile := models.Profile{Name: uname}

	err := o.Read(&profile, "Name")
	if err != nil {
		fmt.Println("查找失败")
		return
	}

	profile.Introduction = bio
	_, err = o.Update(&profile)
	if err != nil {
		fmt.Println("更新失败")
		return
	}

	c.Redirect("/profile", 302)
}
