package todoapp_test

import (
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
				Title:     "Say hello",
				Completed: true,
			},
			getID:   1,
			wantErr: false,
		},
		{
			name: "Second working",
			todo: &model.Todo{
				Title:     "Say goodbye",
				Completed: false,
			},
			getID:   1,
			wantErr: false,
		},
		{
			name: "Not found",
			todo: &model.Todo{
				Title:     "Find it",
				Completed: false,
			},
			getID:   99,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := todoapp.New()

			err := ta.SaveTodo(tt.todo)
			if err != nil {
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

			assert.Equal(t, tt.todo, got)
		})
	}
}

func TestTodoApp_SaveTodo(t *testing.T) {
	tests := []struct {
		name    string
		todo    *model.Todo
		wantErr bool
	}{
		{
			name: "First working",
			todo: &model.Todo{
				Title:     "Say hello",
				Completed: true,
			},
			wantErr: false,
		},
		{
			name: "invalid todo",
			todo: &model.Todo{
				Title:     "",
				Completed: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := todoapp.New()

			err := ta.SaveTodo(tt.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveTodo(%v) error=%v, wantErr=%t", tt.todo, err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			got, err := ta.GetTodo(1)
			if err != nil {
				t.Errorf("GetTodo() failed: %v", err)
				return
			}

			assert.Equal(t, tt.todo, got)
		})
	}
}
