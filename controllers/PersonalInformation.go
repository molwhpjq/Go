package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/models"
)

type PersonalInformation struct {
	beego.Controller
}

func (c *PersonalInformation) Get() {
	o := orm.NewOrm()
	var users []models.Userinfo      // 定义一个用户信息切片
	page, _ := c.GetInt("pageIndex") // 从 URL 参数中获取当前页码索引
	pageSize := 6                    // 设置每页显示的用户数量为 6
	offset := (page - 1) * pageSize  // 根据页码索引和每页显示数量计算偏移量
	//o.QueryTable("Bookinfo").All(&recommendedBooks)

	// 使用 ORM 查询数据库，获取指定页码的用户信息，并存储到 users 切片中
	_, _ = o.QueryTable("Userinfo").Limit(pageSize, offset).All(&users)
	fmt.Println("-------------", users)
	// 计算总页数和总记录数
	total, _ := o.QueryTable("Userinfo").Count()
	pageCount := float64(total) / float64(pageSize)

	// 如果总用户数量不能被每页显示数量整除，则总页数加一
	if int(total)%int(pageSize) > 0 {
		pageCount++
	}
	// 判断是否是首页和是否是最后一页，并设置到 c.Data 中供模板渲染使用
	FirstPage := false //标识是否是首页
	EndPage := false
	if page == 1 {
		FirstPage = true

	}
	if page == int(pageCount) {
		EndPage = true

	}

	c.Data["EndPage"] = EndPage
	c.Data["FirstPage"] = FirstPage
	c.Data["pageIndex"] = page
	c.Data["users"] = users
	c.TplName = "PersonalInformation.html"
}

// Delete 方法处理用户删除请求
func (c *PersonalInformation) Delete() {
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

	c.Redirect("/personalInformation", 302)
}

func (c *PersonalInformation) Show() {
	c.TplName = "update.html"

}

// Update 方法处理用户更新请求
func (c *PersonalInformation) Update() {
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

	c.Redirect("/personalInformation", 302)
}

//查找代码

func (c *PersonalInformation) Post() {
	uname := c.GetString("uname")
	fmt.Println(uname)
	o := orm.NewOrm()
	var users []models.Userinfo
	_, err := o.QueryTable("Userinfo").Filter("Uname", uname).All(&users)
	fmt.Println("----------------------------------")
	fmt.Println(users)
	if err != nil {
		fmt.Println("查询失败", err)
		c.Data["users"] = nil
		fmt.Println("=========================")
		fmt.Println(users)

	} else {

		page, _ := c.GetInt("pageIndex")
		pageSize := 2
		offset := (page - 1) * pageSize
		//o.QueryTable("Bookinfo").All(&recommendedBooks)

		_, _ = o.QueryTable("Userinfo").Filter("Uname", uname).Limit(pageSize, offset).All(&users)
		fmt.Println("-------------", users)
		// 计算总页数和总记录数
		total, _ := o.QueryTable("Userinfo").Filter("Uname", uname).Count()
		pageCount := float64(total) / float64(pageSize)

		if int(total)%int(pageSize) > 0 {
			pageCount++
		}
		FirstPage := false //标识是否是首页
		EndPage := false
		if page == 1 {
			FirstPage = true

		}
		if page == int(pageCount) {
			EndPage = true

		}

		c.Data["EndPage"] = EndPage

		c.Data["FirstPage"] = FirstPage

		c.Data["pageIndex"] = page

		c.Data["users"] = users

		c.TplName = "PersonalInformation.html"

	}

}
