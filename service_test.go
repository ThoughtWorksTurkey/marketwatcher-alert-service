package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenIInsertValidAlert_ShouldReturnOK(t *testing.T) {
	save = MockSave

	a, err := CreateAlert(SampleAlert)
	assert.Equal(t, a.Name, SampleAlert.Name, "Create should return OK for valid return")
	assert.Nil(t, err, nil, "Create should not return OK for valid return")
}

func TestWhenIInsertAlertWithoutName_ShouldReturnError(t *testing.T) {
	save = MockSave
	alertWithoutName := SampleAlert
	alertWithoutName.Name = ""

	a, err := CreateAlert(alertWithoutName)
	assert.EqualError(t, err, "Validation failed", "Create alert without name should return error")
	assert.Equal(t, Alert{}, a, "Create alert without name should return error")
}

func TestWhenIInsertAlertWithoutOwnerID_ShouldReturnError(t *testing.T) {
	save = MockSave
	alertWithoutOwnerID := SampleAlert
	alertWithoutOwnerID.OwnerID = -2

	a, err := CreateAlert(alertWithoutOwnerID)
	assert.EqualError(t, err, "Validation failed", "Create alert without owner id should return error")
	assert.Equal(t, Alert{}, a, "Create alert without owner id should return error")
}

func TestWhenIInsertAlertWithNameLengthIsMoreThanMax_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.Name = "ayse jkajshdkjashdjsahd kashdjashdkjahsdkjhaskjdhaksjhdkjashd jashdkjahsdkjahsdjhsd"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with name length is more than max should return error")
	assert.Equal(t, Alert{}, a, "Create alert with name length is more than max should return error")
}

func TestWhenIInsertTurkishCharacterForCriteria_ShouldReturnOk(t *testing.T) {
	save = MockSave

	alert := SampleAlert
	alert.RequiredCriteria = "ayçe çç öö ğ ü ı şşşşşşş"

	a, err := CreateAlert(SampleAlert)
	assert.Equal(t, a.Name, SampleAlert.Name, "Create alert with turkish character for criteria should return OK")
	assert.Nil(t, err, nil, "Create alert with turkish character for criteria should not return OK")
}

func TestWhenIInsertNonAlphanumericCharacterForCriteria_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.RequiredCriteria = ">>> < | ~~~ ]"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with non alphanumeric creiteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with  non alphanumeric creiteria should return error")
}

func TestWhenIInsertRequiredCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.RequiredCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with long required criteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with long required criteria should return error")
}

func TestWhenIInsertNiceToHaveCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.NiceToHaveCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with long nice-to-have criteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with long nice-to-have criteria should return error")
}

func TestWhenIInsertExcludedCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.ExcludedCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with long excluded criteria should return error")
	assert.Equal(t, Alert{}, a, "Create alert with long excluded criteria should return error")
}

func TestWhenIInsertAlertWithInvalidThreshold_ShouldReturnError(t *testing.T) {
	save = MockSave
	alert := SampleAlert
	alert.Threshold = 2000000

	a, err := CreateAlert(alert)
	assert.EqualError(t, err, "Validation failed", "Create alert with greater than max valued threshold should return error")
	assert.Equal(t, Alert{}, a, "Create alert with greater than max valued threshold should return error")
}

func TestWhenIProvideAlertId_ShouldReturnSampleAlert(t *testing.T) {
	find = MockFind
	id := SampleAlert.ID
	a, err := FindAlert(id.String())

	assert.Equal(t, err, nil, "Find alert returns nil error when id provided")
	assert.Equal(t, SampleAlert.Name, a.Name, "Find alert by id should return alert when id provided")
}

func TestWhenIProvideEmptyAlertId_ShouldReturnError(t *testing.T) {
	find = MockFind
	id := ""
	_, err := FindAlert(id)

	assert.EqualError(t, err, "id should be provided", "Find alert by id should return error when id not provided")
}