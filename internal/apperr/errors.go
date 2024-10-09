package apperr

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("entity already exists")
)

type HttpError struct {
	StatusCode int
	Message    string `json:"error_message"`
}

func (e *HttpError) Error() string {
	return e.Message
}
