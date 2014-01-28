package routers

import (
	"github.com/astaxie/beego"
	"github.com/tobyzxj/passwordkeeper/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
