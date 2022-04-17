package main

import (
	"github.com/astaxie/beego"
	_ "openid_connect/example/routers"
	_ "openid_connect/example/models/db"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	
	
	beego.Run()
}