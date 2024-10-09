package users

import (
	"context"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/config"
	"log/slog"
)

type Repository struct {
	cfg    *config.Config
	logger *slog.Logger
}

func NewRepository(cfg *config.Config, logger *slog.Logger) *Repository {
	return &Repository{
		cfg:    cfg,
		logger: logger,
	}
}

func (r *Repository) GetUser(ctx context.Context, userID string) (*User, error) {
	// todo: implement
	return nil, nil
}
