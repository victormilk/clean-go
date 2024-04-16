package update

import (
	"errors"
	"todoList/internal/entity"
	"todoList/internal/usecase"
	"todoList/utils"
)

type UpdateTodoInput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodo struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewUpdateTodo(todoRepository entity.TodoRepositoryInterface) *UpdateTodo {
	return &UpdateTodo{todoRepository: todoRepository}
}

func (u *UpdateTodo) Execute(input UpdateTodoInput) error {
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

	if err := todo.Update(input.Title, input.Description); err != nil {
		return err
	}

	return u.todoRepository.Update(todo)
}
