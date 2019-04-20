package main

import (
	"log"
	"net/http"
	"time"
	"todoapp"
	"todoapp/cmd/server"
	"todoapp/store"
)

func main() {
	service := todoapp.New(store.NewInMemoryStore())

	srv := &http.Server{
		Handler:      server.New(service),
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
