package routers

import (
	"github.com/astaxie/beego"
	"github.com/tobyzxj/passwordkeeper/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})

	// UI of spuer user
	beego.Router("/user", &controllers.UserController{})
	beego.AutoRouter(&controllers.UserController{})

	// UI of general user
	beego.Router("/passwdtag", &controllers.PasswdtagController{})
	beego.AutoRouter(&controllers.PasswdtagController{})
}
