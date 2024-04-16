package complete

import (
	"errors"
	"todoList/internal/entity"
	"todoList/internal/usecase"
	"todoList/utils"
)

type CompleteTodoInput struct {
	ID string `json:"id"`
}

type CompleteTodo struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewCompleteTodo(todoRepository entity.TodoRepositoryInterface) *CompleteTodo {
	return &CompleteTodo{todoRepository: todoRepository}
}

func (u *CompleteTodo) Execute(input CompleteTodoInput) error {
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

	if err := todo.Complete(); err != nil {
		return err
	}

	return u.todoRepository.Complete(todo)
}
