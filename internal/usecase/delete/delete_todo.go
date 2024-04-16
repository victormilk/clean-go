package delete

import (
	"todoList/internal/entity"
	"todoList/internal/usecase"
	"todoList/utils"
)

type DeleteTodoInput struct {
	ID string `json:"id"`
}

type DeleteTodo struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewDeleteTodo(todoRepository entity.TodoRepositoryInterface) *DeleteTodo {
	return &DeleteTodo{todoRepository: todoRepository}
}

func (u *DeleteTodo) Execute(input DeleteTodoInput) error {
	id, err := utils.ParseToUUID(input.ID)
	if err != nil {
		return err
	}

	exists, err := u.todoRepository.ExistsByID(id)
	if err != nil {
		return err
	}

	if !exists {
		return usecase.ErrTodoNotFound
	}

	return u.todoRepository.DeleteByID(id)
}
