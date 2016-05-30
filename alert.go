package main

// Alert is the primary entity of this microservice
type Alert struct {
	id                 int      `cql:"id"`
	name               string   `cql:"name"`
	requiredCriteria   []string `cql:"required_criteria"`
	niceToHaveCriteria []string `cql:"nice_to_have_criteria"`
	excludedCriteria   []string `cql:"excluded_criteria"`
	threshold          int      `cql:"threshold"`
	ownerID            int64    `cql:"owner_id"`
}

func (a *Alert) validate() bool {
	return a.id > 0 &&
		a.name != "" &&
		a.requiredCriteria != nil && len(a.requiredCriteria) > 0 &&
		a.threshold > 0 && a.ownerID > 0
}

var sampleAlert = Alert{
	id:                 1,
	name:               "Test Alert",
	requiredCriteria:   []string{"ali", "veli"},
	niceToHaveCriteria: []string{"ali", "veli"},
	excludedCriteria:   []string{"ali", "veli"},
	threshold:          1000,
	ownerID:            1,
}
