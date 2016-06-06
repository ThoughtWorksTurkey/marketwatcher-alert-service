package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var SampleAlert = Alert{
	ID:                 1,
	OwnerID:            1,
	Name:               "Test Alert",
	RequiredCriteria:   "TW,ThoughtWorks,Thought Works,Thoughtworks",
	NiceToHaveCriteria: "good,best office",
	ExcludedCriteria:   "bad,sucks,not good enough",
	Threshold:          1000,
	Status:             ACTIVE,
}

func TestWhenIInsertValidAlert_ShouldReturnOK(t *testing.T) {
	upsert = MockUpsert

	a, err := CreateAlert(SampleAlert)
	assert.Equal(t, a.Name, SampleAlert.Name, "Create should return OK for valid return")
	assert.Nil(t, err, nil, "Create should not return OK for valid return")
}

func TestWhenIInsertAlertWithoutName_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alertWithoutName := SampleAlert
	alertWithoutName.Name = ""

	a, err := CreateAlert(alertWithoutName)
	assert.EqualError(t, err, "Validation failed", "Create alert without name should return error")
	assert.Equal(t, Alert{}, a, "Create alert without name should return error")
}

func TestWhenIInsertAlertWithoutOwnerID_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alertWithoutOwnerID := SampleAlert
	alertWithoutOwnerID.OwnerID = -2

	a, err := CreateAlert(alertWithoutOwnerID)
	assert.EqualError(t, err, "Validation failed", "Create alert without owner id should return error")
	assert.Equal(t, Alert{}, a, "Create alert without owner id should return error")
}

func TestWhenIInsertAlertWithNameLengthIsMoreThanMax_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alert := SampleAlert
	alert.Name = "ayse jkajshdkjashdjsahd kashdjashdkjahsdkjhaskjdhaksjhdkjashd jashdkjahsdkjahsdjhsd"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with name length is more than max should return error")
	assert.Equal(t, Alert{}, a, "Create alert with name length is more than max should return error")
}

func TestWhenIInsertTurkishCharacterForCriteria_ShouldReturnOk(t *testing.T) {
	upsert = MockUpsert

	alert := SampleAlert
	alert.RequiredCriteria = "ayçe çç öö ğ ü ı şşşşşşş"

	a, err := CreateAlert(SampleAlert)
	assert.Equal(t, a.Name, SampleAlert.Name, "Create alert with turkish character for criteria should return OK")
	assert.Nil(t, err, nil, "Create alert with turkish character for criteria should not return OK")
}

func TestWhenIInsertNonAlphanumericCharacterForCriteria_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alert := SampleAlert
	alert.RequiredCriteria = ">>> < | ~~~ ]"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with non alphanumeric creiteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with  non alphanumeric creiteria should return error")
}

func TestWhenIInsertRequiredCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alert := SampleAlert
	alert.RequiredCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with long required criteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with long required criteria should return error")
}

func TestWhenIInsertNiceToHaveCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alert := SampleAlert
	alert.NiceToHaveCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with long nice-to-have criteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with long nice-to-have criteria should return error")
}

func TestWhenIInsertExcludedCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alert := SampleAlert
	alert.ExcludedCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with long excluded criteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with long excluded criteria should return error")
}

func TestWhenIInsertAlertWithInvalidThreshold_ShouldReturnError(t *testing.T) {
	upsert = MockUpsert
	alert := SampleAlert
	alert.Threshold = 2000000

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with greater than max valued threshold should return error")
	assert.Equal(t, Alert{}, a, "Create alert with greater than max valued threshold should return error")
}
