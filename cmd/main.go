package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todoapp"
	"todoapp/model"
	"todoapp/store"

	"github.com/gorilla/mux"
)

const (
	contentTypeKey  = "Content-Type"
	applicationJSON = "application/json"
)

var todoService = todoapp.New(store.NewInMemoryStore())

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v0/todos", getTodos).Methods(http.MethodGet)
	r.HandleFunc("/v0/todos", addTodo).Methods(http.MethodPost)
	r.HandleFunc("/v0/todos", updateTodo).Methods(http.MethodPut)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todoService.GetTodos()
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

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed decoding request: %v", err), http.StatusBadRequest)
		return
	}

	err = todoService.SaveTodo(&todo)
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

func updateTodo(w http.ResponseWriter, r *http.Request) {

}
