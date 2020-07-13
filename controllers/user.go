package controllers

import (
	"github.com/astaxie/beego/validation"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"youku/models"
	"youku/utils"

	"github.com/astaxie/beego"
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
	mobile := u.GetString("mobile")
	if mobile == "" {
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

// @router /login/do [*]
func (u *UserController) Login() {
	mobile := u.GetString("mobile")
	if mobile == "" {
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
	uid, name := models.IsMobileLogin(mobile, MD5(password))
	if uid != 0 {
		u.Data["json"] = Success(0, "登录成功", map[string]interface{}{
			"uid":      uid,
			"username": name,
		}, 1)
		u.ServeJSON()
	} else {
		u.Data["json"] = Fail(4004, "手机号或密码错误")
		u.ServeJSON()
	}
}

//批量发送消息
// @router /send/message [*]
func (c *UserController) Send() {
	uids := c.GetString("uids")
	content := c.GetString("content")

	valid := validation.Validation{}
	valid.Required(uids, "uids").Message("接收人不能为空")
	valid.Required(content, "content").Message("发送内容不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	id, err := models.SendMsg(content)
	if err == nil {
		for _, v := range strings.Split(uids, ",") {
			userID, _ := strconv.Atoi(v)
			_ = models.Sender(userID, id)
		}
		c.Data["json"] = Success(0, "发送成功", "", 1)
	} else {
		c.Data["json"] = Fail(5000, "发送失败请联系客服~")
	}

	c.ServeJSON()
}

func (c *UserController) VideoUpload() {
	r := c.Ctx.Request
	uid := r.FormValue("uid")
	// 获取文件流
	file, header, _ := r.FormFile("file")
	// 转为二进制流
	b, _ := ioutil.ReadAll(file)
	// 生成文件名
	filename := strings.Split(header.Filename, ".")
	filename[0] = utils.GetVideoName(uid)
	// 文件保存位置
	fileDir := "/Users/xuyunlong/go/src/fyouku/static" + filename[0] + "." + filename[1]
	playUrl := "/static/video/" + filename[0] + "." + filename[1]
	err := ioutil.WriteFile(fileDir, b, os.ModePerm)
	var title string
	if err != nil {
		title = Success(0, playUrl, nil, 1)
	} else {
		title = Fail(5000, "上传失败，请联系客服")
	}
	c.Ctx.WriteString(title)
}
