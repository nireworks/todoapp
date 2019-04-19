package store

import (
	"reflect"
	"testing"
	"todoapp/model"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryStore_Add(t *testing.T) {
	type args struct {
		in0 *model.Todo
	}
	tests := []struct {
		name    string
		ims     *InMemoryStore
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ims.Add(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryStore_GetById(t *testing.T) {
	type args struct {
		in0 int
	}
	tests := []struct {
		name    string
		ims     *InMemoryStore
		args    args
		want    *model.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ims.GetById(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryStore.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryStore_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		ims     *InMemoryStore
		want    *[]model.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ims.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryStore.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryStore_Delete(t *testing.T) {
	type args struct {
		in0 *model.Todo
	}
	tests := []struct {
		name    string
		ims     *InMemoryStore
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ims.Delete(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryStore.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewInMemoryStore(t *testing.T) {
	ims := NewInMemoryStore()

	assert.NotNil(t, ims.todoMap)
}
