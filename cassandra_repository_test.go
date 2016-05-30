package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cr = CassandraRepository{}

func TestWhenIFindWithID1_ShoudReturnSucessful(t *testing.T) {
	actualValue, _ := cr.find(1)
	assert.Equal(t, sampleAlert, actualValue, "Must return correct Alert")
}

/*
func TestWhenIUpsertAlert_ShouldReturnOK(t *testing.T) {
	insertedAlert, err := cr.upsert(sampleAlert)
	assert.Equal(t, sampleAlert, insertedAlert, "Insert alert to map should be successful")
	assert.Nil(t, err, "Error must be nil")
}*/
