package server

import (
	"log"
	"net/http"

)

func main() {
	log.Printf("Server started")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
