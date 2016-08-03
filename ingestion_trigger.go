package main

import (
	"bytes"
	"net/http"
	"errors"
	"strconv"
)

var IngestionUrl = "http://localhost:9000/init"
var AlertNotCreatedErr = "Alert could not be created"

var triggerIngestion = func(a Alert) error {
	url := IngestionUrl
	alertBytes := []byte(`{"name":"` + a.Name + `","requiredCriteria":"` + a.RequiredCriteria + `"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(alertBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(AlertNotCreatedErr)
	}
	defer resp.Body.Close()

	if (resp.Status == (strconv.Itoa(http.StatusOK) + " OK")) {
		return nil
	} else {
		return errors.New(AlertNotCreatedErr)
	}
}
