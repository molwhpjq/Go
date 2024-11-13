package models

import "github.com/astaxie/beego/orm"

type Bookinfo struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Publisher   string  `json:"publisher"`
	Publishtime string  `json:"bpublishtime"` // 发布时间
}

func (book *Bookinfo) BookToRespDesc() interface{} {
	respBook := map[string]interface{}{
		"id":          book.Id,
		"name":        book.Name,
		"author":      book.Author,
		"price":       book.Price,
		"publisher":   book.Publisher,
		"publishtime": book.Publishtime,
	}
	return respBook
}

// 新增获取推荐图书的方法
func GetRecommendedBooks() ([]Bookinfo, error) {
	o := orm.NewOrm()
	var books []Bookinfo
	_, err := o.QueryTable("Bookinfo").OrderBy("-id").Limit(5).All(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
