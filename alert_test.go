package main

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWhenIInsertValidAlert_ShouldReturnOK(t *testing.T) {
	err := SampleAlert.validate()
	assert.Equal(t, nil, err, "Create should not return OK for valid return")
}

func TestWhenIInsertAlertWithoutName_ShouldReturnError(t *testing.T) {
	alertWithoutName := SampleAlert
	alertWithoutName.Name = ""

	err := alertWithoutName.validate()
	assert.Equal(t, VALIDATION_MESSAGE_NAME_EMPTY, err.Error(), "Create alert without name should return error")
}

func TestWhenIInsertAlertWithNameLengthIsMoreThanMax_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.Name = "ayse jkajshdkjashdjsahd kashdjashdkjahsdkjhaskjdhaksjhdkjashd jashdkjahsdkjahsdjhsd"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_NAME_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_NAME), err.Error(), "Create alert with name length is more than max should return error")
}

func TestWhenIInsertAlertWithoutOwnerID_ShouldReturnError(t *testing.T) {
	alertWithoutOwnerID := SampleAlert
	alertWithoutOwnerID.OwnerID = -2

	err := alertWithoutOwnerID.validate()
	assert.Equal(t, VALIDATION_MESSAGE_OWNER_ID, err.Error(), "Create alert without owner id should return error")
}

func TestWhenIInsertTurkishCharacterForCriteria_ShouldReturnOk(t *testing.T) {
	alert := SampleAlert
	alert.RequiredCriteria = "ayçe çç öö ğ ü ı şşşşşşş"
	err := alert.validate()
	assert.Equal(t, nil, err, "Create alert with turkish character for criteria should not return OK")
}

func TestWhenIInsertNonAlphanumericCharacterForCriteria_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.RequiredCriteria = ">>> < | ~~~ ]"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_CRITERIA_PHRASES_ALPHANUMERIC, err.Error(), "Create alert with non alphanumeric creiteria should return error")
}

func TestWhenIInsertRequiredCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.RequiredCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_REQUIRED_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long required criteria should return error")
}

func TestWhenIInsertNiceToHaveCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.NiceToHaveCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_NICE_TO_HAVE_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long nice-to-have criteria should return error")
}

func TestWhenIInsertExcludedCriteriaLongerThan140_ShouldReturnError(t *testing.T) {
	alert := SampleAlert
	alert.ExcludedCriteria = "aaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqwaaaaaaaqw"

	err := alert.validate()
	assert.Equal(t, VALIDATION_MESSAGE_EXCLUDED_CRITERIA_LENGTH+strconv.Itoa(MAX_LENGTH_FOR_CRITERIA), err.Error(), "Create alert with long excluded criteria should return error")
}

func TestWhenIInsertAlertWithInvalidThreshold_ShouldReturnError(t *testing.T) {
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
