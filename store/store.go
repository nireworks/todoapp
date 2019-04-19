package store

import (
	"errors"
	"sync"
	"sync/atomic"
	"todoapp/model"
)

type Store interface {
	Add(*model.Todo) error
	GetById(int) (*model.Todo, error)
	GetAll() ([]*model.Todo, error)
	Delete(*model.Todo) error
}

var (
	ErrInvalidTodo  = errors.New("invalid todo")
	ErrNilTodo      = errors.New("nil todo")
	ErrTodoNotFound = errors.New("todo not found")
)

type InMemoryStore struct {
	counter int64

	mu      sync.RWMutex
	todoMap map[int]*model.Todo
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{todoMap: make(map[int]*model.Todo)}
}

func (ims *InMemoryStore) Add(todo *model.Todo) error {
	if err := isValid(todo); err != nil {
		return err
	}

	todo.Id = ims.getId()

	ims.mu.Lock()
	defer ims.mu.Unlock()

	ims.todoMap[todo.Id] = todo

	return nil
}

func (ims *InMemoryStore) GetById(id int) (*model.Todo, error) {
	ims.mu.RLock()
	defer ims.mu.RUnlock()

	todo, ok := ims.todoMap[id]
	if !ok {
		return nil, ErrTodoNotFound
	}

	return todo, nil
}

func (ims *InMemoryStore) GetAll() ([]*model.Todo, error) {
	ims.mu.RLock()
	defer ims.mu.RUnlock()

	list := make([]*model.Todo, len(ims.todoMap))

	counter := 0
	for _, todo := range ims.todoMap {
		list[counter] = todo
		counter++
	}

	return list, nil
}

func (ims *InMemoryStore) Delete(todo *model.Todo) error {
	ims.mu.Lock()
	defer ims.mu.Unlock()

	delete(ims.todoMap, todo.Id)

	return nil
}

func (ims *InMemoryStore) getId() int {
	atomic.AddInt64(&ims.counter, 1)

	return int(ims.counter)
}

func isValid(todo *model.Todo) error {
	if todo == nil {
		return ErrNilTodo
	}

	if todo.Title == "" {
		return ErrInvalidTodo
	}

	return nil
}
