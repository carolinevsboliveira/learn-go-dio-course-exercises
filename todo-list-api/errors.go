package todo_list_api

import "errors"

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrInvalidID    = errors.New("invalid ID format")
	ErrEmptyTitle   = errors.New("title cannot be empty")
)
