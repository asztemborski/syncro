package model_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/asztemborski/syncro/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewAppErr(t *testing.T) {
	t.Run("should create AppErr with default values when code and message are empty", func(t *testing.T) {
		err := model.NewAppErr("", "")
		assert.Equal(t, model.DefaultAppErrCode, err.Code)
		assert.Equal(t, model.DefaultAppErrMessage, err.Message)
		assert.Equal(t, http.StatusBadRequest, err.Status)
		assert.Nil(t, err.Details)
	})

	t.Run("should create AppErr with provided code and message", func(t *testing.T) {
		err := model.NewAppErr("custom.code", "custom message")
		assert.Equal(t, "custom.code", err.Code)
		assert.Equal(t, "custom message", err.Message)
		assert.Equal(t, http.StatusBadRequest, err.Status)
		assert.Nil(t, err.Details)
	})
}

func TestAppErrError(t *testing.T) {
	t.Run("should return correct error string for AppErr with no wrapped error", func(t *testing.T) {
		err := model.NewAppErr("test.code", "test message")
		assert.Equal(t, "test.code: test message", err.Error())
	})

	t.Run("should include wrapped error in the error string", func(t *testing.T) {
		wrappedErr := errors.New("wrapped error")
		err := model.NewAppErr("test.code", "test message").Wrap(wrappedErr)
		assert.Equal(t, "test.code: test message: wrapped error", err.Error())
	})

	t.Run("should handle missing code and message gracefully", func(t *testing.T) {
		err := model.NewAppErr("", "").Wrap(errors.New("wrapped error"))
		assert.Equal(t, "core.error: core error has occurred: wrapped error", err.Error())
	})
}

func TestAppErrWithStatus(t *testing.T) {
	newStatusCode := http.StatusForbidden
	err := model.NewAppErr("test.code", "test message").WithStatus(newStatusCode)
	assert.Equal(t, err.Status, newStatusCode)
}

func TestAppErrWithDetails(t *testing.T) {
	details := []model.AppErrDetail{
		{Path: "field1", Message: "invalid value"},
		{Path: "field2", Message: "missing value"},
	}
	err := model.NewAppErr("test.code", "test message").WithDetails(details...)

	assert.Equal(t, details, err.Details)
	assert.Len(t, err.Details, 2)
}

func TestAppErrWrap(t *testing.T) {
	wrappedErr := errors.New("wrapped error")
	err := model.NewAppErr("test.code", "test message").Wrap(wrappedErr)
	assert.Equal(t, wrappedErr, err.Unwrap())
}

func TestAppErrUnwrap(t *testing.T) {
	t.Run("should return nil if no error is wrapped", func(t *testing.T) {
		err := model.NewAppErr("test.code", "test message")
		assert.Nil(t, err.Unwrap())
	})

	t.Run("should return the wrapped error", func(t *testing.T) {
		wrappedErr := errors.New("wrapped error")
		err := model.NewAppErr("test.code", "test message").Wrap(wrappedErr)
		assert.Equal(t, wrappedErr, err.Unwrap())
	})
}
