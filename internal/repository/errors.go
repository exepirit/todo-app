package repository

import "errors"

var (
	ErrNotFound = errors.New("object not found in repository")
	ErrExists   = errors.New("entity with same ID already exists")
)
