package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.Router("/alerts", &AlertController{})
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"http://*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}
