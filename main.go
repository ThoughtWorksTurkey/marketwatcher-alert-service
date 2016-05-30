package main

import (
	"errors"
)

// Alert is the primary entity of this microservice
type Alert struct {
	id                 int
	name               string
	requiredCriteria   []string
	niceToHaveCriteria []string
	excludedCriteria   []string
	threshold          int
	ownerID            int64
}

func (a *Alert) validate() bool {
	return a.id > 0 &&
		a.name != "" &&
		a.requiredCriteria != nil && len(a.requiredCriteria) > 0 &&
		a.threshold > 0 && a.ownerID > 0
}

var alertMap = make(map[int]Alert)

func findAlertByID(id int) Alert {
	v, _ := alertMap[id]
	return v
}

func insertAlert(a Alert) (Alert, error) {
	if !a.validate() {
		return Alert{}, errors.New("Validation failed")
	}
	alertMap[a.id] = a
	v, ok := alertMap[a.id]
	if ok {
		return v, nil
	}

	return v, errors.New("Could NOT insert Alert")
}
