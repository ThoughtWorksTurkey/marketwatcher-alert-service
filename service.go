package main

import (
	"errors"
)

func findUserAlerts(r Repository, userID int) []Alert {
	v, _ := r.findByOwnerID(userID)
	// TODO catch errors
	return v
}

func saveAlert(r Repository, a Alert) (Alert, error) {
	if !a.validate() {
		return Alert{}, errors.New("Validation failed")
	}

	v, err := r.upsert(a)
	return v, err
}
