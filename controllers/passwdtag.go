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

	// 读取当前用户所有设置的密码标签
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
	tid := this.Input().Get("tid")
	tpasswdtag := this.Input().Get("tpasswdtag")
	taccount := this.Input().Get("taccount")
	tpassword := this.Input().Get("tpassword")
	turl := this.Input().Get("turl")
	tremark := this.Input().Get("tremark")

	// save to database
	uid, ok := getLoginUserId(this.Ctx)
	if ok {
		var errs string
		if len(tid) != 0 {
			errs = models.ModifyPasswdTag(tid, tpasswdtag, taccount, tpassword, turl, tremark)
		} else {
			errs = models.AddPasswdTag(uid, tpasswdtag, taccount, tpassword, turl, tremark)
		}

		if errs != "ERR_OK" {
			if len(tid) != 0 {
				this.Data["ModifyTagStatus"] = "Modify Tag failed"
			} else {
				this.Data["AddTagStatus"] = "Add Tag failed"
			}

			return
		} else {
			this.Redirect("/passwdtag", 302)
		}
	}

	if len(tid) != 0 {
		this.Data["ModifyTagStatus"] = "Modify Tag failed"
	} else {
		this.Data["AddTagStatus"] = "Add Tag failed"
	}
}

// create a new password tag
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

// delete a password tag
func (this *PasswdtagController) Delete() {
	// 判断用户是否有权限进入此页面
	if !checkLogin(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	err := models.DeletePasswdTagByTid(tid)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/passwdtag", 302)
}

// modify a password tag
func (this *PasswdtagController) Modify() {
	// 判断用户是否有权限进入此页面
	if !checkLogin(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "passwdtag_modify.html"

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

	// get a passwdtag by tid, which need to modify
	tid := this.Input().Get("tid")
	passwdtag, err := models.GetPasswdTagByTid(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/passwdtag", 302)
		return
	}

	this.Data["Passwdtag"] = passwdtag
	this.Data["Tid"] = tid
}

func (this *PasswdtagController) View() {
	// 判断用户是否有权限进入此页面
	if !checkLogin(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "passwdtag_view.html"

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

	// view a password tag
	// get a passwdtag by tid, which need to modify
	tid := this.Input().Get("tid")
	passwdtag, err := models.GetPasswdTagByTid(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/passwdtag", 302)
		return
	}

	this.Data["Passwdtag"] = passwdtag
	this.Data["Tid"] = tid
}
