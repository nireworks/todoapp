package todoapp

import (
	"fmt"
	"todoapp/model"
	"todoapp/store"
)

type TodoApp struct {
	backend store.Store
}

func New() *TodoApp {
	return &TodoApp{
		backend: store.NewInMemoryStore(),
	}
}

func (t *TodoApp) GetTodo(index int) (*model.Todo, error) {
	return nil, fmt.Errorf("not implemented")
}

func (t *TodoApp) GetTodos() ([]*model.Todo, error) {
	return nil, fmt.Errorf("not implemented")
}

func (t *TodoApp) SaveTodo(todo *model.Todo) error {
	return fmt.Errorf("not implemented")
}
