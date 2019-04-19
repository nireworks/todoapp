package todoapp

import "todoapp/model"

type TodoService interface {
	GetTodo(int) (*model.Todo, error)
	GetTodos() ([]*model.Todo, error)
	SaveTodo(*model.Todo) error
}
