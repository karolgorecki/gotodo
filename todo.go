package main

import (
	"log"
	"net/http"

	_ "github.com/karolgorecki/todo/boltstore" // pick the store
	"github.com/karolgorecki/todo/server"
)

func main() {
	// Create new router with all handlers
	rt := server.RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8000", rt))
}
