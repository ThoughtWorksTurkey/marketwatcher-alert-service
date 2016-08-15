package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

var IngestionUrl = os.Getenv("DATA_INGESTION_URL")
var IngestionServiceNotReachableErr = "Ingestion service could not be reached"
var AlertNotCreatedErr = "Alert could not be created"

var triggerIngestion = func(a Alert) error {
	alertBytes := []byte(`{"id":"` + a.ID.String() + `","name":"` + a.Name + `","requiredCriteria":"` + a.RequiredCriteria + `"}`)
	req, err := http.NewRequest("POST", IngestionUrl, bytes.NewBuffer(alertBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		if reflect.TypeOf(err).String() == "*url.Error" {
			return errors.New(IngestionServiceNotReachableErr)
		}

		return errors.New(AlertNotCreatedErr)
	}
	defer resp.Body.Close()

	if resp.Status == (strconv.Itoa(http.StatusOK) + " OK") {
		return nil
	} else {
		return errors.New(AlertNotCreatedErr)
	}
}
