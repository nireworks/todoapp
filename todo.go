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
	todo, err := t.backend.GetById(index)
	if err != nil {
		if err == store.ErrTodoNotFound {
			return nil, err
		}

		return nil, fmt.Errorf("unexpected error from backend: %v", err)
	}

	return todo, nil
}

func (t *TodoApp) GetTodos() ([]*model.Todo, error) {
	todos, err := t.backend.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get all todos: %v", err)
	}

	return todos, nil
}

func (t *TodoApp) SaveTodo(todo *model.Todo) error {
	if err := todo.IsValid(); err != nil {
		return fmt.Errorf("save todo: %v", err)
	}

	err := t.backend.Add(todo)
	if err != nil {
		return fmt.Errorf("save todo: %v", err)
	}

	return nil
}
