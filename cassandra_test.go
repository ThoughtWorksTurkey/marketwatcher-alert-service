// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenIUpsert_ShouldReturnSucessful(t *testing.T) {
	generatedName := GenerateAlertId().String()
	inserted, err := save(createSampleAlert(generatedName))

	assert.Equal(t, generatedName, inserted.Name, "Upsert should return OK")
	assert.NoError(t, err)
}

func TestWhenSameAlertNameAndOwnerIdIsInsertedShouldReturnError(t *testing.T) {
	_, err := save(createSampleAlert("duplicate-name"))
	_, err = save(createSampleAlert("duplicate-name"))

	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), ALERT_NAME_MUST_BE_UNIQUE_PER_OWNER, "Alert name must be unique per owner")
	}
}

func TestFindById_ShouldReturnSuccessful(t *testing.T) {
	testAlert := createSampleAlert("sample-name")
	save(testAlert)

	foundAlert, err := find(testAlert.ID)
	assert.Equal(t, "sample-name", foundAlert.Name, "Find should return OK")
	assert.NoError(t, err)
}
