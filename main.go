package main

import (
	_ "beeverse/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{

		AllowAllOrigins: true,

		AllowOrigins: []string{"*"},
		//Optional parameters "GET", "POST", "PUT", "DELETE", "OPTIONS" (* is all)

		AllowMethods: []string{"*"},
		// refers to the type of Header allowed
		AllowHeaders: []string{"*"},
		// public HTTP header list
		ExposeHeaders: []string{"*"},

		AllowCredentials: true,
	}))

	beego.Run()
}
