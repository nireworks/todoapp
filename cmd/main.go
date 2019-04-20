package main

import (
	"log"
	"todoapp"
	"todoapp/cmd/server"
	"todoapp/store"
)

func main() {
	service := todoapp.New(store.NewInMemoryStore())

	srv := server.New(service)

	log.Fatal(srv.Server(":8000").ListenAndServe())
}
