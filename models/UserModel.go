package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

type Userinfo struct {
	Id    int    `json:"id"`
	Uname string `json:"uname"`
	Upwd  string `json:"upwd"`
	Uage  int    `json:"uage"`
	Usex  string `json:"usex"`
}

type UnifiedResponse struct {
	Code    int         `json:"code"`    // 响应状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

func (user *Userinfo) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":    user.Id,
		"uname": user.Uname,
		"uage":  user.Uage,
		"usex":  user.Usex,
	}
	return respInfo
}

func GetUsers(page, pageSize int) ([]Userinfo, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("userinfo").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var users []Userinfo
	_, err := o.QueryTable("userinfo").Limit(pageSize, (page-1)*pageSize).All(&users)
	return users, err
}

func GetUserById(id int) (*Userinfo, error) {
	o := orm.NewOrm()
	user := Userinfo{Id: id}
	err := o.Read(&user)
	if err != nil {

		return nil, err

	}
	return &user, nil

}

func GetBooks(page, pageSize int) ([]Bookinfo, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("bookinfo").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var books []Bookinfo
	_, err := o.QueryTable("bookinfo").Limit(pageSize, (page-1)*pageSize).All(&books)
	return books, err
}
