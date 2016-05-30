package main

import (
	"github.com/stretchr/testify/assert"
	//"log"
	"testing"
)

var alertMap map[int]Alert

func generateSampleAlertMap() {
	alertMap = make(map[int]Alert)
	alertMap[1] = getSampleAlert()
}

func getSampleAlert() Alert {
	return Alert{
		id:                 1,
		name:               "Test Alert",
		requiredCriteria:   []string{"ali", "veli"},
		niceToHaveCriteria: []string{"ali", "veli"},
		excludedCriteria:   []string{"ali", "veli"},
		threshold:          1000,
		ownerid:            1,
	}
}

func TestWhenIPassID_1_FindAlertByID_ShoudReturnSucessful(t *testing.T) {
	var expectedValue = getSampleAlert()
	var actualValue = findAlertByID(1)

	assert.Equal(t, expectedValue, actualValue, "Must return correct Alert")
}

func TestWhenIPassID_666_FindAlertByID_ShouldReturnNothing(t *testing.T) {
	assert.Equal(t, Alert{id: 0}, findAlertByID(666), "Must return nothing")
}

func findAlertByID(id int) Alert {
	generateSampleAlertMap()
	v, _ := alertMap[id]
	return v
}
