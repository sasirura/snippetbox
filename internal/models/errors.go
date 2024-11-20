package models

import "errors"

var (
	ErrNoRecord           = errors.New("Models: no matching records found")
	ErrInvalidCredentials = errors.New("Models: invalid credentials")
	ErrDuplicateEmail     = errors.New("Models: duplicate email")
)
