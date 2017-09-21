package server

import (
	"log"
	"net/http"

	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/v1alpha1"
)

func main() {
	log.Printf("Server started")
	s := storage.NewMemStore()
	router := NewRouter(v1alpha1.Grafeas{(s)})
	log.Fatal(http.ListenAndServe(":8080", router))
}
