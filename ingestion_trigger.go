package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"strconv"
)

var IngestionUrl = os.Getenv("DATA_INGESTION_URL")
var AlertNotCreatedErr = "Alert could not be created"

var triggerIngestion = func(a Alert) error {
	alertBytes := []byte(`{"id":` + a.ID.String() + `","name":"` + a.Name + `","requiredCriteria":"` + a.RequiredCriteria + `"}`)

	req, err := http.NewRequest("POST", IngestionUrl, bytes.NewBuffer(alertBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(AlertNotCreatedErr)
	}
	defer resp.Body.Close()

	if resp.Status == (strconv.Itoa(http.StatusOK) + " OK") {
		return nil
	} else {
		return errors.New(AlertNotCreatedErr)
	}
}
