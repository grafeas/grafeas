// Copyright 2017 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/v1alpha1"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(g v1alpha1.Grafeas) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := Handler{g}
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

func routes(api Handler) Routes {
	return Routes{
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
			api.CreateOccurrence,
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
			"DeleteOperation",
			"DELETE",
			"/v1alpha1/projects/{projectsId}/operations/{operationId}",
			api.DeleteOperation,
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
			"GetOperation",
			"GET",
			"/v1alpha1/projects/{projectsId}/operations/{operationsId}",
			api.GetOperation,
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
			"ListOperations",
			"GET",
			"/v1alpha1/projects/{projectsId}/operations",
			api.ListOperations,
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
