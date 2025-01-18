package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asztemborski/syncro/internal/api/handler"
	"github.com/asztemborski/syncro/internal/api/middleware"
	"github.com/asztemborski/syncro/internal/core"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler interface {
	Register(e *echo.Echo)
}

type Middleware interface {
	Register(e *echo.Echo)
}

type Server struct {
	app  *core.App
	http *http.Server
	echo *echo.Echo
}

func NewServer(app *core.App) *Server {
	echo := echo.New()
	echo.Validator = handler.NewRequestValidator()
	echo.HTTPErrorHandler = handler.NewErrorHandler(app).HandleError

	http := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config().Http.Port),
		Handler:      echo,
		IdleTimeout:  app.Config().Http.IdleTimeout,
		ReadTimeout:  app.Config().Http.ReadTimeout,
		WriteTimeout: app.Config().Http.WriteTimeout,
	}

	server := &Server{
		app:  app,
		http: http,
		echo: echo,
	}

	server.RegisterHandlers(
		handler.NewHealthHandler(app),
		handler.NewAccountHandler(app),
	)

	server.RegisterMiddlewares(
		middleware.NewLoggerMiddleware(app),
	)

	return server
}

func (s *Server) RegisterHandlers(handlers ...Handler) {
	for _, handler := range handlers {
		handler.Register(s.echo)
	}
}

func (s *Server) RegisterMiddlewares(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		middleware.Register(s.echo)
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		s.app.Logger().Info("starting server", zap.String("addr", s.http.Addr))
		if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.app.Logger().Error("server start failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return s.Shutdown(ctx)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.app.Logger().Info("shutting down server")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.app.Logger().Error("server shutdown failed", zap.Error(err))
		return err
	}
	return nil
}
