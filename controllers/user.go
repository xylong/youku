package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"youku/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /register/save [post]
func (u *UserController) Post() {
	mobile:=u.GetString("mobile")
	if mobile=="" {
		u.Data["json"] = Fail(4001, "手机号不能为空")
		u.ServeJSON()
	}
	if ok, _ := regexp.MatchString(`^1(3|4|5|6|7|8|9)[0-9\d{8$}]`, mobile); !ok {
		u.Data["json"] = Fail(4002, "手机号格式错误")
		u.ServeJSON()
	}

	password := u.GetString("password")
	if password == "" {
		u.Data["json"] = Fail(4003, "密码不能为空")
		u.ServeJSON()
	}
	if models.IsUserMobile(mobile) {
		u.Data["json"] = Fail(4005, "手机号已注册")
		u.ServeJSON()
	} else {
		err := models.Save(mobile, MD5(password))
		if err != nil {
			u.Data["json"] = Fail(5000, err)
			u.ServeJSON()
		} else {
			u.Data["json"] = Success(0, "注册成功", nil, 0)
			u.ServeJSON()
		}
	}
}
