package main

import (
	"errors"
	"github.com/gocql/gocql"
)

var CreateAlert = func(a Alert) (Alert, error) {
	validationErr := a.validate()
	if validationErr != nil {
		return a, validationErr
	}

	a.ID = GenerateAlertId()

	err := triggerIngestion(a)
	if err != nil {
		return a, err
	}
	a, err = save(a)
	if err != nil {
		err = errors.New("Alert could not be created")
	}
	return a, err
}

var ListAlerts = func(ownerID int) ([]Alert, error) {
	return findByOwner(ownerID)
}

var FindAlert = func(id string) (Alert, error) {
	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		return Alert{}, errors.New("id should be provided")
	} else {
		return find(uuid)
	}
}
