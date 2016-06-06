// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenIUpsert_ShoudReturnSucessful(t *testing.T) {
	inserted, err := upsert(SampleAlert)
	assert.Equal(t, inserted, SampleAlert, "Upsert should return OK")
	assert.NoError(t, err)
}