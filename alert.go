package main

import (
	"errors"
	"github.com/gocql/gocql"
	"regexp"
	"strconv"
	"strings"
)

// Constants for Alert entity
const (
	ACTIVE                                           = 1
	DEACTIVE                                         = 2
	MAX_LENGTH_FOR_CRITERIA                          = 140
	MAX_LENGTH_FOR_NAME                              = 32
	MAX_THRESHOLD_VALUE                              = 1000000
	VALIDATION_MESSAGE_NAME_EMPTY                    = "Alert name cannot be null"
	VALIDATION_MESSAGE_NAME_LENGTH                   = "Alert name length must be smaller than "
	VALIDATION_MESSAGE_NAME_ALPHANUMERIC             = "Alert name must be all alphanumerical characters"
	VALIDATION_MESSAGE_REQUIRED_CRITERIA_EMPTY       = "Must included criteria cannot be null"
	VALIDATION_MESSAGE_REQUIRED_CRITERIA_LENGTH      = "Required criteria length must be smaller than "
	VALIDATION_MESSAGE_NICE_TO_HAVE_CRITERIA_LENGTH  = "Optional criteria length must be smaller than "
	VALIDATION_MESSAGE_EXCLUDED_CRITERIA_LENGTH      = "Excluded criteria length must be smaller than "
	VALIDATION_MESSAGE_CRITERIA_PHRASES_ALPHANUMERIC = "Criteria phrases must be alphanumerical!"
	VALIDATION_MESSAGE_THRESHOLD                     = "Threshold must be positive number, it cannot be null, zero or negative number!"
	VALIDATION_MESSAGE_OWNER_ID                      = "Validation failed: ownerId!"
)

var alphanumericRegex = regexp.MustCompile("^[a-zA-Z0-9_\\s]*$")
var validationMessage string

// Alert is the primary entity of this microservice
type Alert struct {
	ID                 gocql.UUID `cql:"id" json:"id"`
	OwnerID            int        `cql:"owner_id" json:"owner_id"`
	Name               string     `cql:"name" json:"name"`
	RequiredCriteria   string     `cql:"required_criteria" json:"required_criteria"`
	NiceToHaveCriteria string     `cql:"nice_to_have_criteria" json:"nice_to_have_criteria"`
	ExcludedCriteria   string     `cql:"excluded_criteria" json:"excluded_criteria"`
	Threshold          int        `cql:"threshold" json:"threshold"`
	Status             int        `cql:"status" json:"status"`
}

func (a *Alert) validate() error {
	err := validateName(a)
	if err != nil {
		return err
	}
	err = validateRequiredCriteria(a)
	if err != nil {
		return err
	}
	err = validateNiceToHaveCriteria(a.NiceToHaveCriteria)
	if err != nil {
		return err
	}
	err = validateExcludedCriteria(a.ExcludedCriteria)
	if err != nil {
		return err
	}
	err = validateThreshold(a.Threshold)
	if err != nil {
		return err
	}
	err = validateOwnerID(a.OwnerID)
	if err != nil {
		return err
	}
	return nil
}

func validateName(a *Alert) error {
	if a.Name == "" {
		return errors.New(VALIDATION_MESSAGE_NAME_EMPTY)
	}
	if len(a.Name) > MAX_LENGTH_FOR_NAME {
		return errors.New(VALIDATION_MESSAGE_NAME_LENGTH + strconv.Itoa(MAX_LENGTH_FOR_NAME))
	}
	if !alphanumericRegex.MatchString(a.Name) {
		return errors.New(VALIDATION_MESSAGE_NAME_ALPHANUMERIC)
	}
	return nil
}

func validateRequiredCriteria(a *Alert) error {
	if a.RequiredCriteria == "" {
		return errors.New(VALIDATION_MESSAGE_REQUIRED_CRITERIA_EMPTY)
	}
	if len(a.RequiredCriteria) > MAX_LENGTH_FOR_CRITERIA {
		return errors.New(VALIDATION_MESSAGE_REQUIRED_CRITERIA_LENGTH + strconv.Itoa(MAX_LENGTH_FOR_CRITERIA))
	}
	return validateCriteriaPhrases(a.RequiredCriteria)
}

func validateNiceToHaveCriteria(a string) error {
	if a == "" {
		return nil
	}
	if len(a) > MAX_LENGTH_FOR_CRITERIA {
		return errors.New(VALIDATION_MESSAGE_NICE_TO_HAVE_CRITERIA_LENGTH + strconv.Itoa(MAX_LENGTH_FOR_CRITERIA))
	}
	return validateCriteriaPhrases(a)
}

func validateExcludedCriteria(a string) error {
	if a == "" {
		return nil
	}
	if len(a) > MAX_LENGTH_FOR_CRITERIA {
		return errors.New(VALIDATION_MESSAGE_EXCLUDED_CRITERIA_LENGTH + strconv.Itoa(MAX_LENGTH_FOR_CRITERIA))
	}
	return validateCriteriaPhrases(a)
}

func validateCriteriaPhrases(a string) error {
	phrases := strings.Split(a, ",")
	var b = true
	for _, phrase := range phrases {
		if !(alphanumericRegex.MatchString(phrase)) {
			b = false
			break
		}
	}
	if !b {
		return errors.New(VALIDATION_MESSAGE_CRITERIA_PHRASES_ALPHANUMERIC)
	}
	return nil
}

func validateThreshold(a int) error {
	if a > MAX_THRESHOLD_VALUE || a <= 0 {
		return errors.New(VALIDATION_MESSAGE_THRESHOLD)
	}
	return nil
}

func validateOwnerID(ownerID int) error {
	if ownerID < 0 {
		return errors.New(VALIDATION_MESSAGE_OWNER_ID)
	}
	return nil
}

func GenerateAlertId() gocql.UUID {
	return gocql.TimeUUID()
}
