package main

import (
	"errors"
)

var CreateAlert = func(a Alert) (Alert, error) {
	if !a.validate() {
		return Alert{}, errors.New("Validation failed")
	}
	a, err := UpsertAlert(a)
	return a, err
}

var UpsertAlert = func(a Alert) (Alert, error) {
	v, err := PersistAlert(a)
	return v, err
}
