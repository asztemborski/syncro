package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	validator := validator.New()
	return &RequestValidator{validator: validator}
}

func (rv *RequestValidator) Validate(input any) error {
	if err := rv.validator.Struct(input); err != nil {
		return err
	}
	return nil
}

func retrieveAndValidateRequestBody[T any](c echo.Context) (T, error) {
	var request T
	if err := c.Bind(&request); err != nil {
		return request, err
	}
	if err := c.Validate(&request); err != nil {
		return request, err
	}

	return request, nil
}
