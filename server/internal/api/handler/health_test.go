package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/asztemborski/syncro/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestHealthCheckHandler(t *testing.T) {
	tests := []struct {
		name            string
		debug           bool
		expectedEnv     string
		expectedVersion string
	}{
		{
			name:            "Production Environment",
			expectedEnv:     "production",
			debug:           false,
			expectedVersion: "1.0.0",
		},
		{
			name:            "Development Environment",
			expectedEnv:     "development",
			debug:           true,
			expectedVersion: "1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testConfig := &config.Configuration{
				App: config.AppConfig{
					Environment: tt.expectedEnv,
					Version:     tt.expectedVersion,
				},
			}

			testApp := app.NewApp(testConfig, zaptest.NewLogger(t), &app.AccountService{})

			e := echo.New()

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			handler := NewHealthHandler(testApp)

			err := handler.healthCheck(c)
			require.NoError(t, err)

			assert.Equal(t, http.StatusOK, rec.Code)

			var response map[string]string
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, "available", response["status"])
			assert.Equal(t, tt.expectedEnv, response["environment"])
			assert.Equal(t, tt.expectedVersion, response["version"])
		})
	}
}
