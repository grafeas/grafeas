package server

import (
	"encoding/json"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
	"io/ioutil"
	"log"
	"net/http"
)

type Grafeas struct {
	s storage.Store
}

func (g *Grafeas) CreateNote(w http.ResponseWriter, r *http.Request) {
	n := swagger.Note{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &n)
	if n.Name == "" {
		log.Printf("Invalid note name: %v", n.Name)
		http.Error(w, "Invalid note name", http.StatusBadRequest)
	}
	// TODO: Validate that operation exists if it is specified
	g.s.CreateNote(&n)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, err := json.Marshal(&n)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", err)
	}
	w.Write(bytes)
}

func (g *Grafeas) CreateOccurrence(w http.ResponseWriter, r *http.Request) {
	o := swagger.Occurrence{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &o)
	if o.Name == "" {
		log.Printf("Invalid occurrence name: %v", o.Name)
		http.Error(w, "Invalid occurrences name", http.StatusBadRequest)
	}
	// TODO: Validate that note exists
	// TODO: Validate that operation exists if it is specified
	g.s.CreateOccurrence(&o)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, err := json.Marshal(&o)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", err)
	}
	w.Write(bytes)
}

func (g *Grafeas) CreateOperation(w http.ResponseWriter, r *http.Request) {
	o := swagger.Operation{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &o)
	if o.Name == "" {
		log.Printf("Invalid occurrence name: %v", o.Name)
		http.Error(w, "Invalid occurrences name", http.StatusBadRequest)
	}
	g.s.CreateOperation(&o)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, err := json.Marshal(&o)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", err)
	}
	w.Write(bytes)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) DeleteOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) GetOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) GetOccurrenceNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) ListNoteOccurrences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) ListOccurrences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) UpdateOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (g *Grafeas) UpdateOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
