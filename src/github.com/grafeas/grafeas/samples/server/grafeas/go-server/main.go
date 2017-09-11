package main

import (
	"log"
	"net/http"

	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/go"
)

func main() {
	log.Printf("Server started")

	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
