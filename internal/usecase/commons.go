package usecase

import "errors"

var (
	ErrNoResultRows = errors.New("no rows in result set")
	ErrTodoNotFound = errors.New("the todo with the provided ID was not found")
)
