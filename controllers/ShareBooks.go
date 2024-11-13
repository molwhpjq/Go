package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/models"
)

type ShareBooks struct {
	beego.Controller
}

// Get 方法是一个HTTP GET请求的处理函数
func (c *ShareBooks) Get() {
	c.TplName = "ShareBooks.html"
}
func (c *ShareBooks) Post() {
	o := orm.NewOrm()
	book := models.ShareBooks{}
	book.Title = c.GetString("title")
	book.Author = c.GetString("author")
	book.Publisher = c.GetString("publisher")
	book.Category = c.GetString("category")
	_, err := o.Insert(&book)
	if err != nil {
		fmt.Println("插入失败")
		return
		c.Redirect("/ShareBooks", 302)

	}

	c.Redirect("/ShareBooks", 302)
}
