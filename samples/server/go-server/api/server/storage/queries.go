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
		CREATE TABLE IF NOT EXISTS project (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE
		);
		CREATE TABLE IF NOT EXISTS note (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			data TEXT
		);
		CREATE TABLE IF NOT EXISTS occurrence (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			data TEXT,
			note_id int REFERENCES note NOT NULL
		);
		CREATE TABLE IF NOT EXISTS operation (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			data TEXT
		);`

	insertProject = `INSERT INTO project(name) VALUES ($1)`
	projectExists = `SELECT EXISTS (SELECT 1 FROM project WHERE name = $1)`
	deleteProject = `DELETE FROM project WHERE name = $1`
	listProjects  = `SELECT name FROM project`

	insertOccurrence = `INSERT INTO occurrence(name, note_id, data) VALUES ($1, (SELECT id FROM note WHERE name = $2), $3)`
	searchOccurrence = `SELECT data FROM occurrence WHERE name = $1`
	updateOccurrence = `UPDATE occurrence SET data = $2 WHERE name = $1`
	deleteOccurrence = `DELETE FROM occurrence WHERE name = $1`
	listOccurrences  = `SELECT data FROM occurrence WHERE name LIKE $1 || '%'`

	insertNote          = `INSERT INTO note(name, data) VALUES ($1, $2)`
	searchNote          = `SELECT data FROM note WHERE name = $1`
	updateNote          = `UPDATE note SET data = $2 WHERE name = $1`
	deleteNote          = `DELETE FROM note WHERE name = $1`
	listNotes           = `SELECT data FROM note WHERE name LIKE $1 || '%'`
	listNoteOccurrences = `SELECT o.data FROM occurrence as o, note as n WHERE n.id = o.note_id AND n.name = $1`

	insertOperation = `INSERT INTO operation(name, data) VALUES ($1, $2)`
	searchOperation = `SELECT data FROM operation WHERE name = $1`
	deleteOperation = `DELETE FROM operation WHERE name = $1`
	updateOperation = `UPDATE operation SET data = $2 WHERE name = $1`
	listOperations  = `SELECT data FROM operation WHERE name LIKE $1 || '%'`
)
