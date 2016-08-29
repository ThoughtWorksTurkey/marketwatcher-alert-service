//+build !unit

package main

import (
	"github.com/astaxie/beego"
)

func main() {
	defer session.Close()
	beego.Router("/api/alerts/id/:id", &AlertController{}, "get:GetAlertById")
	beego.Router("/api/alerts", &AlertController{}, "post:CreateAlert")
    beego.Router("/api/alerts/owner_id/:id", &AlertController{}, "get:GetAlertsOfOwner")

	beego.Run()
}
