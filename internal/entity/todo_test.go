package entity

import (
	"strings"
	"testing"
)

func TestGivenATitle_WhenCallsNewTodo_ThenShouldReceiveErrorOrCreate(t *testing.T) {
	for i, tc := range []struct {
		title   string
		wantErr bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{strings.Repeat("invalid title", 51), true},
		{"valid title", false},
	} {
		_, err := NewTodo(tc.title, "description")
		if tc.wantErr && err == nil {
			t.Errorf("Test case %d: expected error, got nil", i)
		}
	}
}

func TestGivenADescription_WhenCallsNewTodo_ThenShouldReturnErrorOrCreate(t *testing.T) {
	for i, tc := range []struct {
		description string
		wantErr     bool
	}{
		{strings.Repeat("invalid description", 1001), true},
		{"", false},
		{"valid description", false},
	} {
		_, err := NewTodo("title", tc.description)
		if tc.wantErr && err == nil {
			t.Errorf("Test case %d: expected error, got nil", i)
		}
	}
}

func TestGivenATodo_WhenCallsComplete_ThenShouldChangeIsCompletedToTrue(t *testing.T) {
	todo, _ := NewTodo("valid title", "valid description")
	_ = todo.Complete()
	if !todo.IsCompleted {
		t.Errorf("Expected true, got %v", todo.IsCompleted)
	}
}

func TestGivenATodo_WhenCallsUndo_ThenShouldChangeIsCompletedToFalse(t *testing.T) {
	todo, _ := NewTodo("valid title", "valid description")
	_ = todo.Complete()
	_ = todo.Undo()
	if todo.IsCompleted {
		t.Errorf("Expected false, got %v", todo.IsCompleted)
	}
}

func TestGivenATitleAndDescription_WhenCallsUpdate_ThenShouldUpdateTitleAndDescription(t *testing.T) {
	for i, tc := range []struct {
		title       string
		description string
		wantErr     bool
	}{
		{"", "valid description", true},
		{"valid title", strings.Repeat("invalid description", 1001), true},
		{"valid title", "", false},
		{"valid title", "valid description", false},
	} {
		todo, _ := NewTodo("title", "description")
		err := todo.Update(tc.title, tc.description)
		if tc.wantErr && err == nil {
			t.Errorf("Test case %d: expected error, got nil", i)
		}
	}
}
