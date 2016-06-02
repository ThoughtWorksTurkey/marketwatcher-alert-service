// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenIUpsert_ShoudReturnSucessful(t *testing.T) {
	inserted, err := UpsertAlertCassandra(SampleAlert)
	assert.Equal(t, inserted, SampleAlert, "Upsert should return OK")
	assert.NoError(t, err)
}

/*
func TestWhenIFindWithID1_ShoudReturnSucessful(t *testing.T) {
	actualValue, err := Find(1)
	assert.Equal(t, SampleAlert, actualValue, "Must return correct Alert")
	assert.NoError(t, err)
}*/
