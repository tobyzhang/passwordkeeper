package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tobyzxj/passwordkeeper/models"
)

type PasswdtagController struct {
	beego.Controller
}

func (this *PasswdtagController) Get() {
	// 判断用户是否有权限进入此页面
	if !checkLogin(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "passwdtag.html"

	// Update navbar
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

	// // 读取当前用户所有设置的密码标签
	uid, ok := getLoginUserId(this.Ctx)
	if ok {
		var err error
		this.Data["Passwdtag"], err = models.GetPasswdtagByUid(uid)
		if err != nil {
			beego.Error(err)
		}
	}
}

func (this *PasswdtagController) Add() {
	// 判断用户是否有权限进入此页面
	if !checkLogin(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "passwdtag_add.html"

	// Update navbar
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

	// add new passwdtag

}
