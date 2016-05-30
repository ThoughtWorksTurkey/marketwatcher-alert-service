package main

import (
	"errors"
)

var alertMap = make(map[int]Alert)

// FakeRepository is a fake repository
type FakeRepository struct{}

func (fr FakeRepository) find(id int) (Alert, error) {
	v, ok := alertMap[id]

	if ok {
		return v, nil
	}

	return v, errors.New("Could NOT find Alert")
}

func (fr FakeRepository) upsert(a Alert) (Alert, error) {
	if !a.validate() {
		return Alert{}, errors.New("Validation failed")
	}

	alertMap[a.id] = a

	v, ok := alertMap[a.id]

	if ok {
		return v, nil
	}

	return v, errors.New("Could NOT upsert Alert")
}
