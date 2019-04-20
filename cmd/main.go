package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todoapp"
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
	r.HandleFunc("/v0/todos", getTodos)

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
