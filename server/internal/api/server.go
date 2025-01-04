package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asztemborski/syncro/internal/app"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Register(e *echo.Echo)
}

type Middleware interface {
	Register(e *echo.Echo)
}

type Server struct {
	App  *app.App
	http *http.Server
	echo *echo.Echo
}

func NewServer(app *app.App) *Server {
	echo := echo.New()

	http := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config().Http.Port),
		Handler:      echo,
		IdleTimeout:  app.Config().Http.IdleTimeout,
		ReadTimeout:  app.Config().Http.ReadTimeout,
		WriteTimeout: app.Config().Http.WriteTimeout,
	}

	server := &Server{
		App:  app,
		http: http,
		echo: echo,
	}

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
	go s.http.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return s.Shutdown(ctx)
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return s.http.Shutdown(ctx)
}
