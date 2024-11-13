package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/models"
)

type UserShow struct {
	beego.Controller
}

func (c *UserShow) Get() {
	o := orm.NewOrm()
	var user []models.Userinfo

	_, err := o.QueryTable("Userinfo").Filter("Uname", username).All(&user)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	fmt.Println("--------------")
	fmt.Println(user)
	c.Data["users"] = user
	fmt.Println(username)
	c.TplName = "UserShow.html"
}

// Delete 方法处理用户删除请求
func (c *UserShow) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("id查询失败")
		return
	}

	o := orm.NewOrm()
	// 根据 ID 查询用户信息，并设置到 user 变量中
	user := models.Userinfo{}
	user.Id = id
	err = o.Read(&user)

	if err != nil {
		fmt.Println("查询失败")
		return
	}
	o.Delete(&user)

	c.Redirect("/UserShow", 302)
}

func (c *UserShow) Show() {
	c.TplName = "update.html"

}

// Update 方法处理用户更新请求
func (c *UserShow) Update() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("用户ID获取失败")
		return
	}

	o := orm.NewOrm()
	user := models.Userinfo{}
	user.Id = id
	err = o.Read(&user)
	if err != nil {
		fmt.Println("用户信息查询失败")
		return
	}

	user.Uname = c.GetString("newUname")
	user.Uage, _ = c.GetInt("newage")

	_, err = o.Update(&user)
	if err != nil {
		fmt.Println("用户信息更新失败")
		return
	}

	c.Redirect("/login", 302)
}
