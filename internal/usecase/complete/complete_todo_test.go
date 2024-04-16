package complete

import (
	"errors"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
	"todoList/internal/entity"
	"todoList/internal/usecase"
	"todoList/test/mocks"
)

func TestGivenAnInput_WhenCallsExecute_ThenShouldCompleteTodoOrReceiveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tc := range []struct {
		id             string
		shouldFind     bool
		shouldComplete bool
		wantErr        bool
	}{
		{"invalid id", false, true, true},
		{uuid.New().String(), false, true, true},
		{uuid.New().String(), true, true, true},
		{uuid.New().String(), true, false, true},
		{uuid.New().String(), true, true, false},
	} {
		mock := mocks.NewMockTodoRepositoryInterface(ctrl)

		if tc.wantErr && tc.shouldFind && tc.shouldComplete {
			mock.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("error")).AnyTimes()
		} else if tc.wantErr && !tc.shouldFind && tc.shouldComplete {
			mock.EXPECT().GetByID(gomock.Any()).Return(nil, usecase.ErrNoResultRows).AnyTimes()
		} else if tc.wantErr && tc.shouldFind && !tc.shouldComplete {
			mock.EXPECT().GetByID(gomock.Any()).Return(&entity.Todo{
				ID:          uuid.MustParse(tc.id),
				Title:       "",
				Description: strings.Repeat("invalid_description", 1001),
				IsCompleted: false,
			}, nil).AnyTimes()
		} else {
			mock.EXPECT().GetByID(gomock.Any()).Return(&entity.Todo{
				ID:          uuid.MustParse(tc.id),
				Title:       "title",
				Description: "description",
				IsCompleted: false,
			}, nil).AnyTimes()
			mock.EXPECT().Complete(gomock.Any()).Return(nil).AnyTimes()
		}

		u := NewCompleteTodo(mock)
		input := CompleteTodoInput{ID: tc.id}
		err := u.Execute(input)
		if tc.wantErr && err == nil {
			t.Errorf("Test case %d: expected error, got nil", i)
			mock.EXPECT().Complete(gomock.Any()).Times(0)
		}
	}
}
