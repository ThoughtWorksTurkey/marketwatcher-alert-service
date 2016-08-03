package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"bytes"
)

var alertsCreateUrl = "/api/alerts"

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "/")))
	beego.BConfig.CopyRequestBody = true
	beego.TestBeegoInit(apppath)
	beego.Router(alertsCreateUrl, &AlertController{}, "post:PostNewAlert")
	ingestionServer := stubbedIngestionServiceForBadRequest()
	IngestionUrl = ingestionServer.URL
	save = MockSave
}

func stubbedIngestionServiceForBadRequest() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Ingestion failed", http.StatusBadRequest)
	}))
}

func TestWhenIngestionServiceReturnsBadRequest_alertShouldNotBeCreated(t *testing.T) {
	alertJson := `{
	"owner_id":            4,
	"name":               "Test2",
	"required_criteria":   "TW,ThoughtWorks,Thought Works,Thoughtworks",
	"nice_to_have_criteria": "good,best office",
	"excluded_criteria":   "bad,sucks,not good enough",
	"threshold":          1000,
	"status":             1
}`
	alertBuffer := bytes.NewBuffer([]byte(alertJson))
	r, _ := http.NewRequest("POST", alertsCreateUrl, alertBuffer)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	errorObj := AlertErrorMessage{}
	buffer := w.Body
	json.Unmarshal(buffer.Bytes(), &errorObj)
	assert.Equal(t, AlertNotCreatedErr, errorObj.Message, "When Ingestion Service returns bad request Alert service should return error")
}
