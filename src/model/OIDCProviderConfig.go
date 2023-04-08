package model

import (
	"fmt"

	"github.com/astaxie/beego"
)

type OIDCProviderConfig struct {
	Issuer       string
	AuthURL      string
	TokenURL     string
	UserInfoURL  string
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

var OPConfig OIDCProviderConfig

func init() {
    fmt.Println(beego.AppConfig.String("Issuer"))
	OPConfig.Issuer = beego.AppConfig.String("Issuer")
	OPConfig.ClientID = beego.AppConfig.String("ClientID")
	OPConfig.ClientSecret = beego.AppConfig.String("ClientSecret")
	OPConfig.RedirectURI = beego.AppConfig.String("RedirectURI")
}