package service

import "errors"

// ошибки
var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidID     = errors.New("invalid ID")
	ErrEmptyName     = errors.New("empty name")
	ErrInvalidName   = errors.New("invalid name characters")
	ErrDuplicateID   = errors.New("duplicate ID")
	ErrDuplicateName = errors.New("duplicate name")
)
