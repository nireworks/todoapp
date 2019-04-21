package server_test

import (
	"bytes"
	"fmt"
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
				t.Errorf("failed constructing get request: %v", err)
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
		wantStatus int
		wantBodies []string
	}{
		{
			name:       "nothing to add",
			addTodos:   []string{""},
			wantStatus: http.StatusBadRequest,
			wantBodies: []string{"{\"error\":\"failed decoding request body: EOF\"}"},
		},
		{
			name:       "Add one",
			addTodos:   []string{"{\"id\":1,\"title\":\"Hey\",\"completed\":false}"},
			wantStatus: http.StatusOK,
			wantBodies: []string{"{\"id\":1,\"title\":\"Hey\",\"completed\":false}"},
		},
		{
			name:       "Add one with no ID",
			addTodos:   []string{"{\"title\":\"Hey\",\"completed\":false}"},
			wantStatus: http.StatusOK,
			wantBodies: []string{"{\"id\":1,\"title\":\"Hey\",\"completed\":false}"},
		},
		{
			name:       "Add one with no id and completed true",
			addTodos:   []string{"{\"title\":\"Hey\",\"completed\":true}"},
			wantStatus: http.StatusOK,
			wantBodies: []string{"{\"id\":1,\"title\":\"Hey\",\"completed\":true}"},
		},
		{
			name: "Add two",
			addTodos: []string{
				"{\"title\":\"Hey\",\"completed\":true}",
				"{\"title\":\"Hey Again\",\"completed\":false}",
			},
			wantStatus: http.StatusOK,
			wantBodies: []string{
				"{\"id\":1,\"title\":\"Hey\",\"completed\":true}",
				"{\"id\":2,\"title\":\"Hey Again\",\"completed\":false}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := store.NewInMemoryStore()

			srv := server.New(todoapp.New(mockStore))

			for idx, todo := range tt.addTodos {
				req, err := http.NewRequest(http.MethodPost, "/v0/todos", bytes.NewBuffer([]byte(todo)))
				assert.NoError(t, err)

				w := httptest.NewRecorder()
				srv.ServeHTTP(w, req)

				assert.Equal(t, tt.wantStatus, w.Code)
				assert.Equal(t, tt.wantBodies[idx], w.Body.String())
			}
		})
	}
}

func TestHandler_GetTodoById(t *testing.T) {
	tests := []struct {
		name       string
		todos      []*model.Todo
		fetchId    string
		wantStatus int
		wantBody   string
	}{
		{
			name: "Fetch non-existent",
			todos: []*model.Todo{
				{Title: "Hey"},
			},
			fetchId:    "99",
			wantStatus: http.StatusNotFound,
			wantBody:   "{\"error\":\"fetch with id 99: todo not found\"}",
		},
		{
			name: "Fetch non-int",
			todos: []*model.Todo{
				{Title: "Hey"},
			},
			fetchId:    "asd",
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"error\":\"invalid parameter: 'asd' cannot be converted to int\"}",
		},
		{
			name: "Fetch one",
			todos: []*model.Todo{
				{Title: "Hey"},
			},
			fetchId:    "1",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":1,\"title\":\"Hey\",\"completed\":false}",
		},
		{
			name: "Fetch first of three",
			todos: []*model.Todo{
				{Title: "Hey", Completed: true},
				{Title: "Hey Again"},
				{Title: "Good Bye"},
			},
			fetchId:    "1",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":1,\"title\":\"Hey\",\"completed\":true}",
		},
		{
			name: "Fetch second of three",
			todos: []*model.Todo{
				{Title: "Hey", Completed: true},
				{Title: "Hey Again"},
				{Title: "Good Bye"},
			},
			fetchId:    "2",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":2,\"title\":\"Hey Again\",\"completed\":false}",
		},
		{
			name: "Fetch third of three",
			todos: []*model.Todo{
				{Title: "Hey", Completed: true},
				{Title: "Hey Again"},
				{Title: "Good Bye"},
			},
			fetchId:    "3",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":3,\"title\":\"Good Bye\",\"completed\":false}",
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

			url := fmt.Sprintf("/v0/todos/%s", tt.fetchId)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Errorf("failed constructing get request: %v", err)
				return
			}

			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestHandle_UpdateTodo(t *testing.T) {
	tests := []struct {
		name       string
		todos      []*model.Todo
		updateTodo string
		updateId   string
		wantStatus int
		wantBody   string
		wantTodos  []*model.Todo
	}{
		{
			name:       "empty request",
			todos:      []*model.Todo{},
			updateId:   "1",
			updateTodo: "",
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"error\":\"failed decoding request body: EOF\"}",
			wantTodos:  []*model.Todo{},
		},
		{
			name: "Update existing",
			todos: []*model.Todo{
				{Title: "Hello"},
			},
			updateTodo: "{\"id\":1,\"title\":\"Updated\",\"completed\":false}",
			updateId:   "1",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":1,\"title\":\"Updated\",\"completed\":false}",
			wantTodos: []*model.Todo{
				{Id: 1, Title: "Updated", Completed: false},
			},
		},
		{
			name: "Update first of three",
			todos: []*model.Todo{
				{Title: "First"},
				{Title: "Second"},
				{Title: "Third"},
			},
			updateTodo: "{\"id\":1,\"title\":\"Updated\",\"completed\":true}",
			updateId:   "1",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":1,\"title\":\"Updated\",\"completed\":true}",
			wantTodos: []*model.Todo{
				{Id: 1, Title: "Updated", Completed: true},
				{Id: 2, Title: "Second", Completed: false},
				{Id: 3, Title: "Third", Completed: false},
			},
		},
		{
			name: "Update second of three",
			todos: []*model.Todo{
				{Title: "First"},
				{Title: "Second"},
				{Title: "Third"},
			},
			updateTodo: "{\"id\":2,\"title\":\"Updated\",\"completed\":true}",
			updateId:   "2",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":2,\"title\":\"Updated\",\"completed\":true}",
			wantTodos: []*model.Todo{
				{Id: 1, Title: "First", Completed: false},
				{Id: 2, Title: "Updated", Completed: true},
				{Id: 3, Title: "Third", Completed: false},
			},
		},
		{
			name: "Update third of three",
			todos: []*model.Todo{
				{Title: "First"},
				{Title: "Second"},
				{Title: "Third"},
			},
			updateTodo: "{\"id\":3,\"title\":\"Updated\",\"completed\":true}",
			updateId:   "3",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":3,\"title\":\"Updated\",\"completed\":true}",
			wantTodos: []*model.Todo{
				{Id: 1, Title: "First", Completed: false},
				{Id: 2, Title: "Second", Completed: false},
				{Id: 3, Title: "Updated", Completed: true},
			},
		},
		{
			name: "Ignore ID in payload",
			todos: []*model.Todo{
				{Title: "First"},
				{Title: "Second"},
				{Title: "Third"},
			},
			updateTodo: "{\"id\":1,\"title\":\"Updated\",\"completed\":true}",
			updateId:   "3",
			wantStatus: http.StatusOK,
			wantBody:   "{\"id\":3,\"title\":\"Updated\",\"completed\":true}",
			wantTodos: []*model.Todo{
				{Id: 1, Title: "First", Completed: false},
				{Id: 2, Title: "Second", Completed: false},
				{Id: 3, Title: "Updated", Completed: true},
			},
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
			url := fmt.Sprintf("/v0/todos/%s", tt.updateId)
			req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer([]byte(tt.updateTodo)))
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantBody, w.Body.String())

			todos, _ := mockStore.GetAll()

			model.SortById(todos)
			model.SortById(tt.wantTodos)

			for idx, todo := range todos {
				assert.Equal(t, tt.wantTodos[idx], todo)
			}
		})
	}
}
