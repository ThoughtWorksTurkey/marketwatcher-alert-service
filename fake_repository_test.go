package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fr = FakeRepository{}

func TestWhenIUpsertAlert_ShouldReturnOK(t *testing.T) {
	insertedAlert, err := fr.upsert(sampleAlert)
	assert.Equal(t, sampleAlert, insertedAlert, "Insert alert to map should be successful")
	assert.Nil(t, err, "Error must be nil")
}

func TestWhenIUpsertAlertWithNewName_ShouldReturnOK(t *testing.T) {
	alert, _ := fr.find(1)
	alert.name = "Alert 1"
	updatedAlert, _ := fr.upsert(alert)
	assert.Equal(t, updatedAlert.name, alert.name, "Update alert must return successful")
}

func TestWhenIFindAlertWithID1_ShoudReturnSucessful(t *testing.T) {
	fr.upsert(sampleAlert)
	actualValue, _ := fr.find(1)
	assert.Equal(t, sampleAlert, actualValue, "Must return correct Alert")
}

func TestWhenIFindAlertWithoutExistingID_ShouldReturnNothing(t *testing.T) {
	expected := Alert{}
	actual, _ := fr.find(666)
	assert.Equal(t, expected, actual, "Must return nothing")
}
