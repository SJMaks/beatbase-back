package apperrs

import "errors"

var (
	EnvNotFound = errors.New("env not found")
)
