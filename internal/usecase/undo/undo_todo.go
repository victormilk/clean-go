package undo

import (
	"errors"
	"todoList/internal/entity"
	"todoList/internal/usecase"
	"todoList/utils"
)

type UndoTodoInput struct {
	ID string `json:"id"`
}

type UndoTodo struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewUndoTodo(todoRepository entity.TodoRepositoryInterface) *UndoTodo {
	return &UndoTodo{todoRepository: todoRepository}
}

func (u *UndoTodo) Execute(input *UndoTodoInput) error {
	id, err := utils.ParseToUUID(input.ID)
	if err != nil {
		return err
	}

	todo, err := u.todoRepository.GetByID(id)
	if err != nil {
		if errors.Is(err, usecase.ErrNoResultRows) {
			return usecase.ErrTodoNotFound
		}
		return err
	}

	if err := todo.Undo(); err != nil {
		return err
	}

	return u.todoRepository.Undo(todo)
}
