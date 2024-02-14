package exception

import (
	"github.com/MrRytis/apple-business-manager/internal/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPError struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

type Error struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
	Code   string `json:"code"`
	Value  string `json:"value"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"
	errors := []Error{}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)

		c.Logger().Error(err)
		returnError(c, code, message, errors)

		return
	}

	if v, ok := err.(*validator.ValidationErrors); ok {
		code = http.StatusUnprocessableEntity
		message = "Validation failed"
		for _, e := range v.Errors {
			errors = append(errors, Error{
				Field:  e.Field,
				Reason: e.Reason,
				Code:   e.Code,
				Value:  e.Value,
			})
		}

		c.Logger().Info(v)
		returnError(c, code, message, errors)

		return
	}

	c.Logger().Error(err)
	returnError(c, code, message, errors)

	return
}

func returnError(c echo.Context, code int, message string, errors []Error) {
	httpErr := HTTPError{
		Code:    code,
		Message: message,
		Errors:  errors,
	}

	c.JSON(code, httpErr)
}
