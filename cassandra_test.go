// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gocql/gocql"
)

var generatedId gocql.UUID

func TestWhenIUpsert_ShoudReturnSucessful(t *testing.T) {
	inserted, err := save(SampleAlert)
	generatedId = inserted.ID
	assert.Equal(t, SampleAlert.Name, inserted.Name, "Upsert should return OK")
	assert.NoError(t, err)
}

func TestFindById_ShouldReturnSuccessful(t *testing.T){
 	alert,err := find(generatedId)
	assert.Equal(t, alert.Name, "Test Alert", "Find should return OK")
	assert.NoError(t, err)
}