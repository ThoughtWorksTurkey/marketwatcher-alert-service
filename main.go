package main

import "github.com/astaxie/beego"

func main() {
	beego.Router("/", &AlertController{})
	beego.Run()
}
