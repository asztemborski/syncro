package middleware

import (
	"strconv"
	"time"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	app *app.App
}

func NewLoggerMiddleware(app *app.App) *LoggerMiddleware {
	return &LoggerMiddleware{app: app}
}

func (m *LoggerMiddleware) Register(e *echo.Echo) {
	e.Use(m.LogRequest)
}

func (m *LoggerMiddleware) LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		start := time.Now()

		if err := next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()

		m.app.Logger().Info("new request",
			zap.String("remote ip", c.RealIP()),
			zap.String("host", req.Host),
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
			zap.String("user agent", req.UserAgent()),
			zap.Int("status", res.Status),
			zap.String("latency", stop.Sub(start).String()),
			zap.String("bytes out", strconv.FormatInt(res.Size, 10)),
		)

		return nil
	}
}
