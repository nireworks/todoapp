package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp"
	"todoapp/model"

	"github.com/gorilla/mux"
)

const (
	contentTypeKey  = "Content-Type"
	applicationJSON = "application/json"
)

type Server struct {
	service todoapp.TodoService
	router  *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func New(service todoapp.TodoService) *Server {
	s := &Server{
		service: service,
		router:  mux.NewRouter(),
	}
	s.routes()

	return s
}

func (s *Server) routes() {
	routes := []struct {
		path    string
		handler http.HandlerFunc
		methods []string
	}{
		{
			path:    "/v0/todos",
			handler: s.getTodos(),
			methods: []string{http.MethodGet},
		},
		{
			path:    "/v0/todos",
			handler: s.addTodo(),
			methods: []string{http.MethodPost},
		},
		{
			path:    "/v0/todos",
			handler: s.updateTodo(),
			methods: []string{http.MethodPut},
		},
	}

	for _, route := range routes {
		s.router.
			HandleFunc(route.path, route.handler).
			Methods(route.methods...)
	}
}

func (s *Server) getTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := s.service.GetTodos()

		if err != nil {
			http.Error(w, fmt.Sprintf("failed fetching todos: %v", err), http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed marshalling response: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set(contentTypeKey, applicationJSON)
		w.WriteHeader(http.StatusOK)

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed writing response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) addTodo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var todo model.Todo

		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed decoding request: %v", err), http.StatusBadRequest)
			return
		}

		err = s.service.SaveTodo(&todo)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed saving todo: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set(contentTypeKey, applicationJSON)
		w.WriteHeader(http.StatusOK)

		_, err = w.Write([]byte("Success!"))
		if err != nil {
			http.Error(w, fmt.Sprintf("failed writing response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) updateTodo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
