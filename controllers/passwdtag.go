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

func (this *PasswdtagController) Post() {
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
	tpasswdtag := this.Input().Get("tpasswdtag")
	taccount := this.Input().Get("taccount")
	tpassword := this.Input().Get("tpassword")
	turl := this.Input().Get("turl")
	tremark := this.Input().Get("tremark")

	// save to database
	uid, ok := getLoginUserId(this.Ctx)
	if ok {
		err := models.AddPasswdTag(uid, tpasswdtag, taccount, tpassword, turl, tremark)
		if err != "ERR_OK" {
			this.Data["AddTagStatus"] = "Add Tag failed"
			return
		} else {
			this.Redirect("/passwdtag", 302)
		}
	}

	this.Data["AddTagStatus"] = "Add Tag failed"
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
}
