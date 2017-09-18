package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(s storage.Store) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := Grafeas{s}
	for _, route := range routes(api) {
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

func routes(api Grafeas) Routes {
	return Routes{
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
			api.CreateNote,
		},

		Route{
			"CreateOccurrence",
			"POST",
			"/v1alpha1/projects/{projectsId}/occurrences",
			api.CreateOccurrence
		},

		Route{
			"CreateOperation",
			"POST",
			"/v1alpha1/projects/{projectsId}/operations",
			api.CreateOperation,
		},

		Route{
			"DeleteNote",
			"DELETE",
			"/v1alpha1/projects/{projectsId}/notes/{notesId}",
			api.DeleteNote,
		},

		Route{
			"DeleteOccurrence",
			"DELETE",
			"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
			api.DeleteOccurrence,
		},

		Route{
			"GetNote",
			"GET",
			"/v1alpha1/projects/{projectsId}/notes/{notesId}",
			api.GetNote,
		},

		Route{
			"GetOccurrence",
			"GET",
			"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
			api.GetOccurrence,
		},

		Route{
			"GetOccurrenceNote",
			"GET",
			"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes",
			api.GetOccurrenceNote,
		},

		Route{
			"ListNoteOccurrences",
			"GET",
			"/v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences",
			api.ListNoteOccurrences,
		},

		Route{
			"ListNotes",
			"GET",
			"/v1alpha1/projects/{projectsId}/notes",
			api.ListNotes,
		},

		Route{
			"ListOccurrences",
			"GET",
			"/v1alpha1/projects/{projectsId}/occurrences",
			api.ListOccurrences,
		},

		Route{
			"UpdateNote",
			"PUT",
			"/v1alpha1/projects/{projectsId}/notes/{notesId}",
			api.UpdateNote,
		},

		Route{
			"UpdateOccurrence",
			"PUT",
			"/v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
			api.UpdateOccurrence,
		},

		Route{
			"UpdateOperation",
			"PUT",
			"/v1alpha1/projects/{projectsId}/operations/{operationsId}",
			api.UpdateOperation,
		},
	}
}
