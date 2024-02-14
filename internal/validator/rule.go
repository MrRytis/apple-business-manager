package validator

import (
	"fmt"
	"time"
)

func ValidateString(val string, field string, required bool, min int, max int) []ValidationError {
	var errors []ValidationError

	if required && val == "" {
		errors = append(errors, ValidationError{
			Field:  field,
			Reason: "Field is required",
			Code:   "required",
			Value:  val,
		})
	}

	if min > 0 && len(val) < min {
		errors = append(errors, ValidationError{
			Field:  field,
			Reason: fmt.Sprintf("Must be at least %d characters long", min),
			Code:   "min",
			Value:  val,
		})
	}

	if max > 0 && len(val) > max {
		errors = append(errors, ValidationError{
			Field:  field,
			Reason: fmt.Sprintf("Must be at most %d characters long", max),
			Code:   "max",
			Value:  val,
		})
	}

	return errors
}

func ValidateDateGreater(val time.Time, field string, time time.Time) []ValidationError {
	var errors []ValidationError

	if val.Before(time) {
		errors = append(errors, ValidationError{
			Field:  field,
			Reason: fmt.Sprintf("Date must be greater than %s", time.Format("2006-01-02T15:04:05Z")),
			Code:   "date-greater",
			Value:  val.Format("2006-01-02T15:04:05Z"),
		})
	}

	return errors
}
