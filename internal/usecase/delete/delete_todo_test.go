package delete

import (
	"errors"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"testing"
	"todoList/test/mocks"
)

func TestGivenAnInput_WhenCallsExecute_ThenShouldDeleteATodoOrReceiveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tc := range []struct {
		id          string
		shouldExist bool
		wantErr     bool
	}{
		{"invalid id", false, true},
		{uuid.New().String(), false, true},
		{uuid.New().String(), true, true},
		{uuid.New().String(), true, false},
	} {
		mock := mocks.NewMockTodoRepositoryInterface(ctrl)

		if tc.wantErr && tc.shouldExist {
			mock.EXPECT().ExistsByID(gomock.Any()).Return(false, errors.New("error")).AnyTimes()
		} else if tc.wantErr && !tc.shouldExist {
			mock.EXPECT().ExistsByID(gomock.Any()).Return(false, nil).AnyTimes()
		} else {
			mock.EXPECT().ExistsByID(gomock.Any()).Return(true, nil).AnyTimes()
			mock.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes()
		}

		u := NewDeleteTodo(mock)
		err := u.Execute(DeleteTodoInput{ID: tc.id})
		if tc.wantErr && err == nil {
			t.Errorf("Test case %d: expected error, got nil", i)
			mock.EXPECT().DeleteByID(gomock.Any()).Times(0)
		}
	}
}
