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

package storage

const (
	createTables = `
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			data TEXT
		);
		CREATE TABLE IF NOT EXISTS occurrences (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			data TEXT,
			note_id int REFERENCES notes NOT NULL
		);
		CREATE TABLE IF NOT EXISTS operations (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			data TEXT
		);`

	insertProject = `INSERT INTO projects(name) VALUES ($1)`
	projectExists = `SELECT EXISTS (SELECT 1 FROM projects WHERE name = $1)`
	deleteProject = `DELETE FROM projects WHERE name = $1`
	listProjects  = `SELECT name FROM projects`

	insertOccurrence = `INSERT INTO occurrences(name, note_id, data) VALUES ($1, (SELECT id FROM notes WHERE name = $2), $3)`
	searchOccurrence = `SELECT data FROM occurrences WHERE name = $1`
	updateOccurrence = `UPDATE occurrences SET data = $2 WHERE name = $1`
	deleteOccurrence = `DELETE FROM occurrences WHERE name = $1`
	listOccurrences  = `SELECT data FROM occurrences WHERE name LIKE $1 || '%'`

	insertNote          = `INSERT INTO notes(name, data) VALUES ($1, $2)`
	searchNote          = `SELECT data FROM notes WHERE name = $1`
	updateNote          = `UPDATE notes SET data = $2 WHERE name = $1`
	deleteNote          = `DELETE FROM notes WHERE name = $1`
	listNotes           = `SELECT data FROM notes WHERE name LIKE $1 || '%'`
	listNoteOccurrences = `SELECT o.data FROM occurrences as o, notes as n WHERE n.id = o.note_id AND n.name = $1`

	insertOperation = `INSERT INTO operations(name, data) VALUES ($1, $2)`
	searchOperation = `SELECT data FROM operations WHERE name = $1`
	deleteOperation = `DELETE FROM operations WHERE name = $1`
	updateOperation = `UPDATE operations SET data = $2 WHERE name = $1`
	listOperations  = `SELECT data FROM operations WHERE name LIKE $1 || '%'`
)
