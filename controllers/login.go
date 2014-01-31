package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/tobyzxj/passwordkeeper/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("passwd", "", -1, "/")
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "login.html"
	this.Data["IsInOut"] = true
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

func (this *LoginController) Post() {
	this.TplNames = "login.html"
	this.Data["IsInOut"] = true
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

	// 获取表单信息
	uname := this.Input().Get("uname")
	passwd := this.Input().Get("passwd")
	autoLogin := this.Input().Get("autoLogin") == "remember-me"

	if len(uname) == 0 || len(passwd) == 0 {
		this.Data["LoginStatus"] = "Login failed"
		return
	}

	b := models.CheckAccount(uname, passwd)
	if b != true {
		this.Data["LoginStatus"] = "Login failed"
		return
	}

	// 登录成功
	maxAge := 0
	if autoLogin {
		maxAge = 1<<31 - 1
	}
	this.Ctx.SetCookie("uname", uname, maxAge, "/")
	this.Ctx.SetCookie("passwd", passwd, maxAge, "/")

	// 判断是否是超级用户 "root" "admin"
	// 如果是超级用户,打开用户管理界面
	// 否则进入用户操作界面，密码标签管理界面
	if uname == "root" || uname == "admin" {
		this.Redirect("/user", 302)
		return
	}

	this.Redirect("/passwdtag", 302)
}

// 检查用户是否已登录
func checkLogin(ctx *context.Context) bool {
	// 请求用户本地Cookie用户名密码值
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("passwd")
	if err != nil {
		return false
	}
	passwd := ck.Value

	return models.CheckAccount(uname, passwd)
}

// 获取登录的用户名
func getLoginUserName(ctx *context.Context) (string, bool) {
	// 请求用户本地Cookie用户名密码值
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return "", false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("passwd")
	if err != nil {
		return "", false
	}
	passwd := ck.Value

	return uname, models.CheckAccount(uname, passwd)
}
