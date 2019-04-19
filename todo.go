package todoapp

import (
	"fmt"
	"todoapp/model"
)

type TodoApp struct {
}

func New() *TodoApp {
	return &TodoApp{}
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
