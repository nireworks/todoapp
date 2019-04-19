package store

import (
	"testing"
	"todoapp/model"

	"github.com/stretchr/testify/assert"
)

func TestNewInMemoryStore(t *testing.T) {
	ims := NewInMemoryStore()

	assert.NotNil(t, ims.todoMap)
}

func TestInMemoryStore_Add(t *testing.T) {
	tests := []struct {
		name    string
		todo    *model.Todo
		wantErr bool
	}{
		{
			"make it successful",
			&model.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: false,
			},
			false,
		},
		{
			"missing title",
			&model.Todo{
				Id:        1,
				Title:     "",
				Completed: false,
			},
			true,
		},
		{
			"nil todo",
			nil,
			true,
		},
	}

	for _, tt := range tests {
		ims := NewInMemoryStore()

		t.Run(tt.name, func(t *testing.T) {
			if err := ims.Add(tt.todo); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			readTodo, err := ims.GetById(tt.todo.Id)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.todo, readTodo)
			}
		})
	}
}

func TestInMemoryStore_GetById(t *testing.T) {
	tests := []struct {
		name    string
		todo    *model.Todo
		byId    int
		wantErr bool
	}{
		{
			"make it successful",
			&model.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: false,
			},
			1,
			false,
		},
		{
			"make it fail",
			&model.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: false,
			},
			2,
			true,
		},
	}

	for _, tt := range tests {
		ims := NewInMemoryStore()

		t.Run(tt.name, func(t *testing.T) {
			err := ims.Add(tt.todo)
			if err != nil {
				t.Errorf("Failed adding todo: %v", err)
				return
			}

			readTodo, err := ims.GetById(tt.byId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById(%d) wantErr=%t, but got: %v", tt.byId, tt.wantErr, err)
				return
			}

			if tt.wantErr {
				return
			}

			assert.Equal(t, tt.todo, readTodo)
		})
	}
}

func TestInMemoryStore_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		todos   []*model.Todo
		wantErr bool
	}{
		{
			"no elements",
			[]*model.Todo{},
			false,
		},
		{
			"one element",
			[]*model.Todo{
				{Title: "Say hello", Completed: false},
			},
			false,
		},
		{
			"two elements",
			[]*model.Todo{
				{Title: "Say hello", Completed: false},
				{Title: "Say Goodbye", Completed: false},
			},
			false,
		},
		{
			"three elements",
			[]*model.Todo{
				{Title: "Say hello", Completed: false},
				{Title: "Say Goodbye", Completed: false},
				{Title: "Say Whatever", Completed: false},
			},
			false,
		},
	}

	for _, tt := range tests {
		ims := NewInMemoryStore()

		t.Run(tt.name, func(t *testing.T) {
			for _, todo := range tt.todos {
				err := ims.Add(todo)
				if err != nil {
					t.Errorf("Failed adding todo: %v", err)
					return
				}
			}

			readTodos, err := ims.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() wantErr=%t, but got: %v", tt.wantErr, err)
				return
			}

			if tt.wantErr {
				return
			}

			if assert.Equal(t, len(tt.todos), len(readTodos)) {
				for i, todo := range readTodos {
					assert.Equal(t, tt.todos[i], todo)
				}
			}
		})
	}
}

func TestInMemoryStore_Delete(t *testing.T) {
	tests := []struct {
		name           string
		todo           *model.Todo
		wantErr        bool
		deleteId       int
		expectedLength int
	}{
		{
			"Delete one entry",
			&model.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: false,
			},
			false,
			1,
			0,
		},
		{
			"Delete non-existent entry",
			&model.Todo{
				Id:        1,
				Title:     "Say hello",
				Completed: false,
			},
			false,
			2,
			1,
		},
	}
	for _, tt := range tests {
		ims := NewInMemoryStore()

		t.Run(tt.name, func(t *testing.T) {
			err := ims.Add(tt.todo)
			if err != nil {
				t.Errorf("Failed adding todo: %v", err)
				return
			}

			err = ims.Delete(&model.Todo{Id: tt.deleteId})
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete(%d) wantErr=%t, but got: %v", tt.deleteId, tt.wantErr, err)
				return
			}

			if tt.wantErr {
				return
			}

			assert.Equal(t, len(ims.todoMap), tt.expectedLength)
		})
	}
}
