package main

// Alert is
type Alert struct {
	id                 int
	name               string
	requiredCriteria   []string
	niceToHaveCriteria []string
	excludedCriteria   []string
	threshold          int
	ownerid            int64
}
