package main

import (
	"bytes"
	"net/http"
	"errors"
	"strconv"
	"os"
)

var AlertNotCreatedErr = "Aalert was not created"

var triggerIngestion = func(a Alert) error {
	url := os.Getenv("DATA_INGESTION_URL")
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
