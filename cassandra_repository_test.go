// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cr = CassandraRepository{}

func TestWhenIUpsert_ShoudReturnSucessful(t *testing.T) {
	inserted, err := cr.upsert(sampleAlert)
	assert.NotEqual(t, inserted.UpdateDate, sampleAlert.UpdateDate "Update date should not be null when inserted")
	assert.NoError(t, err)
}

func TestWhenIFindWithID1_ShoudReturnSucessful(t *testing.T) {
	actualValue, err := cr.find(1)
	assert.Equal(t, sampleAlert, actualValue, "Must return correct Alert")
	assert.NoError(t, err)
}
