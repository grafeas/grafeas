package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateNote",
		"POST",
		"/v1alpha1/projects/{projectsId}/notes",
		CreateNote,
	},

	Route{
		"CreateOccurrence",
		"POST",
		"/v1alpha1/projects/{projectsId}/occurrences",
		CreateOccurrence,
	},

	Route{
		"CreateOperation",
		"POST",
		"/v1alpha1/projects/{projectsId}/operations",
		CreateOperation,
	},

	Route{
		"DeleteNote",
		"DELETE",
		"/v1alpha1/projects/{projectsId}/notes/{notesId}",
		DeleteNote,
	},

	Route{
		"DeleteOccurrence",
		"DELETE",
		"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
		DeleteOccurrence,
	},

	Route{
		"GetNote",
		"GET",
		"/v1alpha1/projects/{projectsId}/notes/{notesId}",
		GetNote,
	},

	Route{
		"GetOccurrence",
		"GET",
		"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
		GetOccurrence,
	},

	Route{
		"GetOccurrenceNote",
		"GET",
		"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes",
		GetOccurrenceNote,
	},

	Route{
		"ListNoteOccurrences",
		"GET",
		"/v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences",
		ListNoteOccurrences,
	},

	Route{
		"ListNotes",
		"GET",
		"/v1alpha1/projects/{projectsId}/notes",
		ListNotes,
	},

	Route{
		"ListOccurrences",
		"GET",
		"/v1alpha1/projects/{projectsId}/occurrences",
		ListOccurrences,
	},

	Route{
		"UpdateNote",
		"PUT",
		"/v1alpha1/projects/{projectsId}/notes/{notesId}",
		UpdateNote,
	},

	Route{
		"UpdateOccurrence",
		"PUT",
		"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
		UpdateOccurrence,
	},

	Route{
		"UpdateOperation",
		"PUT",
		"/v1alpha1/projects/{projectsId}/operations/{operationsId}",
		UpdateOperation,
	},
}
