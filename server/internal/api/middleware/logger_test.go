package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLoggerMiddleware(t *testing.T) {
	observedZapCore, logs := observer.New(zap.InfoLevel)
	testLogger := zap.New(observedZapCore)

	testApp := app.NewApp(&config.Configuration{}, testLogger, &app.AccountService{})

	middleware := NewLoggerMiddleware(testApp)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("User-Agent", "test-agent")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test response")
	}

	err := middleware.LogRequest(testHandler)(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Equal(t, 1, logs.Len())
	logEntry := logs.All()[0]

	assert.Equal(t, "http request", logEntry.Message)

	fields := make(map[string]any)
	for _, field := range logEntry.Context {
		switch field.Type {
		case zapcore.StringType:
			fields[field.Key] = field.String
		case zapcore.Int64Type, zapcore.Int32Type:
			fields[field.Key] = field.Integer
		}
	}

	expectedFields := map[string]bool{
		"remote ip":  true,
		"host":       true,
		"method":     true,
		"uri":        true,
		"user agent": true,
		"status":     true,
		"latency":    true,
		"bytes out":  true,
	}

	for field := range expectedFields {
		assert.Contains(t, fields, field)
	}

	assert.Equal(t, "GET", fields["method"])
	assert.Equal(t, "/test", fields["uri"])
	assert.Equal(t, "test-agent", fields["user agent"])
	assert.Equal(t, int64(200), fields["status"])
}
