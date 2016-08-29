package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"os"
	"strconv"
)

// AlertController is a controller
type AlertController struct {
	beego.Controller
}

type AlertErrorMessage struct {
	Message string `json:"message"`
}

func (alertReceiver *AlertController) ServeErrorWithStatus(httpStatus int, err error) {
	alertReceiver.ServeErrorWithStatusAndMessage(httpStatus, err, err.Error())
}

func (alertReceiver *AlertController) ServeErrorWithStatusAndMessage(httpStatus int, err error, message string) {
	if err != nil {
		alertReceiver.Data["json"] = AlertErrorMessage{Message: message}
		alertReceiver.ServeJSONWithStatus(httpStatus)
	}
}

// Get is a receiver method
func (alertReceiver *AlertController) GetAlertsOfOwner() {
	requestedOwnerID, _ := strconv.Atoi(alertReceiver.Ctx.Input.Param(":id"))

	alerts, err := ListAlerts(requestedOwnerID)

	if err != nil {
		alertReceiver.ServeErrorWithStatus(404, err)
	} else {
		alertReceiver.Data["json"] = alerts
		alertReceiver.ServeJSONWithStatus(200)
	}
}

// Post is a receiver method
func (alertReceiver *AlertController) CreateAlert() {
	var alert Alert
	json.Unmarshal(alertReceiver.Ctx.Input.RequestBody, &alert)

	validationErr := alert.validate()

	if validationErr != nil {
		alertReceiver.ServeErrorWithStatus(400, validationErr)
		return
	}

	alert.ID = GenerateAlertId()

	err := triggerIngestion(alert)
	if err != nil {
		log.Println("Could not trigger ingestion service at (" + os.Getenv("DATA_INGESTION_URL") + ")")
		alertReceiver.ServeErrorWithStatus(500, err)
		return
	}

	createdAlert, err := save(alert)
	if err != nil {
		if err.Error() == ALERT_NAME_MUST_BE_UNIQUE_PER_OWNER {
			alertReceiver.ServeErrorWithStatus(409, err)
		} else {
			alertReceiver.ServeErrorWithStatus(500, err)
		}
		return
	}

	alertReceiver.Data["json"] = createdAlert
	alertReceiver.ServeJSONWithStatus(200)
}

func (alertReceiver *AlertController) GetAlertById() {
	alertId := alertReceiver.Ctx.Input.Param(":id")

	alert, err := FindAlert(alertId)

	if err != nil {
		alertReceiver.Ctx.WriteString(err.Error())
		alertReceiver.ServeJSONWithStatus(500)
	} else if (Alert{}) == alert {
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
