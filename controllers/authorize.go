package controllers

import (
	"fmt"
	"openid_connect_op/models"

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
	fmt.Println(clientID)
	fmt.Println(redirectURL)
	fmt.Println(scopes)
	fmt.Println(response_type)
	fmt.Println(state)
	a.Data["json"] = models.Response{
		Code:    "200",
		Message: "授权成功",
		Data:    nil,
	}
	a.ServeJSON()
}