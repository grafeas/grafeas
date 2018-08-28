// Package name deals with parsing and formatting resource names used in the Grafeas API.
package name

import (
	"fmt"
	"strings"

	"github.com/grafeas/grafeas/go/errors"
	"google.golang.org/grpc/codes"
)

// ParseProject parses the project ID from a project resource name.
func ParseProject(name string) (string, error) {
	params := strings.Split(name, "/")
	if len(params) != 2 {
		return "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]', got %q", name)
	}
	if params[0] != "projects" {
		return "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]', got %q", name)
	}
	if params[1] == "" {
		return "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]', got %q", name)
	}

	return params[1], nil
}

// ParseNote parses the project ID and note ID from a note resource name.
func ParseNote(name string) (string, string, error) {
	params := strings.Split(name, "/")
	if len(params) != 4 {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/notes/[NOTE_ID]', got %q", name)
	}
	if params[0] != "projects" {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/notes/[NOTE_ID]', got %q", name)
	}
	if params[2] != "notes" {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/notes/[NOTE_ID]', got %q", name)
	}
	if params[1] == "" || params[3] == "" {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/notes/[NOTE_ID]', got %q", name)
	}

	return params[1], params[3], nil
}

// ParseOccurrence parses the project ID and occurrence ID from an occurrence resource name.
func ParseOccurrence(name string) (string, string, error) {
	params := strings.Split(name, "/")
	if len(params) != 4 {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]', got %q", name)
	}
	if params[0] != "projects" {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]', got %q", name)
	}
	if params[2] != "occurrences" {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]', got %q", name)
	}
	if params[1] == "" || params[3] == "" {
		return "", "", errors.Newf(codes.InvalidArgument, "name must be in the form 'projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]', got %q", name)
	}

	return params[1], params[3], nil
}

// FormatProject formats the specified project ID into a project resource name.
func FormatProject(pID string) string {
	return fmt.Sprintf("projects/%s", pID)
}

// FormatNote formats the specified project ID and note ID into a note resource name.
func FormatNote(pID, nID string) string {
	return fmt.Sprintf("projects/%s/notes/%s", pID, nID)
}

// FormatOccurrence formats the specified project ID and occurrence ID into an occurrence resource
// name.
func FormatOccurrence(pID, oID string) string {
	return fmt.Sprintf("projects/%s/occurrences/%s", pID, oID)
}
