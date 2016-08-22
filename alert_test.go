package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWhenInsertValidAlert_ShouldReturnOK(t *testing.T) {
	err := SampleAlert.validate()
	assert.Equal(t, nil, err, "Create should not return OK for valid return")
}

func TestWhenInsertAlertWithoutName_ShouldReturnError(t *testing.T) {
	alertWithoutName := SampleAlert
	alertWithoutName.Name = ""

	err := alertWithoutName.validate()
	assert.Equal(t, VALIDATION_MESSAGE_NAME_EMPTY, err.Error(), "Create alert without name should return error")
}

func TestWhenInsertAlertWithNameLengthIsMoreThanMax_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.Name = "ayse jkajshdkjashdjsahd kashdjashdkjahsdkjhaskjdhaksjhdkjashd jashdkjahsdkjahsdjhsd"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_NAME_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_NAME), err.Error(), "Create alert with name length is more than max should return error")
}

func TestWhenInsertAlertWithoutOwnerID_ShouldReturnError(t *testing.T) {
	alertWithoutOwnerID := SampleAlert
	alertWithoutOwnerID.OwnerID = -2

	err := alertWithoutOwnerID.validate()
	assert.Equal(t, VALIDATION_MESSAGE_OWNER_ID, err.Error(), "Create alert without owner id should return error")
}

func TestWhenInsertTurkishCharacterForCriteria_ShouldReturnOk(t *testing.T) {
	alert := SampleAlert
	alert.RequiredCriteria = "ayçe çç öö ğ ü ı şşşşşşş"
	err := alert.validate()
	assert.Equal(t, nil, err, "Create alert with turkish character for criteria should not return OK")
}

func TestWhenInsertNonAlphanumericCharacterForCriteria_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.RequiredCriteria = ">>> < | ~~~ ]"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_CRITERIA_PHRASES_ALPHANUMERIC, err.Error(), "Create alert with non alphanumeric creiteria should return error")
}

func TestWhenInsertRequiredCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.RequiredCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_REQUIRED_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long required criteria should return error")
}

func TestWhenInsertNiceToHaveCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.NiceToHaveCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_NICE_TO_HAVE_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long nice-to-have criteria should return error")
}

func TestWhenInsertExcludedCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.ExcludedCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_EXCLUDED_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long excluded criteria should return error")
}

func TestWhenInsertAlertWithInvalidThreshold_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.Threshold = 2000000

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_THRESHOLD, err.Error(), "Create alert with greater than max valued threshold should return error")
}

func TestWhenCriteriaIsLongerThanMaxWithComma_ShouldReturnOk(t *testing.T) {
	alert := SampleAlert
	alert.ExcludedCriteria = "01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789,"

	err := alert.validate()
	assert.Equal(t, err, nil, "Create alert with criteria longer than max with comma should be ok")
}
