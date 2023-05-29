package errors

import "errors"

var (
	ErrPostNotFound     = errors.New("post not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrCommentNotFound  = errors.New("comment not found")
)
