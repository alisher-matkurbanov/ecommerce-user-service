package http

import (
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/config"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	cfg        *config.Config
	httpServer *http.Server
}

func NewServer(cfg *config.Config, router *chi.Mux) *Server {
	httpServer := &http.Server{
		Handler:                      router,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}
	return &Server{
		cfg:        cfg,
		httpServer: httpServer,
	}
}

func (s *Server) ServeHTTP() error {
	return s.httpServer.ListenAndServe()
}
