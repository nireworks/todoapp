package store

import (
	"testing"
	"todoapp/model"

	"github.com/stretchr/testify/assert"
)

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
			if assert.Error(t, err) {
				t.Fail()
				return
			}

			readTodo, err := ims.GetById(tt.byId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById(%d) wantErr=%t, but got: %v", tt.byId, tt.wantErr, err)
				return
			}

			assert.Equal(t, tt.todo, readTodo)
		})
	}
}

func TestNewInMemoryStore(t *testing.T) {
	ims := NewInMemoryStore()

	assert.NotNil(t, ims.todoMap)
}
