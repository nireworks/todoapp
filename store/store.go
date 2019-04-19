package store

import (
	"fmt"
	"todoapp/model"
)

type Store interface {
	Add(*model.Todo) error
	GetById(int) (*model.Todo, error)
	GetAll() ([]*model.Todo, error)
	Delete(*model.Todo) error
}

type InMemoryStore struct {
	todoMap map[int]*model.Todo
}

func (ims *InMemoryStore) Add(*model.Todo) error {
	return fmt.Errorf("not implemented")

}

func (ims *InMemoryStore) GetById(int) (*model.Todo, error) {
	return nil, fmt.Errorf("not implemented")
}

func (ims *InMemoryStore) GetAll() ([]*model.Todo, error) {
	return nil, fmt.Errorf("not implemented")
}

func (ims *InMemoryStore) Delete(*model.Todo) error {
	return fmt.Errorf("not implemented")
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{todoMap: make(map[int]*model.Todo)}
}
