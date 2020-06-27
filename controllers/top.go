package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"youku/models"
)

type TopController struct {
	beego.Controller
}

// ChannelTop 排行榜
// @router /channel/top [*]
func (c *TopController) ChannelTop() {
	id, _ := c.GetInt("channelId")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("必须指定频道")
	valid.Min(id, 1, "id").Message("频道错误")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetChannelTop(id)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}
	c.ServeJSON()
}

// TypeTop 根据类型获取排行榜
// @router /type/top [*]
func (c *TopController) TypeTop() {
	id, _ := c.GetInt("typeId")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("必须指定类型")
	valid.Min(id, 1, "id").Message("类型错误")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = Fail(4001, err.Message)
			c.ServeJSON()
		}
	}

	num, videos, err := models.GetTypeTop(id)
	if err == nil {
		c.Data["json"] = Success(0, "success", videos, num)
	} else {
		c.Data["json"] = Fail(4004, "没有相关内容")
	}
	c.ServeJSON()
}
