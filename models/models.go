package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// 设置数据库路径
	DB_NAME = "data/passwd.db"
	// 设置数据库名称
	SQLITE3_DRIVER = "sqlite3"
)

// 用户注册信息
type User struct {
	Id       int64
	Name     string `orm:"index"`
	Email    string `orm:"index"`
	Passwd   string
	Created  time.Time `orm:"index"`
	Modified time.Time
}

// 用户操作记录
type UserOpRecord struct {
	Id         int64
	Uid        int64
	LoginTimes int64
}

// 密码标签
type Passwdtag struct {
	Id       int64
	Uid      int64
	Tag      string `orm:"index"`
	Account  string
	Passwd   string
	Url      string    `orm:"index"`
	Remark   string    `orm:"size(5000)"`
	Created  time.Time `orm:"index"`
	Modified time.Time
}

func RegisterDB() {
	// 检查数据库文件是否已存在
	if !com.IsExist(DB_NAME) {
		os.MkdirAll(path.Dir(DB_NAME), os.ModePerm)
		os.Create(DB_NAME)
	}

	// 注册驱动模型
	orm.RegisterModel(new(User), new(UserOpRecord), new(Passwdtag))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver(SQLITE3_DRIVER, orm.DR_Sqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", SQLITE3_DRIVER, DB_NAME, 10)
}

// 添加用户
func AddUser(name, email, passwd string) string {
	if len(name) == 0 || len(email) == 0 || len(passwd) == 0 {
		return "ERR_PARAM"
	}

	o := orm.NewOrm()

	// 检查此用户名和邮箱是否已被注册过
	// A.检测用户名
	user_name := &User{Name: name}
	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user_name)
	if err == nil {
		beego.Info("AdderUser:user_name exist")
		return "ERR_EXIST_UNAME"
	} else {
		// B.检测邮箱
		user_email := &User{Email: email}
		err = qs.Filter("email", email).One(user_email)
		if err == nil {
			beego.Info("AdderUser:user_email exist")
			return "ERR_EXIST_UEMAIL"
		}
	}

	h := md5.New()
	io.WriteString(h, passwd)
	passwd_md5 := fmt.Sprintf("%x", h.Sum(nil))

	// 添加新用户
	user := &User{
		Name:     name,
		Email:    email,
		Passwd:   passwd_md5,
		Created:  time.Now(),
		Modified: time.Now(),
	}
	_, err = o.Insert(user)
	if err != nil {
		return "ERR_ORM"
	}

	return "ERR_OK"
}

// 修改用户
func ModifyUser(uid, name, email, passwd string) string {
	if len(uid) == 0 || len(name) == 0 || len(email) == 0 {
		return "ERR_PARAM"
	}

	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return "ERR_CONV"
	}

	// 检查此用户名和邮箱是否已被注册过
	// A.检测用户名
	o := orm.NewOrm()

	user_name := &User{Name: name}
	qs := o.QueryTable("user")
	err = qs.Filter("name", name).One(user_name)
	if err == nil {
		if user_name.Id != id {
			beego.Info("AdderUser:user_name exist")
			return "ERR_EXIST_UNAME"
		}
	} else {
		// B.检测邮箱
		user_email := &User{Email: email}
		err = qs.Filter("email", email).One(user_email)
		if err == nil {
			if user_email.Id != id {
				beego.Info("AdderUser:user_email exist")
				return "ERR_EXIST_UEMAIL"
			}
		}
	}

	h := md5.New()
	io.WriteString(h, passwd)
	passwd_md5 := fmt.Sprintf("%x", h.Sum(nil))

	// 修改用户信息
	user := &User{Id: id}
	err = o.Read(user)
	if err == nil {
		user.Name = name
		user.Email = email
		if len(passwd) != 0 {
			user.Passwd = passwd_md5
		}
		user.Modified = time.Now()
		_, err = o.Update(user)
		if err != nil {
			return "ERR_ORM"
		}
	} else {
		return "ERR_ORM"
	}

	return "ERR_OK"
}

// 检测用户名和密码是否正确
func CheckAccount(account, passwd string) bool {
	o := orm.NewOrm()

	h := md5.New()
	io.WriteString(h, passwd)
	passwd_md5 := fmt.Sprintf("%x", h.Sum(nil))

	// 验证用户名和密码是否正确
	// A. 用户帐号为用户名帐号
	user_name := &User{Name: account}
	qs := o.QueryTable("user")
	err := qs.Filter("name", account).One(user_name)
	if err == nil {
		// 找到用户记录，判断密码是否正确
		if user_name.Passwd != passwd_md5 {
			return false
		} else {
			return true
		}
	}

	// B.用户帐号可能是邮箱帐号
	user_email := &User{Email: account}
	qs = o.QueryTable("user")
	err = qs.Filter("email", account).One(user_email)
	if err == nil {
		// 找到用户记录，判断密码是否正确
		if user_email.Passwd != passwd_md5 {
			return false
		} else {
			return true
		}
	}

	// 用户帐号没有找到
	return false
}

// 获取用户列表
func GetUserList() ([]*User, error) {
	o := orm.NewOrm()

	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}

// 删除用户
func DeleteUserByUid(uid string) error {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	user := &User{Id: id}
	_, err = o.Delete(user)
	return err
}

// 根据uid获取用户信息
func GetUserByUid(uid string) (User, error) {
	var user User

	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return user, err
	}

	o := orm.NewOrm()
	user.Id = id
	err = o.Read(&user)
	return user, err
}

// 根据用户名查询用户信息
func GetUserByName(uname string) (User, error) {
	var user User

	o := orm.NewOrm()

	// A.帐号可能是用户名
	qs := o.QueryTable("user")
	err := qs.Filter("name", uname).One(&user)
	if err == nil {
		return user, nil
	}

	// B.帐号可能是
	err = qs.Filter("email", uname).One(&user)
	if err == nil {
		return user, nil
	}

	// not find
	return user, err
}

// 根据用户名查找该用户的所有密码标签
func GetPasswdtagByUid(uid string) ([]*Passwdtag, error) {
	var passwdtag Passwdtag
	passwdtags := make([]*Passwdtag, 0)

	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	passwdtag.Uid = id
	qs := o.QueryTable("passwdtag")
	_, err = qs.Filter("Uid", id).All(&passwdtags)
	return passwdtags, err
}

// 添加一个密码标签
func AddPasswdTag(uid, passwdtag, account, password, url, remark string) string {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return "ERR_CONV"
	}

	o := orm.NewOrm()

	// 添加新密码标签
	tag := &Passwdtag{
		Uid:      id,
		Tag:      passwdtag,
		Account:  account,
		Passwd:   password,
		Url:      url,
		Remark:   remark,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	_, err = o.Insert(tag)
	if err != nil {
		return "ERR_ORM"
	}

	return "ERR_OK"
}

// 修改一个密码标签
func ModifyPasswdTag(tid, passwdtag, account, password, url, remark string) string {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return "ERR_CONV"
	}

	o := orm.NewOrm()

	tag := &Passwdtag{Id: id}
	err = o.Read(tag)
	if err == nil {
		tag.Tag = passwdtag
		tag.Account = account
		tag.Passwd = password
		tag.Url = url
		tag.Remark = remark
		tag.Modified = time.Now()
		_, err = o.Update(tag)
		if err != nil {
			return "ERR_ORM"
		}
	} else {
		return "ERR_ORM"
	}

	return "ERR_OK"
}

// 删除一个密码标签
func DeletePasswdTagByTid(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	tag := &Passwdtag{Id: id}
	_, err = o.Delete(tag)
	return err
}

// 根据tid获取一个指定的password tag
func GetPasswdTagByTid(tid string) (Passwdtag, error) {
	var tag Passwdtag

	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return tag, err
	}

	o := orm.NewOrm()
	tag.Id = id
	err = o.Read(&tag)
	return tag, err
}
