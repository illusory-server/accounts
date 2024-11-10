package app

import "github.com/pkg/errors"

var (
	ErrEnvNotFound       = errors.New("environment not found")
	ErrIncorrectEnvValue = errors.New("incorrect environment value")
)
