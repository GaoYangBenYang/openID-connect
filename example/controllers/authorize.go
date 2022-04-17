package controllers

import (
	"openid_connect/example/models"

	"github.com/astaxie/beego"
)

//授权
type Authorize struct {
	beego.Controller
}

func (a *Authorize) Authorize() {
	clientID:=a.GetString("client_id")
	redirectURL:=a.GetString("redirect_uri")
	scopes:=a.GetString("scope")
	response_type:=a.GetString("response_type")
	state:=a.GetString("state")
	a.Data["json"] = models.Response{
		Code:    "200",
		Message: "Authorize",
		Data:    []string{clientID,redirectURL,scopes,response_type,state},
	}
	a.ServeJSON()
}

func (a *Authorize) CallBack() {
	code:=a.GetString("code")
	state:=a.GetString("state")
	a.Data["json"] = models.Response{
		Code:    "200",
		Message: "CallBack",
		Data:    []string{code,state},
	}
	a.ServeJSON()
}