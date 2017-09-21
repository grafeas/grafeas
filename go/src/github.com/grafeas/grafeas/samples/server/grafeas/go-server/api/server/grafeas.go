package server

import (
	"encoding/json"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/v1alpha1"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	g v1alpha1.Grafeas
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	n := swagger.Note{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &n)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := h.g.CreateNote(&n); err != nil {

		log.Printf("Error creating note: %v", err)
		http.Error(w, err.Err, err.StatusCode)
	}
	bytes, err := json.Marshal(&n)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *Handler) CreateOccurrence(w http.ResponseWriter, r *http.Request) {
	o := swagger.Occurrence{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &o)
	if err := h.g.CreateOccurrence(&o); err != nil {
		log.Printf("Error creating occurrence: %v", err)
		http.Error(w, err.Err, err.StatusCode)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	bytes, err := json.Marshal(&o)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h *Handler) CreateOperation(w http.ResponseWriter, r *http.Request) {
	o := swagger.Operation{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &o)
	if err := h.g.CreateOperation(&o); err != nil {
		log.Printf("Error creating occurrence: %v", err)
		http.Error(w, err.Err, err.StatusCode)
	}

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

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOccurrenceNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListNoteOccurrences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListOccurrences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
