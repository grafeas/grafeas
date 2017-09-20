package server

import (
	"log"
	"net/http"

	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
)

func main() {
	log.Printf("Server started")
	s := storage.NewMemStore()
	router := NewRouter(s)
	log.Fatal(http.ListenAndServe(":8080", router))
}
