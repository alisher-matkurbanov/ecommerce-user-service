package main

import (
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/config"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/controllers/http"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/infra"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/users"
	"github.com/go-chi/chi/v5"
	"log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := infra.NewLogger()
	userRepo := users.NewRepository(cfg, logger)
	userService := users.NewService(cfg, logger, userRepo)
	httpUserController := users.NewHttpController(cfg, logger, userService)
	router := chi.NewRouter()
	server := http.NewServer(cfg, router)
	server.InitRoutes(router, logger, httpUserController)

	logger.Info("handling requests")
	if err := server.ServeHTTP(); err != nil {
		log.Fatal(err)
	}
}
