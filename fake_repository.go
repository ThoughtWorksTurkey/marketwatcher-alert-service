package main

import (
	"errors"
)

var alertMap = make(map[int][]Alert)

// FakeRepository is a fake repository
type FakeRepository struct{}

func (fr FakeRepository) findByOwnerID(id int) ([]Alert, error) {
	return []Alert{Alert{}}, nil
}

func (fr FakeRepository) findByName(id int, name string) (Alert, error) {
	return Alert{}, nil
}

func (fr FakeRepository) upsert(a Alert) (Alert, error) {
	return Alert{}, errors.New("Could NOT upsert Alert")
}
