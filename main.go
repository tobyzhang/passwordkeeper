package main

import (
	//"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tobyzxj/passwordkeeper/models"
	_ "github.com/tobyzxj/passwordkeeper/routers"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 打开ORM调试
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
