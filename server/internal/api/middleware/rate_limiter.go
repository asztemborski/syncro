package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/asztemborski/syncro/internal/core"
	"github.com/asztemborski/syncro/internal/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

var ErrRateLimitExceed = model.NewAppErr("core.rate_limited", "rate limit exceed").
	WithStatus(http.StatusTooManyRequests)

type Visitor struct {
	*rate.Limiter
	lastSeen time.Time
}

type RateLimiterMiddleware struct {
	mu       sync.Mutex
	app      *core.App
	visitors map[string]*Visitor
}

func NewRateLimiterMiddleware(app *core.App) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		app:      app,
		visitors: map[string]*Visitor{},
	}
}

func (m *RateLimiterMiddleware) Register(e *echo.Echo) {
	if !m.app.Config().Http.RateLimiter.Enabled {
		return
	}

	e.Use(m.RateLimit)
	go func() {
		for range time.Tick(time.Minute) {
			m.RemoveOldVisitors()
		}
	}()
}

func (m *RateLimiterMiddleware) RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		maxRps := m.app.Config().Http.RateLimiter.MaximumRPS
		maxBurst := m.app.Config().Http.RateLimiter.MaximumBurst

		m.mu.Lock()
		if _, exist := m.visitors[ip]; !exist {
			m.visitors[ip] = &Visitor{
				Limiter: rate.NewLimiter(rate.Limit(maxRps), maxBurst),
			}
		}

		m.visitors[ip].lastSeen = time.Now()
		if !m.visitors[ip].Allow() {
			m.mu.Unlock()
			return ErrRateLimitExceed
		}

		m.mu.Unlock()
		return next(c)
	}
}

func (m *RateLimiterMiddleware) RemoveOldVisitors() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for ip, visitor := range m.visitors {
		if time.Since(visitor.lastSeen) > time.Minute {
			delete(m.visitors, ip)
		}
	}
}
