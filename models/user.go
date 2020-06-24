package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(User))
}

const (
	Normal=iota
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

func Profile(id int) (*User,error)  {
	o:=orm.NewOrm()
	user:=&User{Id: id}
	err:=o.Read(user)
	return user,err
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

func Save(mobile,password string) error {
	o:=orm.NewOrm()
	_,err:=o.Insert(&User{
		Name:     mobile,
		Mobile:   mobile,
		Password: password,
		Status: Normal,
		CreatedAt: time.Now().Unix(),
	})
	return err
}
