package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

const (
	Normal = iota
	Disable
)

type User struct {
	Id        int
	Name      string
	Password  string
	Avatar    string
	Mobile    string
	Status    int
	CreatedAt int64
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func Profile(id int) (*User, error) {
	o := orm.NewOrm()
	user := &User{Id: id}
	err := o.Read(user)
	return user, err
}

// IsUserMobile 手机号是否已注册
func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := &User{Mobile: mobile}
	if err := o.Read(user, "Mobile"); err == orm.ErrNoRows || err == orm.ErrMissPK {
		return false
	}
	return true
}

func Save(mobile, password string) error {
	o := orm.NewOrm()
	_, err := o.Insert(&User{
		Name:      mobile,
		Mobile:    mobile,
		Password:  password,
		Status:    Normal,
		CreatedAt: time.Now().Unix(),
	})
	return err
}

func IsMobileLogin(mobile, password string) (int, string) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	}
	return user.Id, user.Name
}

// GetUserInfo 用户信息
func GetUserInfo(id int) (info UserInfo, err error) {
	o := orm.NewOrm()
	err = o.Raw("select id,name,add_time,avatar from user where id=? limit 1", id).QueryRow(&info)
	return
}
