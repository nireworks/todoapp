package todoapp_test

import (
	"reflect"
	"testing"
	"todoapp"
	"todoapp/model"

	"github.com/stretchr/testify/assert"
)

func TestTodoApp_GetTodo(t *testing.T) {
	tests := []struct {
		name    string
		todo    *model.Todo
		getID   int
		wantErr bool
	}{
		{
			name: "First working",
			todo: &model.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: true,
			},
			getID:   1,
			wantErr: false,
		},
		{
			name: "Second working",
			todo: &model.Todo{
				Id:        2,
				Title:     "Say goodbye",
				Completed: false,
			},
			getID:   2,
			wantErr: false,
		},
		{
			name: "Not found",
			todo: &model.Todo{
				Id:        3,
				Title:     "Find it",
				Completed: false,
			},
			getID:   4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := todoapp.New()

			err := ta.SaveTodo(tt.todo)
			if assert.Error(t, err) {
				t.Errorf("Saving todo item failed: %v", err)
				return
			}

			got, err := ta.GetTodo(tt.getID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoApp.GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if !reflect.DeepEqual(got, tt.todo) {
				t.Errorf("TodoApp.GetTodo() = %v, want %v", got, tt.todo)
			}
		})
	}
}
