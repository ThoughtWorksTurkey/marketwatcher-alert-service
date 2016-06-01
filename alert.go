package main

import "time"

// Status
const (
	Active   = 1
	Deactive = 2
)

// Alert is the primary entity of this microservice
type Alert struct {
	ID                 int
	OwnerID            int       `cql:"owner_id"`
	Name               string    `cql:"name"`
	RequiredCriteria   string    `cql:"required_criteria"`
	NiceToHaveCriteria string    `cql:"nice_to_have_criteria"`
	ExcludedCriteria   string    `cql:"excluded_criteria"`
	Threshold          int       `cql:"threshold"`
	Status             int       `cql:"status"`
	UpdateDate         time.Time `cql:"update_date"`
}

func (a *Alert) validate() bool {
	return a.OwnerID > 0 &&
		a.Name != "" &&
		a.RequiredCriteria != "" &&
		a.Threshold > 0
}

var sampleAlert = Alert{
	ID:                 1,
	OwnerID:            1,
	Name:               "Test Alert",
	RequiredCriteria:   "ThoughtWorks",
	NiceToHaveCriteria: "good,best office",
	ExcludedCriteria:   "bad",
	Threshold:          1000,
	Status:             Active,
	UpdateDate:         time.Now(),
}
