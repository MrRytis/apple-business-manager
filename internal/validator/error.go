package validator

import "fmt"

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
	Code   string `json:"code"`
	Value  string `json:"value"`
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed - Field: %s, Reason: %s, Code: %s, Value: %s", ve.Field, ve.Reason, ve.Code, ve.Value)
}

func (ve *ValidationErrors) Error() string {
	return fmt.Sprintf("Validation failed: %d errors", len(ve.Errors))
}
