package main

import (
	"errors"
	"github.com/gocql/gocql"
	"log"
)

var ListAlerts = func(ownerID int) ([]Alert, error) {
	return findByOwner(ownerID)
}

var FindAlert = func(id string) (Alert, error) {
	uuid, err := gocql.ParseUUID(id)
	log.Printf("Parsed UUID: %v\n", uuid.String())
	if err != nil {
		return Alert{}, errors.New("id should be provided")
	} else {
		return find(uuid)
	}
}
