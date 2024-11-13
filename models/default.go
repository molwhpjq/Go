package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"testbeegodemo/util"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	driverName := beego.AppConfig.String("driverName")
	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)
	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		fmt.Println(err) //打印错误
		util.LogError("连接数据库出错")
		return
	}
	util.LogInfo("连接数据库成功")

	//register model : 注册实体模型
	orm.RegisterModel(new(Userinfo))
	orm.RegisterModel(new(Admin))
	orm.RegisterModel(new(Bookinfo))
	orm.RegisterModel(new(ShareBooks))
	orm.RegisterModel(new(Profile))

	//)
	//the last step: create table生成表
	orm.RunSyncdb("default", false, true)
	//启用ORM日志
	orm.Debug = true
}
