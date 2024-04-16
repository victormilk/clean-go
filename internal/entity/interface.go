package entity

import "github.com/google/uuid"

type TodoRepositoryInterface interface {
	Save(todo *Todo) error
	GetAll() ([]*Todo, error)
	GetByID(id uuid.UUID) (*Todo, error)
	DeleteByID(id uuid.UUID) error
	Complete(todo *Todo) error
	Undo(todo *Todo) error
	Update(todo *Todo) error
	ExistsByID(id uuid.UUID) (bool, error)
}
