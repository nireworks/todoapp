package todoapp_test

import (
	"reflect"
	"testing"
	"todoapp"

	"github.com/stretchr/testify/assert"
)

func setup() {

}

func TestTodoApp_GetTodo(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		todo    todoapp.Todo
		getID   int
		wantErr bool
	}{
		{
			name: "First working",
			todo: todoapp.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: true,
			},
			getID:   1,
			wantErr: false,
		},
		{
			name: "Second working",
			todo: todoapp.Todo{
				Id:        2,
				Title:     "Say goodbye",
				Completed: false,
			},
			getID:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := todoapp.New()

			err := ta.SaveTodo(tt.todo)
			assert.NoError(t, err)

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
