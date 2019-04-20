package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp"
	"todoapp/cmd/server"
	"todoapp/model"
	"todoapp/store"

	"github.com/stretchr/testify/assert"
)

func TestHandle_GetTodos(t *testing.T) {
	tests := []struct {
		name       string
		todos      []*model.Todo
		wantBody   string
		wantStatus int
	}{
		{
			name:       "empty store",
			todos:      []*model.Todo{},
			wantBody:   "[]",
			wantStatus: http.StatusOK,
		},
		{
			name: "one todo",
			todos: []*model.Todo{
				{Title: "Hey"},
			},
			wantBody:   "[{\"id\":1,\"title\":\"Hey\",\"completed\":false}]",
			wantStatus: http.StatusOK,
		},
		{
			name: "one todo with all fields",
			todos: []*model.Todo{
				{Id: 1, Title: "Hey", Completed: true},
			},
			wantBody:   "[{\"id\":1,\"title\":\"Hey\",\"completed\":true}]",
			wantStatus: http.StatusOK,
		},
		{
			name: "one todo with wrong Id",
			todos: []*model.Todo{
				{Id: 10, Title: "Hey", Completed: true},
			},
			wantBody:   "[{\"id\":1,\"title\":\"Hey\",\"completed\":true}]",
			wantStatus: http.StatusOK,
		},
		{
			name: "two todos",
			todos: []*model.Todo{
				{Id: 1, Title: "Hey", Completed: true},
				{Id: 2, Title: "Hello", Completed: false},
			},
			wantBody:   "[{\"id\":1,\"title\":\"Hey\",\"completed\":true},{\"id\":2,\"title\":\"Hello\",\"completed\":false}]",
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.NewInMemoryStore()

			for _, todo := range tt.todos {
				err := mockStore.Add(todo)
				assert.NoError(t, err)
			}

			srv := server.New(todoapp.New(mockStore))

			req, err := http.NewRequest(http.MethodGet, "/v0/todos", nil)
			if err != nil {
				t.Errorf("failed constructing post request: %v", err)
				return
			}

			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestHandler_AddTodo(t *testing.T) {
	tests := []struct {
		name       string
		addTodos   []string
		wantTodos  []*model.Todo
		wantStatus int
		wantBody   string
	}{
		{
			name:       "nothing to add",
			addTodos:   []string{""},
			wantTodos:  []*model.Todo{},
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"error\":\"failed decoding request body: EOF\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.NewInMemoryStore()

			srv := server.New(todoapp.New(mockStore))

			for _, todo := range tt.addTodos {
				req, err := http.NewRequest(http.MethodPost, "/v0/todos", bytes.NewBuffer([]byte(todo)))
				assert.NoError(t, err)

				w := httptest.NewRecorder()
				srv.ServeHTTP(w, req)

				assert.Equal(t, tt.wantStatus, w.Code)
				assert.Equal(t, tt.wantBody, w.Body.String())
			}
		})
	}
}