package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sr = FakeRepository{}

func TestWhenIInsertInvalidAlert_ShouldReturnError(t *testing.T) {
	_, err := saveAlert(sr, Alert{RequiredCriteria: ""})
	assert.EqualError(t, err, "Validation failed")
}

func TestWhenIInsertInvalidNamedAlert_ShouldReturnError(t *testing.T) {
	_, err := saveAlert(sr, Alert{Name: ""})
	assert.EqualError(t, err, "Validation failed")
}
