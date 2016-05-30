package main

import (
	"errors"
)

func findAlertByID(r Repository, id int) Alert {
	v, _ := r.find(id)
	return v
}

func saveAlert(r Repository, a Alert) (Alert, error) {
	if !a.validate() {
		return Alert{}, errors.New("Validation failed")
	}

	v, err := r.upsert(a)
	return v, err
}
