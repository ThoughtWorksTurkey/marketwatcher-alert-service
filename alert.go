package main

import (
	"regexp"
	"strings"
	"github.com/gocql/gocql"
)

// Constants for Alert entity
const (
	ACTIVE = 1
	DEACTIVE = 2
	MAX_LENGTH_FOR_CRITERIA = 140
	MAX_LENGTH_FOR_NAME = 32
	MAX_THRESHOLD_VALUE = 1000000
)

var alphanumericRegex = regexp.MustCompile("^[a-zA-Z0-9_\\s]*$")

// Alert is the primary entity of this microservice
type Alert struct {
	ID                 gocql.UUID    `cql:"id" json:"id"`
	OwnerID            int           `cql:"owner_id" json:"owner_id"`
	Name               string        `cql:"name" json:"name"`
	RequiredCriteria   string 	     `cql:"required_criteria" json:"required_criteria"`
	NiceToHaveCriteria string 		 `cql:"nice_to_have_criteria" json:"nice_to_have_criteria"`
	ExcludedCriteria   string 		 `cql:"excluded_criteria" json:"excluded_criteria"`
	Threshold          int    		 `cql:"threshold" json:"threshold"`
	Status             int    		 `cql:"status" json:"status"`
}

func (a *Alert) validate() bool {
	return a.OwnerID > 0 &&
	validateName(a) &&
	validateRequiredCriteria(a) &&
	validateOptionalCriteria(a.NiceToHaveCriteria) &&
	validateOptionalCriteria(a.ExcludedCriteria) &&
	validateThreshold(a.Threshold)
}

func validateName(a *Alert) bool {
	return a.Name != "" && len(a.Name) <= MAX_LENGTH_FOR_NAME &&
	alphanumericRegex.MatchString(a.Name)
}

func validateRequiredCriteria(a *Alert) bool {
	if a.RequiredCriteria == "" {
		return false
	}
	if len(a.RequiredCriteria) > MAX_LENGTH_FOR_CRITERIA {
		return false
	}
	return validateCriteriaPhrases(a.RequiredCriteria)
}

func validateOptionalCriteria(a string) bool {
	if a == "" {
		return true
	}
	if len(a) > MAX_LENGTH_FOR_CRITERIA {
		return false
	}
	return validateCriteriaPhrases(a)
}

func validateCriteriaPhrases(a string) bool {
	phrases := strings.Split(a, ",")
	var b = true
	for _, phrase := range phrases {
		b = alphanumericRegex.MatchString(phrase)
		if !b {
			break
		}
	}
	return b
}

func validateThreshold(a int) bool {
	return a <= MAX_THRESHOLD_VALUE && a > 0
}

func GenerateAlertId() gocql.UUID{
	return gocql.TimeUUID()
}