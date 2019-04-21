package main

import (
	"log"
	"net/http"
	"time"
	"todoapp"
	"todoapp/cmd/server"
	"todoapp/store"

	"github.com/gorilla/handlers"
)

func main() {
	service := todoapp.New(store.NewInMemoryStore())

	headers := handlers.AllowedHeaders([]string{"X-Requested-With"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	srv := &http.Server{
		Handler:      handlers.CORS(headers, origins, methods)(server.New(service)),
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("starting to listen on %v", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
