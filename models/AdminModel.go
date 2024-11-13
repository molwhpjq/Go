package models

//这段代码定义了一个Admin结构体， 表示一个管理员，

type Admin struct {
	Id    int    `json:"id"`
	Aname string `json:"aname"`
	Apwd  string `json:"apwd"`
}

//定义了一个方法AdminToRespDesc，该方法将Admin结构体转换为一个响应描述（即一个键值对的映射），
//其中键是属性的名字，值是属性的值。这种转换通常用于将结构体数据转换为JSON或其他格式的数据

func (admin *Admin) AdminToRespDesc() interface{} {
	respAdmin := map[string]interface{}{
		"id":    admin.Id,
		"aname": admin.Aname,
		"apwd":  admin.Apwd,
	}
	return respAdmin
}
