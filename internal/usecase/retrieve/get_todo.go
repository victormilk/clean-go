package retrieve

import (
	"time"
	"todoList/internal/entity"
	"todoList/utils"
)

type GetTodoInput struct {
	ID string `json:"id"`
}

type GetTodoOutput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetTodo struct {
	todoRepository entity.TodoRepositoryInterface
}

func NewGetTodo(todoRepository entity.TodoRepositoryInterface) *GetTodo {
	return &GetTodo{todoRepository: todoRepository}
}

func (u *GetTodo) Execute(input GetTodoInput) (*GetTodoOutput, error) {
	id, err := utils.ParseToUUID(input.ID)
	if err != nil {
		return nil, err
	}

	todo, err := u.todoRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return toTodoOutput(todo), nil
}

func toTodoOutput(todo *entity.Todo) *GetTodoOutput {
	return &GetTodoOutput{
		ID:          todo.ID.String(),
		Title:       todo.Title,
		Description: todo.Description,
		IsCompleted: todo.IsCompleted,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}
