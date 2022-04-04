package controllers

import (
	"github.com/astaxie/beego"
	"openid_connect_op/models"
)

type Provider struct {
	beego.Controller
}

func (u *Provider) Get() {
	u.Data["json"] = models.Response{
		Code:    "404",
		Message: "请求方法异常",
		Data:    nil,
	}
	u.ServeJSON()
}

func (u *Provider) Post() {
	u.Data["json"]=models.Response{
		Code:    "200",
		Message: "Healthy",
		Data:    nil,
	}
	u.ServeJSON()
}

func (u *Provider) RegisterOp()  {
	u.Data["json"] = models.Response{
		Code:    "200",
		Message: "注册成功",
		Data:    nil,
	}
	u.ServeJSON()
}