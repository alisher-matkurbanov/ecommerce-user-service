package http

import (
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
	"log/slog"
	"time"
)

func (s *Server) InitRoutes(
	r *chi.Mux,
	logger *slog.Logger,
	userController *users.Controller,
) {
	loggerConfig := slogchi.Config{
		WithRequestBody:    true,
		WithResponseBody:   true,
		WithRequestHeader:  true,
		WithResponseHeader: true,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(slogchi.NewWithConfig(logger, loggerConfig))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(s.cfg.HttpTimeout * time.Second))

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userController.CreateUser)
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", userController.GetUser)
			r.Put("/", userController.UpdateUser)
		})
	})
}
