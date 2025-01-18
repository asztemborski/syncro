package model

import (
	"net/http"
	"strings"
)

const (
	DefaultAppErrCode    = "core.error"
	DefaultAppErrMessage = "core error has occurred"
)

type AppErrDetail struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type AppErr struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Details []AppErrDetail `json:"details"`
	wrapped error
}

func NewAppErr(code, message string) *AppErr {
	if code == "" {
		code = DefaultAppErrCode
	}

	if message == "" {
		message = DefaultAppErrMessage
	}

	return &AppErr{
		Code:    code,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func (e *AppErr) Error() string {
	var sb strings.Builder

	if e.Code != "" {
		sb.WriteString(e.Code)
	}

	if e.Message != "" {
		if sb.Len() > 0 {
			sb.WriteString(": ")
		}
		sb.WriteString(e.Message)
	}

	if e.wrapped != nil {
		if sb.Len() > 0 {
			sb.WriteString(": ")
		}
		sb.WriteString(e.wrapped.Error())
	}
	return sb.String()
}

func (e *AppErr) WithDetails(details ...AppErrDetail) *AppErr {
	e.Details = append(e.Details, details...)
	return e
}

func (e *AppErr) WithStatus(status int) *AppErr {
	e.Status = status
	return e
}

func (e *AppErr) Wrap(err error) *AppErr {
	e.wrapped = err
	return e
}

func (e *AppErr) Unwrap() error {
	return e.wrapped
}
