package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"todoList/internal/entity"
	db "todoList/sql/sqlc"
)

type TodoRepository struct {
	queries *db.Queries
}

func NewTodoRepository(queries *db.Queries) *TodoRepository {
	return &TodoRepository{queries: queries}
}

func (r *TodoRepository) GetAll() ([]*entity.Todo, error) {
	todos, err := r.queries.GetTodos(context.Background())
	if err != nil {
		return nil, err
	}

	var result []*entity.Todo
	for _, todo := range todos {
		result = append(result, &entity.Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description.String,
			IsCompleted: todo.Completed,
		})
	}

	return result, nil
}

func (r *TodoRepository) GetByID(id uuid.UUID) (*entity.Todo, error) {
	todo, err := r.queries.GetTodo(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &entity.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description.String,
		IsCompleted: todo.Completed,
		CreatedAt:   todo.CreatedAt.Time,
		UpdatedAt:   todo.UpdatedAt.Time,
	}, nil
}

func (r *TodoRepository) Save(todo *entity.Todo) error {
	err := r.queries.CreateTodo(context.Background(), db.CreateTodoParams{
		ID:    todo.ID,
		Title: todo.Title,
		Description: pgtype.Text{
			String: todo.Description,
			Valid:  true,
		},
		Completed: todo.IsCompleted,
		CreatedAt: pgtype.Timestamp{
			Time:  todo.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  todo.UpdatedAt,
			Valid: true,
		},
	})

	return err
}

func (r *TodoRepository) DeleteByID(id uuid.UUID) error {
	err := r.queries.DeleteTodo(context.Background(), id)
	return err
}

func (r *TodoRepository) Complete(todo *entity.Todo) error {
	err := r.queries.UpdateTodo(context.Background(), db.UpdateTodoParams{
		ID:    todo.ID,
		Title: todo.Title,
		Description: pgtype.Text{
			String: todo.Description,
			Valid:  true,
		},
		Completed: todo.IsCompleted,
		UpdatedAt: pgtype.Timestamp{
			Time:  todo.UpdatedAt,
			Valid: true,
		},
	})
	return err
}

func (r *TodoRepository) Undo(todo *entity.Todo) error {
	err := r.queries.UpdateTodo(context.Background(), db.UpdateTodoParams{
		ID:    todo.ID,
		Title: todo.Title,
		Description: pgtype.Text{
			String: todo.Description,
			Valid:  true,
		},
		Completed: todo.IsCompleted,
		UpdatedAt: pgtype.Timestamp{
			Time:  todo.UpdatedAt,
			Valid: true,
		},
	})
	return err
}

func (r *TodoRepository) Update(todo *entity.Todo) error {
	err := r.queries.UpdateTodo(context.Background(), db.UpdateTodoParams{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: pgtype.Text{String: todo.Description, Valid: true},
		Completed:   todo.IsCompleted,
		UpdatedAt:   pgtype.Timestamp{Time: todo.UpdatedAt, Valid: true},
	})
	return err
}

func (r *TodoRepository) ExistsByID(id uuid.UUID) (bool, error) {
	exists, err := r.queries.ExistsTodo(context.Background(), id)
	return exists, err
}
