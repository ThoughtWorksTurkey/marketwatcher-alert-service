package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"strconv"
	"log"
	"net/http/httputil"
)

var IngestionUrl = os.Getenv("DATA_INGESTION_URL") + "/init"
var IngestionServiceNotReachableErr = "Ingestion service could not be reached"
var AlertNotCreatedErr = "Alert could not be created"

var triggerIngestion = func(a Alert) error {
	alertBytes := []byte(`{"id":"` + a.ID.String() + `","name":"` + a.Name + `","requiredCriteria":"` + a.RequiredCriteria + `"}`)
	log.Printf("REQUEST: %s\n", IngestionUrl)
	req, err := http.NewRequest("POST", IngestionUrl, bytes.NewBuffer(alertBytes))
	req.Header.Set("Content-Type", "application/json")
	
	dumpRequest, err := httputil.DumpRequestOut(req, true)
	
	log.Printf("REQUEST: %q\n", dumpRequest)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return errors.New(IngestionServiceNotReachableErr)
	}
	defer resp.Body.Close()
	
	dumpResponse, err := httputil.DumpResponse(resp, true)
	
	log.Printf("REQUEST: %q\n", dumpResponse)

	if resp.Status == (strconv.Itoa(http.StatusOK) + " OK") {
		return nil
	} else {
		return errors.New(AlertNotCreatedErr)
	}
}
