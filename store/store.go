package store

import (
	"errors"
	"sync"
	"sync/atomic"
	"todoapp/model"
)

type Store interface {
	Add(*model.Todo) error
	Update(int, *model.Todo) (*model.Todo, error)
	GetById(int) (*model.Todo, error)
	GetAll() ([]*model.Todo, error)
	Delete(*model.Todo) error
}

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type InMemoryStore struct {
	counter int64
	todoMap map[int]*model.Todo

	sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{todoMap: make(map[int]*model.Todo)}
}

func (ims *InMemoryStore) Add(todo *model.Todo) error {
	if err := todo.IsValid(); err != nil {
		return err
	}

	todo.Id = ims.getId()

	ims.Lock()
	defer ims.Unlock()

	ims.todoMap[todo.Id] = todo

	return nil
}

func (ims *InMemoryStore) GetById(id int) (*model.Todo, error) {
	ims.RLock()
	defer ims.RUnlock()

	todo, ok := ims.todoMap[id]
	if !ok {
		return nil, ErrTodoNotFound
	}

	return todo, nil
}

func (ims *InMemoryStore) Update(id int, todo *model.Todo) (*model.Todo, error) {
	if err := todo.IsValid(); err != nil {
		return nil, err
	}

	ims.Lock()
	defer ims.Unlock()

	_, ok := ims.todoMap[id]
	if !ok {
		return nil, ErrTodoNotFound
	}

	todo.Id = id

	ims.todoMap[id] = todo

	return todo, nil
}

func (ims *InMemoryStore) GetAll() ([]*model.Todo, error) {
	ims.RLock()
	defer ims.RUnlock()

	list := make([]*model.Todo, len(ims.todoMap))

	counter := 0
	for _, todo := range ims.todoMap {
		list[counter] = todo
		counter++
	}

	return list, nil
}

func (ims *InMemoryStore) Delete(todo *model.Todo) error {
	ims.Lock()
	defer ims.Unlock()

	delete(ims.todoMap, todo.Id)

	return nil
}

func (ims *InMemoryStore) getId() int {
	atomic.AddInt64(&ims.counter, 1)

	return int(ims.counter)
}
