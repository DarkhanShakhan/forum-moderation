package errors

import "errors"

var (
	ErrPostNotFound     = errors.New("post not found")
	ErrCategoryNotFound = errors.New("category not found")
)
