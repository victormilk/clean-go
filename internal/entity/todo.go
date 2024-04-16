package entity

import (
	"errors"
	"time"
)

import "github.com/google/uuid"

type TodoInterface interface {
	IsValid() error
	Complete() error
	Undo() error
	Update(title, description string) error
}

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodo(title, description string) (*Todo, error) {
	todo := &Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := todo.IsValid()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

var (
	ErrTitleIsRequired    = errors.New("title is required")
	ErrTitleTooShort      = errors.New("title is too short")
	ErrTitleTooLong       = errors.New("title is too long")
	ErrDescriptionTooLong = errors.New("description is too long")
)

func (t *Todo) IsValid() error {
	if t.Title == "" {
		return ErrTitleIsRequired
	}
	if len(t.Title) < 3 {
		return ErrTitleTooShort
	}
	if len(t.Title) > 50 {
		return ErrTitleTooLong
	}
	if len(t.Description) > 1000 {
		return ErrDescriptionTooLong
	}
	return nil
}

func (t *Todo) Complete() error {
	t.IsCompleted = true
	t.UpdatedAt = time.Now()
	return t.IsValid()
}

func (t *Todo) Undo() error {
	t.IsCompleted = false
	t.UpdatedAt = time.Now()
	return t.IsValid()
}

func (t *Todo) Update(title, description string) error {
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
	return t.IsValid()
}
