package main

import (
	"github.com/astaxie/beego"
	_ "openid_connect_op/routers"
	_ "openid_connect_op/db"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	beego.Run()
}