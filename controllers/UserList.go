package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/models"
)

type UserlistController struct {
	beego.Controller
}

func IsUserLoggedIn(c *beego.Controller) bool {
	// 获取 session 中的用户名
	username := c.GetSession(USERNAME)
	// 判断用户名是否存在
	if username == nil {
		return false

	}
	return true
}
func (c *UserlistController) Get() {
	//if !IsUserLoggedIn(&c.Controller) {
	//	c.Redirect("/login", 302)
	//	return
	//}

	c.TplName = "userlist.html"
}
func (c *UserlistController) GetAll() {
	var users []models.Userinfo
	o := orm.NewOrm()
	o.QueryTable("userinfo").All(&users)
	if len(users) > 0 {
		var repList []interface{}
		for _, user := range users {
			repList = append(repList, user.UserToRespDesc())
			fmt.Println(user.Id, user.Uname)

		}
		c.Data["json"] = repList
	}
	c.ServeJSON()
}

func (c *UserlistController) FindById() {

	id, _ := c.GetInt(":id")
	var user models.Userinfo
	o := orm.NewOrm()
	//err := o.QueryTable("userinfo").Filter("id", id).One(&user)
	error := o.QueryTable("userinfo").Filter("id", id).One(&user)
	if error != nil {
		fmt.Println(error)
		// 查询出错
		c.Data["json"] = map[string]interface{}{"error": "User not found"}
		c.ServeJSON()
		return
	}

	// 将用户信息以JSON格式返回给前端
	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserlistController) DeleteUser() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"}
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	_, err = o.Delete(&models.Userinfo{Id: id})
	if err != nil {
		c.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	c.ServeJSON()

}
