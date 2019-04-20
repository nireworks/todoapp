package model

import (
	"errors"
	"sort"
)

var (
	ErrInvalidTodo = errors.New("invalid todo")
	ErrNilTodo     = errors.New("nil todo")
)

// Todo is the underlying structure that is bein handled by the TodoService
type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func (t *Todo) IsValid() error {
	if t == nil {
		return ErrNilTodo
	}

	if t.Title == "" {
		return ErrInvalidTodo
	}

	return nil
}

func SortByTitle(todos []*Todo) {
	sort.Slice(todos, func(i, j int) bool { return todos[i].Title < todos[j].Title })
}

func SortById(todos []*Todo) {
	sort.Slice(todos, func(i, j int) bool { return todos[i].Id < todos[j].Id })
}
