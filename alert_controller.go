package main

import (
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
)

// AlertController is a controller
type AlertController struct {
	beego.Controller
}

type AlertErrorMessage struct {
	Message                 string `json:"alert_error_message"`
}

// Get is a receiver method
func (alertReceiver *AlertController) GetAlertsOfOwner() {
	requestedOwnerID, _ := strconv.Atoi(alertReceiver.Ctx.Input.Param(":owner_id"))

	alerts, err := ListAlerts(requestedOwnerID)

	if err != nil {
		alertReceiver.Data["error"] = err
		alertReceiver.ServeJSONWithStatus(404)
	} else {
		alertReceiver.Data["json"] = alerts
		alertReceiver.ServeJSONWithStatus(200)
	}
}

// Post is a receiver method
func (alertReceiver *AlertController) PostNewAlert() {
	var alert Alert
	json.Unmarshal(alertReceiver.Ctx.Input.RequestBody, &alert)
	createdAlert, err := CreateAlert(alert)
	if err != nil {
		alertReceiver.Data["json"] = AlertErrorMessage{ Message: err.Error()}
		alertReceiver.ServeJSONWithStatus(400)
	} else {
		alertReceiver.Data["json"] = createdAlert
		alertReceiver.ServeJSONWithStatus(200)
	}
}

func (alertReceiver *AlertController) GetAlertById() {
	alertId := alertReceiver.Ctx.Input.Param(":id")

	alert, err := FindAlert(alertId)

	if err != nil {
		alertReceiver.Ctx.WriteString(err.Error())
		alertReceiver.ServeJSONWithStatus(404)
	} else {
		alertReceiver.Data["json"] = alert
		alertReceiver.ServeJSONWithStatus(200)
	}
}

// ServeJSONWithStatus decorates responses
func (alertReceiver *AlertController) ServeJSONWithStatus(code int) {
	alertReceiver.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	alertReceiver.Ctx.Output.SetStatus(code)
	alertReceiver.ServeJSON()
}
