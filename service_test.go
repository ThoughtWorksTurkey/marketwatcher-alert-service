package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "/")))
	beego.BConfig.CopyRequestBody = true
	beego.TestBeegoInit(apppath)
	beego.Router(alertsCreateUrl, &AlertController{}, "post:CreateAlert")
}

func TestWhenIProvideAlertId_ShouldReturnSampleAlert(t *testing.T) {
	find = MockFind
	id := SampleAlert.ID
	a, err := FindAlert(id.String())

	assert.Equal(t, nil, err, "Find alert returns nil error when id provided")
	assert.Equal(t, SampleAlert.Name, a.Name, "Find alert by id should return alert when id provided")
}

func TestWhenIProvideEmptyAlertId_ShouldReturnError(t *testing.T) {
	find = MockFind
	id := ""
	_, err := FindAlert(id)

	assert.Equal(t, "id should be provided", err.Error(), "Find alert by id should return error when id not provided")
}

func TestWhenTheresValidationErrorServerShouldReturnBadRequest(t *testing.T) {
	save = MockSave
	invalidAlertJson := `{"owner_id": 4, "name": "", "required_criteria": "TW,ThoughtWorks,Thought Works,Thoughtworks",
        "nice_to_have_criteria": "good,best office", "excluded_criteria": "bad,sucks,not good enough", "threshold": 1000,"status": 1}`

	alertBuffer := bytes.NewBuffer([]byte(invalidAlertJson))
	r, _ := http.NewRequest("POST", alertsCreateUrl, alertBuffer)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	errorObj := AlertErrorMessage{}
	buffer := w.Body
	json.Unmarshal(buffer.Bytes(), &errorObj)
	assert.Equal(t, 400, w.Code, "When there is a validation error, service should return BAD REQUEST")
}

func TestWhenTheresAnExistingOwnerIdAlertNameTupleServerShouldReturnConflict(t *testing.T) {
	validAlertJson := `{"owner_id": 4, "name": "alert name", "required_criteria": "TW,ThoughtWorks,Thought Works,Thoughtworks",
            "nice_to_have_criteria": "good,best office", "excluded_criteria": "bad,sucks,not good enough", "threshold": 1000,"status": 1}`

	save = func(a Alert) (Alert, error) {
		return Alert{}, errors.New("Alert name must be unique per owner")
	}

	triggerIngestion = MockTriggerIngestion

	alertBuffer := bytes.NewBuffer([]byte(validAlertJson))
	r, _ := http.NewRequest("POST", alertsCreateUrl, alertBuffer)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	errorObj := AlertErrorMessage{}
	buffer := w.Body
	json.Unmarshal(buffer.Bytes(), &errorObj)
	assert.Equal(t, 409, w.Code, "When there is an existing owner id alert name pair, service should return CONFLICT")
}
