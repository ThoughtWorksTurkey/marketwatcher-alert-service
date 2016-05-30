package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleAlert = Alert{
	id:                 1,
	name:               "Test Alert",
	requiredCriteria:   []string{"ali", "veli"},
	niceToHaveCriteria: []string{"ali", "veli"},
	excludedCriteria:   []string{"ali", "veli"},
	threshold:          1000,
	ownerID:            1,
}

func TestWhenIInsertAlert_ShouldReturnOK(t *testing.T) {
	insertedAlert, err := insertAlert(sampleAlert)
	assert.Equal(t, sampleAlert, insertedAlert, "Insert alert to map should be successful")
	assert.Nil(t, err, "Error must be nil")
}

func TestWhenIInsertInvalidAlert_ShouldReturnError(t *testing.T) {
	_, err := insertAlert(Alert{id: -1})
	assert.EqualError(t, err, "Validation failed")
}

func TestWhenIInsertInvalidNamedAlert_ShouldReturnError(t *testing.T) {
	_, err := insertAlert(Alert{id: 1, name: ""})
	assert.EqualError(t, err, "Validation failed")
}

func TestWhenIPassID_1_FindAlertByID_ShoudReturnSucessful(t *testing.T) {
	insertAlert(sampleAlert)
	var actualValue = findAlertByID(1)

	assert.Equal(t, sampleAlert, actualValue, "Must return correct Alert")
}

func TestWhenIPassID_666_FindAlertByID_ShouldReturnNothing(t *testing.T) {
	assert.Equal(t, Alert{id: 0}, findAlertByID(666), "Must return nothing")
}
