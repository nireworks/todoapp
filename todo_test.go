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
		todos   []*model.Todo
		wantErr bool
	}{
		{
			name: "One todo",
			todos: []*model.Todo{
				{Title: "Say hello", Completed: true},
			},
			wantErr: false,
		},
		{
			name: "One invalid todo",
			todos: []*model.Todo{
				{Title: "", Completed: false},
			},
			wantErr: true,
		},
		{
			name: "Two todos",
			todos: []*model.Todo{
				{Title: "Say hello", Completed: true},
				{Title: "Say bye", Completed: true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := todoapp.New()

			for _, todo := range tt.todos {
				err := ta.SaveTodo(todo)
				if (err != nil) != tt.wantErr {
					t.Errorf("SaveTodo(%v) error=%v, wantErr=%t", todo, err, tt.wantErr)
					return
				}

				if tt.wantErr {
					return
				}
			}

			allTodos, err := ta.GetTodos()
			if err != nil {
				t.Errorf("GetTodos() failed: %v", err)
				return
			}

			assert.Equal(t, len(allTodos), len(tt.todos))

			got, err := ta.GetTodo(1)
			if err != nil {
				t.Errorf("GetTodo() failed: %v", err)
				return
			}

			assert.Equal(t, tt.todos[0], got)
		})
	}
}

func TestTodoApp_GetTodos(t *testing.T) {
	tests := []struct {
		name  string
		todos []*model.Todo
	}{
		{
			name: "One todo",
			todos: []*model.Todo{
				{Title: "Say hello", Completed: true},
			},
		},
		{
			name: "Two todos",
			todos: []*model.Todo{
				{Title: "Say Hello", Completed: false},
				{Title: "Say Goodbye", Completed: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := todoapp.New()

			for _, todo := range tt.todos {
				err := ta.SaveTodo(todo)
				if err != nil {
					t.Errorf("SaveTodo(%v) failed: %v", todo, err)
					return
				}
			}

			allTodos, err := ta.GetTodos()
			if err != nil {
				t.Errorf("GetTodos() failed: %v", err)
				return
			}

			assert.Equal(t, len(allTodos), len(tt.todos))

			model.SortById(allTodos)
			model.SortById(tt.todos)

			for i, todo := range tt.todos {
				assert.Equal(t, todo, allTodos[i])
			}

			got, err := ta.GetTodo(1)
			if err != nil {
				t.Errorf("GetTodo() failed: %v", err)
				return
			}

			assert.Equal(t, tt.todos[0], got)
		})
	}
}
