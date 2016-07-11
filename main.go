//+build !unit

package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/api/alerts/owner_id/:owner_id([0-9]+)", &AlertController{}, "get:GetAlertsOfOwner")
	beego.Router("/api/alerts/id/:id", &AlertController{}, "get:GetAlertById")
	beego.Router("/api/alerts", &AlertController{}, "post:PostNewAlert")

	beego.Run()
}
