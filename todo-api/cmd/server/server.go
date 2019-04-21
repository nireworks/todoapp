package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todoapp"
	"todoapp/model"
	"todoapp/store"

	"github.com/gorilla/mux"
)

const (
	contentTypeKey  = "Content-Type"
	applicationJSON = "application/json"

	ErrJSONDecodeFailed    = "failed decoding request body"
	ErrJSONEncodeFailed    = "failed encoding response body"
	ErrSaveFailed          = "failed saving todo"
	ErrResponseWriteFailed = "failed writing response"
	ErrFetchTodoFailed     = "failed fetching todos"
	ErrInvalidParameter    = "invalid parameter"
	ErrUnknownError        = "something went wrong"
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
			path:    "/v0/todos/{id}",
			handler: s.updateTodo(),
			methods: []string{http.MethodPut},
		},
		{
			path:    "/v0/todos/{id}",
			handler: s.getTodoById(),
			methods: []string{http.MethodGet},
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
			s.sendFailure(w, ErrFetchTodoFailed, err, http.StatusInternalServerError)
			return
		}

		s.sendSuccess(w, todos)
	}
}

func (s *Server) addTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo model.Todo

		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			s.sendFailure(w, ErrJSONDecodeFailed, err, http.StatusBadRequest)
			return
		}

		err = s.service.SaveTodo(&todo)
		if err != nil {
			s.sendFailure(w, ErrSaveFailed, err, http.StatusInternalServerError)
			return
		}

		s.sendSuccess(w, todo)
	}
}

func (s *Server) getTodoById() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]

		id, err := strconv.Atoi(idString)
		if err != nil {
			s.sendFailure(w, ErrInvalidParameter, fmt.Errorf("'%s' cannot be converted to int", idString), http.StatusBadRequest)
			return
		}

		todo, err := s.service.GetTodo(id)
		if err != nil {
			if err == store.ErrTodoNotFound {
				s.sendFailure(w, "fetch with id "+idString, err, http.StatusNotFound)
				return
			}

			s.sendFailure(w, ErrUnknownError, err, http.StatusInternalServerError)
			return
		}

		s.sendSuccess(w, todo)
	}
}
func (s *Server) updateTodo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]

		id, err := strconv.Atoi(idString)
		if err != nil {
			s.sendFailure(w, ErrInvalidParameter, fmt.Errorf("'%s' cannot be converted to int", idString), http.StatusBadRequest)
			return
		}

		var todo model.Todo
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			s.sendFailure(w, ErrJSONDecodeFailed, err, http.StatusBadRequest)
			return
		}

		updatedTodo, err := s.service.UpdateTodo(id, &todo)
		if err != nil {
			s.sendFailure(w, ErrSaveFailed, err, http.StatusInternalServerError)
			return
		}

		s.sendSuccess(w, updatedTodo)
	}
}

func (s *Server) sendSuccess(w http.ResponseWriter, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		s.sendFailure(w, ErrJSONEncodeFailed, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentTypeKey, applicationJSON)
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		s.sendFailure(w, ErrResponseWriteFailed, err, http.StatusInternalServerError)
		return
	}
}

func (s *Server) sendFailure(w http.ResponseWriter, errMsg string, err error, status int) {
	fr := FailResponse{
		Error: fmt.Sprintf("%v: %v", errMsg, err),
	}

	w.WriteHeader(status)

	err = fr.SendJSON(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed sending error: %v", err), http.StatusInternalServerError)
	}
}
