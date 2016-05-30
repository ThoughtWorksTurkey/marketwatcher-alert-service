package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sr = FakeRepository{}

func TestWhenIInsertInvalidAlert_ShouldReturnError(t *testing.T) {
	_, err := saveAlert(sr, Alert{id: -1})
	assert.EqualError(t, err, "Validation failed")
}

func TestWhenIInsertInvalidNamedAlert_ShouldReturnError(t *testing.T) {
	_, err := saveAlert(sr, Alert{id: 1, name: ""})
	assert.EqualError(t, err, "Validation failed")
}

func TestWhenIUpdateAlertWithInvalidRequiredCriteria_ShouldReturnError(t *testing.T) {
	alert := findAlertByID(sr, 1)
	alert.requiredCriteria = nil
	_, err := saveAlert(sr, alert)
	assert.EqualError(t, err, "Validation failed")
}
