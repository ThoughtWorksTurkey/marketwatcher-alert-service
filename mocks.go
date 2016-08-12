package main

import "github.com/gocql/gocql"

var SampleAlert = Alert{
	ID:                 GenerateAlertId(),
	OwnerID:            1,
	Name:               "Test Alert",
	RequiredCriteria:   "TW,ThoughtWorks,Thought Works,Thoughtworks",
	NiceToHaveCriteria: "good,best office",
	ExcludedCriteria:   "bad,sucks,not good enough",
	Threshold:          1000,
	Status:             ACTIVE,
}

func MockSave(a Alert) (Alert, error) {
	return a, nil
}

func MockFind(id gocql.UUID) (Alert, error) {
	return SampleAlert, nil
}

func MockTriggerIngestion(a Alert) error {
	return nil
}
