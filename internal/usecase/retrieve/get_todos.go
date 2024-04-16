package retrieve

import "todoList/internal/entity"

type GetTodosOutput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

type GetTodos struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewGetTodos(todoRepository entity.TodoRepositoryInterface) *GetTodos {
	return &GetTodos{todoRepository: todoRepository}
}

func (u *GetTodos) Execute() ([]*GetTodosOutput, error) {
	todos, err := u.todoRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return toTodosOutput(todos), nil
}

func toTodosOutput(todos []*entity.Todo) []*GetTodosOutput {
	output := make([]*GetTodosOutput, 0)
	for _, todo := range todos {
		output = append(output, &GetTodosOutput{
			ID:          todo.ID.String(),
			Title:       todo.Title,
			Description: todo.Description,
			IsCompleted: todo.IsCompleted,
		})
	}
	return output
}
