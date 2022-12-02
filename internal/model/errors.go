package model

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrStorage  = errors.New("storage error")
)
