package create

import "todoList/internal/entity"

type CreateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTodo struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewCreateTodo(todoRepository entity.TodoRepositoryInterface) *CreateTodo {
	return &CreateTodo{todoRepository: todoRepository}
}

func (u *CreateTodo) Execute(input CreateTodoInput) error {
	todo, err := entity.NewTodo(input.Title, input.Description)
	if err != nil {
		return err
	}

	return u.todoRepository.Save(todo)
}
