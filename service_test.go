package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenIInsertValidAlert_ShouldReturnOK(t *testing.T) {
	UpsertAlert = UpsertAlertSuccess

	a, err := CreateAlert(sampleAlert)
	assert.Equal(t, a.Name, sampleAlert.Name, "Create should return OK for valid return")
	assert.Nil(t, err, nil, "Create should not return OK for valid return")
}

/*func createAlarm(a Alert){
validateAlert(a);
alert := saveAlert(a);

}*/

/*func TestWhenIInsertInvalidNamedAlert_ShouldReturnError(t *testing.T) {
	_, err := saveAlert(sr, Alert{Name: ""})
	assert.EqualError(t, err, "Validation failed")
}*/
