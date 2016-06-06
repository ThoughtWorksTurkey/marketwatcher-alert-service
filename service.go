package main

import (
	"errors"
)

var CreateAlert = func(a Alert) (Alert, error) {
	if !a.validate() {
		return Alert{}, errors.New("Validation failed")
	}
	a, err := upsert(a)
	return a, err
}

var ListAlerts = func(ownerID int) ([]Alert, error) {
	return findByOwner(ownerID)
}
