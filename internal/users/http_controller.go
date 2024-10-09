package users

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/apperr"
	"github.com/alisher-matkurbanov/ecommerce-user-service/internal/config"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type (
	Controller struct {
		cfg         *config.Config
		logger      *slog.Logger
		userService userService
	}

	userService interface {
		CreateUser(ctx context.Context, user *UserCreate) (*User, error)
		GetUser(ctx context.Context, userID UserID) (*User, error)
		UpdateUser(ctx context.Context, user *UserUpdate) (*User, error)
	}
)

func NewHttpController(
	cfg *config.Config,
	logger *slog.Logger,
	userService userService,
) *Controller {
	return &Controller{
		cfg:         cfg,
		logger:      logger,
		userService: userService,
	}
}

func writeError(httpErr *apperr.HttpError, w http.ResponseWriter, logger *slog.Logger) {
	data, err := json.Marshal(httpErr)
	if err != nil {
		logger.Error(err.Error())
		_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
	w.WriteHeader(httpErr.StatusCode)
	return
}

func handleError(err error, w http.ResponseWriter, logger *slog.Logger) {
	logger.Error(err.Error())

	var httpErr *apperr.HttpError
	if errors.As(err, &httpErr) {
		writeError(httpErr, w, logger)
		return
	}

	if errors.Is(err, apperr.ErrAlreadyExists) {
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
		writeError(httpErr, w, logger)
		return
	}

	httpErr.Message = err.Error()
	httpErr.StatusCode = http.StatusInternalServerError
	writeError(httpErr, w, logger)
	return
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userCreateRequest := UserCreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userCreateRequest); err != nil {
		handleError(err, w, c.logger)
		return
	}

	userCreate := fromUserCreateRequestToUserCreate(&userCreateRequest)
	createdUser, err := c.userService.CreateUser(ctx, userCreate)
	if err != nil {
		handleError(err, w, c.logger)
		return
	}

	response := fromUserToUserResponse(createdUser)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		handleError(err, w, c.logger)
		return
	}
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := chi.URLParam(r, "user_id")
	user, err := c.userService.GetUser(ctx, UserID(userID))
	if err != nil {
		handleError(err, w, c.logger)
		return
	}

	response := fromUserToUserResponse(user)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		handleError(err, w, c.logger)
		return
	}
}

func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userUpdateRequest := UserUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userUpdateRequest); err != nil {
		handleError(err, w, c.logger)
		return
	}

	user := fromUserUpdateRequestToUserUpdate(&userUpdateRequest)
	createdUser, err := c.userService.UpdateUser(ctx, user)
	if err != nil {
		handleError(err, w, c.logger)
		return
	}

	response := fromUserToUserResponse(createdUser)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		handleError(err, w, c.logger)
		return
	}
}
