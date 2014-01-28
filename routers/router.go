package routers

import (
	"github.com/tobyzxj/passwordkeeper/passwordkeeper/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
