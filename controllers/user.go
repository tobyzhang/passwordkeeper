package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tobyzxj/passwordkeeper/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	// 判断用户是否有权限进入此页面
	if checkLogin(this.Ctx) {
		ck, err := this.Ctx.Request.Cookie("uname")
		if err != nil {
			this.Redirect("/login", 302)
			return
		}

		uname := ck.Value
		if uname != "root" && uname != "admin" {
			this.Redirect("/passwdtag", 302)
			return
		}
	} else {
		this.Redirect("/login", 302)
	}

	// 进入超级用户管理界面
	this.TplNames = "user.html"

	// 检测是否有用户操作
	op := this.Input().Get("op")
	switch op {
	case "add":
		var err string
		name := this.Input().Get("uname")
		email := this.Input().Get("uemail")
		passwd := this.Input().Get("upasswd")
		err = models.AddUser(name, email, passwd)
		if err != "ERR_OK" {
			if err == "ERR_EXIST_UNAME" {
				this.Data["Uname"] = "This username is exist"
			} else if err == "ERR_EXIST_UEMAIL" {
				this.Data["Uemail"] = "This email is exist"
			}

			beego.Error(err)
		} else {
			this.Data["Succeed"] = "Add user succeed"
		}

		this.Redirect("/user", 302)
		return

	case "delete":
		uid := this.Input().Get("uid")
		if len(uid) == 0 {
			break
		}

		err := models.DeleteUserByUid(uid)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/user", 302)
		return
	}

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

	var err error
	this.Data["User"], err = models.GetUserList()
	if err != nil {
		beego.Error(err)
	}
}

// 用户帐号修改
func (this *UserController) Modify() {
	// 判断用户是否有权限进入此页面
	if checkLogin(this.Ctx) {
		ck, err := this.Ctx.Request.Cookie("uname")
		if err != nil {
			this.Redirect("/login", 302)
			return
		}

		uname := ck.Value
		if uname != "root" && uname != "admin" {
			this.Redirect("/passwdtag", 302)
			return
		}
	} else {
		this.Redirect("/login", 302)
	}

	// 进入用户帐号修改界面
	this.TplNames = "user_modify.html"

	// 获取用户帐号信息
	uid := this.Input().Get("uid")
	if len(uid) == 0 {
		this.Redirect("/user", 302)
		return
	}

	// 根据uid得到当前的用户帐号信息
	user, err := models.GetUserByUid(uid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/user", 302)
		return
	}

	this.Data["Uid"] = uid
	this.Data["Username"] = user.Name
	this.Data["Email"] = user.Email
	this.Data["Password"] = "1234567890"

	errno := this.Input().Get("errno")
	if len(errno) != 0 {
		switch errno {
		case "uname":
			this.Data["ModifyStatus"] = "This username " + this.Input().Get("v") + " is exist, modify failed"
		case "uemail":
			this.Data["ModifyStatus"] = "This email " + this.Input().Get("v") + " is exist, modify failed"
		case "succeed":
			this.Data["ModifyStatus"] = "Modify succeed"
		}
	}

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

// 用户帐号修改保存
func (this *UserController) Save() {
	// 判断用户是否有权限进入此页面
	if checkLogin(this.Ctx) {
		ck, err := this.Ctx.Request.Cookie("uname")
		if err != nil {
			this.Redirect("/login", 302)
			return
		}

		uname := ck.Value
		if uname != "root" && uname != "admin" {
			this.Redirect("/passwdtag", 302)
			return
		}
	} else {
		this.Redirect("/login", 302)
	}

	// 解析表单
	var err string
	this.TplNames = "user_modify.html"
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
		err = models.ModifyUser(uid, name, email, passwd)
	}

	if err != "ERR_OK" {
		if err == "ERR_EXIST_UNAME" {
			//this.Data["ModifyStatus"] = "This username is exist, modify failed"
			this.Redirect("/user/modify?uid="+uid+"&errno=uname"+"&v="+name, 302)
		} else if err == "ERR_EXIST_UEMAIL" {
			//this.Data["ModifyStatus"] = "This email is exist, modify failed"
			this.Redirect("/user/modify?uid="+uid+"&errno=uemail"+"&v="+email, 302)
		}

		beego.Error(err)

	} else {
		//this.Data["ModifyStatus"] = "Modify succeed"
		this.Redirect("/user/modify?uid="+uid+"&errno=succeed", 302)
	}
}
