package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/models"
)

type ViewBooksController struct {
	beego.Controller
}

func (c *ViewBooksController) Get() {

	o := orm.NewOrm()
	var recommendedBooks []models.Bookinfo
	page, _ := c.GetInt("pageIndex")
	pageSize := 2
	offset := (page - 1) * pageSize
	//o.QueryTable("Bookinfo").All(&recommendedBooks)

	_, _ = o.QueryTable("Bookinfo").Limit(pageSize, offset).All(&recommendedBooks)

	// 计算总页数和总记录数
	total, _ := o.QueryTable("Bookinfo").Count()
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

	c.Data["recommendedBooks"] = recommendedBooks

	c.TplName = "ViewBooks.html"

}

func (c *ViewBooksController) Post() {
	uname := c.GetString("name")
	fmt.Println(uname)
	o := orm.NewOrm()
	var recommendedBooks []models.Bookinfo
	_, err := o.QueryTable("Bookinfo").Filter("Name", uname).All(&recommendedBooks)
	fmt.Println("----------------------------------")
	fmt.Println(recommendedBooks)
	if err != nil {
		fmt.Println("查询失败", err)
		c.Data["recommendedBooks"] = nil
		fmt.Println("=========================")
		fmt.Println(recommendedBooks)

	} else {

		page, _ := c.GetInt("pageIndex")
		pageSize := 2
		offset := (page - 1) * pageSize
		//o.QueryTable("Bookinfo").All(&recommendedBooks)

		_, _ = o.QueryTable("Bookinfo").Filter("Name", uname).Limit(pageSize, offset).All(&recommendedBooks)
		fmt.Println("-------------", recommendedBooks)
		// 计算总页数和总记录数
		total, _ := o.QueryTable("Bookinfo").Filter("Name", uname).Count()
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

		c.Data["recommendedBooks"] = recommendedBooks

		c.TplName = "ViewBooks.html"

	}

}
