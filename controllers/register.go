package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tobyzxj/passwordkeeper/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplNames = "register.html"
	this.Data["IsUp"] = true
	this.Data["IsLogin"] = checkLogin(this.Ctx)

	// 用户操作界面
	username, bLogin := getLoginUserName(this.Ctx)
	if bLogin {
		this.Data["UserName"] = username
		if username == "root" || username == "admin" {
			this.Data["UserUrl"] = "user"
		} else {
			this.Data["UserUrl"] = "passwdtag"
		}
	}
}

func (this *RegisterController) Post() {
	// 解析表单
	var err string
	this.TplNames = "register.html"
	this.Data["IsUp"] = true
	this.Data["IsLogin"] = checkLogin(this.Ctx)

	// 用户操作界面
	username, bLogin := getLoginUserName(this.Ctx)
	if bLogin {
		this.Data["UserName"] = username
		if username == "root" || username == "admin" {
			this.Data["UserUrl"] = "user"
		} else {
			this.Data["UserUrl"] = "passwdtag"
		}
	}

	uid := this.Input().Get("uid")
	name := this.Input().Get("uname")
	email := this.Input().Get("uemail")
	passwd := this.Input().Get("upasswd")

	if len(uid) == 0 {
		err = models.AddUser(name, email, passwd)
	} else {
		this.Redirect("/register", 302)
	}

	if err != "ERR_OK" {
		if err == "ERR_EXIST_UNAME" {
			this.Data["Uname"] = "This username is exist"
		} else if err == "ERR_EXIST_UEMAIL" {
			this.Data["Uemail"] = "This email is exist"
		}

		beego.Error(err)
		//this.Redirect("/register", 302)
	} else {
		this.Redirect("/login", 302)
	}
}
