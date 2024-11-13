package routers

import (
	"github.com/astaxie/beego"
	"testbeegodemo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/userlist", &controllers.UserlistController{})
	beego.Router("/getall", &controllers.UserlistController{}, "GET:GetAll")
	beego.Router("/findById/:id", &controllers.UserlistController{}, "GET:FindById")
	beego.Router("/deleteUser/:id", &controllers.UserlistController{}, "DELETE:DeleteUser")

	//首页
	beego.Router("/HomeBlock", &controllers.HomeController{})
	beego.Router("/SelectPage", &controllers.SelectPage{})

	//用户模块
	beego.Router("/register", &controllers.RegisterController{}, "get:Get;post:Register")
	beego.Router("/login", &controllers.UserController{})
	beego.Router("/UserPage", &controllers.UserPageController{})
	beego.Router("/UserShow", &controllers.UserShow{})
	beego.Router("/personalInformation", &controllers.PersonalInformation{})
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/ShareBooks", &controllers.ShareBooks{})

	//管理员模块
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/AdminEntry", &controllers.AdminEntryController{})
	beego.Router("/ViewBooks", &controllers.ViewBooksController{})
	beego.Router("/delete", &controllers.PersonalInformation{}, "get:Delete")
	beego.Router("/userdelete", &controllers.UserShow{}, "get:Delete")
	beego.Router("/update", &controllers.PersonalInformation{}, "get:Show;post:Update")
	beego.Router("/userupdate", &controllers.UserShow{}, "get:Show;post:Update")

}
