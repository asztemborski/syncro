package handler

import (
	"errors"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/asztemborski/syncro/internal/core"
	"github.com/asztemborski/syncro/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	ValidationErrorCode = "thrive.validation"
)

type ErrorHandler struct {
	app *core.App
}

func NewErrorHandler(app *core.App) *ErrorHandler {
	return &ErrorHandler{app: app}
}

func (h *ErrorHandler) HandleError(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	appErr := createAppError(err)
	if respErr := c.JSON(appErr.Status, appErr); respErr != nil {
		h.app.Logger().Error("failed to send error response", zap.Error(respErr))
	}
}

func createAppError(err error) *model.AppErr {
	var appErr *model.AppErr
	var httpErr *echo.HTTPError
	var validationErrs validator.ValidationErrors

	switch {
	case errors.As(err, &appErr):
		return appErr
	case errors.As(err, &httpErr):
		return model.NewAppErr("", firstToLower(http.StatusText(httpErr.Code))).WithStatus(httpErr.Code)
	case errors.As(err, &validationErrs):
		return processValidationErrors(validationErrs)
	}

	return model.NewAppErr("", firstToLower(http.StatusText(http.StatusInternalServerError))).
		WithStatus(http.StatusInternalServerError)
}

func processValidationErrors(validationErrs validator.ValidationErrors) *model.AppErr {
	appError := model.NewAppErr(ValidationErrorCode, "some validation errors have occured")

	for _, err := range validationErrs {
		appError.WithDetails(model.AppErrDetail{
			Path:    firstToLower(err.Field()),
			Message: firstToLower(extractValidationFieldErrorMessage(err)),
		})
	}
	return appError
}

func extractValidationFieldErrorMessage(err validator.FieldError) string {
	parts := strings.Split(err.Error(), "Error:")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}
	return ""
}

func firstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}
