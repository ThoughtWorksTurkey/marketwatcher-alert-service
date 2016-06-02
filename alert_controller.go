package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

// AlertController is a controller
type AlertController struct {
	beego.Controller
}

// Get is a receiver method
func (alertReceiver *AlertController) Get() {
	alertReceiver.Ctx.WriteString("hello get")
}

// Post is a receiver method
func (alertReceiver *AlertController) Post() {
	var alert Alert
	json.Unmarshal(alertReceiver.Ctx.Input.RequestBody, &alert)
	createdAlert, err := CreateAlert(alert)
	if err != nil {
		alertReceiver.Data["error"] = err
		alertReceiver.ServeJSONWithStatus(400)
	} else {
		alertReceiver.Data["json"] = createdAlert
		alertReceiver.ServeJSONWithStatus(200)
	}

}

// ServeJSONWithStatus encapsulates responsing with http status code
func (alertReceiver *AlertController) ServeJSONWithStatus(code int) {
	alertReceiver.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	alertReceiver.Ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	if enablegzip, err := beego.AppConfig.Bool("enablegzip"); err == nil && enablegzip {
		alertReceiver.Ctx.Output.Header("Content-Encoding", "gzip")
	}
	alertReceiver.Ctx.Output.SetStatus(code)
	alertReceiver.ServeJSON()
}
