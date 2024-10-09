package users

import (
	"context"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/config"
	"log/slog"
)

type Service struct {
	cfg      *config.Config
	logger   *slog.Logger
	userRepo *Repository
}

func NewService(cfg *config.Config, logger *slog.Logger, userRepo *Repository) *Service {
	return &Service{
		cfg:      cfg,
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *Service) CreateUser(ctx context.Context, user *UserCreate) (*User, error) {
	return &User{}, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *UserUpdate) (*User, error) {
	return &User{}, nil
}

func (s *Service) GetUser(ctx context.Context, userID UserID) (*User, error) {
	return &User{}, nil
}
