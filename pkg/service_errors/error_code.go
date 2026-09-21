package service_errors

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrNotFound           = errors.New("not found")
	ErrConflict           = errors.New("conflict")
	ErrUnavailable        = errors.New("service unavailable")
	ErrValidation         = errors.New("validation failed")
)

const (
	EmailExists      = "Email exists"
	UsernameExists   = "Username exists"
	RecordNotFound   = "record not found"
	PermissionDenied = "Permission denied"
	TokenRequired    = "token required"
	TokenExpired     = "token expired"
	TokenInvalid     = "token invalid"
)
