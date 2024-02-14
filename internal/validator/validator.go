package validator

import (
	"github.com/pkg/errors"
)

type Validator interface {
	Validate() ValidationErrors
}

type CustomRequestValidator struct {
}

func New() *CustomRequestValidator {
	return &CustomRequestValidator{}
}

func (cv *CustomRequestValidator) Validate(i interface{}) error {
	if v, ok := i.(Validator); ok {
		err := v.Validate()

		if len(err.Errors) > 0 {
			return &err
		}

		return nil
	}

	return errors.New("To validate a request, it must implement the Validator interface.")
}
