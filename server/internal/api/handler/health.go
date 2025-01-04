package handler

import (
	"net/http"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	app *app.App
}

func NewHealthHandler(app *app.App) *HealthHandler {
	return &HealthHandler{app: app}
}

func (h *HealthHandler) Register(e *echo.Echo) {
	e.GET("/health", h.healthCheck)
}

func (h *HealthHandler) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":      "available",
		"environment": h.app.Config().App.Environment,
		"version":     h.app.Config().App.Version,
	})
}
