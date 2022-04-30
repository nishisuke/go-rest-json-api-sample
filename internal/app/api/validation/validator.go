package validation

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	Validation struct {
		validator *validator.Validate
	}
)

func NewValidation() *Validation {
	return &Validation{validator: validator.New()}
}

func (cv *Validation) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
