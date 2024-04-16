package create

import (
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
	"todoList/test/mocks"
)

func TestGivenAnInput_WhenCallsExecute_ThenShouldCreateATodoOrReceiveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tc := range []struct {
		title       string
		description string
		wantErr     bool
	}{
		{"", "", true},
		{"valid title", "", false},
		{"valid title", "valid description", false},
		{strings.Repeat("invalid title", 51), "valid description", true},
		{"valid title", strings.Repeat("invalid description", 1001), true},
	} {
		mock := mocks.NewMockTodoRepositoryInterface(ctrl)
		if !tc.wantErr {
			mock.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()
		}
		u := NewCreateTodo(mock)
		err := u.Execute(CreateTodoInput{Title: tc.title, Description: tc.description})
		if tc.wantErr && err == nil {
			t.Errorf("Test case %d: expected error, got nil", i)
			mock.EXPECT().Save(gomock.Any()).Times(0)
		}
	}
}
