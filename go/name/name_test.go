package name

import (
	"testing"
)

func TestParseProject(t *testing.T) {
	tests := []struct {
		name string
		id   string
		err  bool
	}{{
		name: "",
		err:  true,
	}, {
		name: "asdf/bear-sheep",
		err:  true,
	}, {
		name: "projects/bear-sheep/notes/CVE-UH-OH/foo/bar",
		err:  true,
	}, {
		name: "projects/",
		err:  true,
	}, {
		name: "projects/bear-sheep",
		id:   "bear-sheep",
	}}

	for _, tt := range tests {
		id, err := ParseProject(tt.name)
		if err != nil {
			if !tt.err {
				t.Errorf("Got err when parsing project name %q: %v, want success", tt.name, err)
			}
		} else if tt.err {
			t.Errorf("Got success when parsing project name %q, want error", tt.name)
		}

		if id != tt.id {
			t.Errorf("Got project ID %q, want %q", id, tt.id)
		}
	}
}

func TestParseNote(t *testing.T) {
	tests := []struct {
		name string
		pID  string
		nID  string
		err  bool
	}{{
		name: "",
		err:  true,
	}, {
		name: "asdf/bear-sheep",
		err:  true,
	}, {
		name: "projects/bear-sheep/notes/CVE-UH-OH/foo/bar",
		err:  true,
	}, {
		name: "projects/",
		err:  true,
	}, {
		name: "asdf/bear-sheep/notes/CVE-UH-OH",
		err:  true,
	}, {
		name: "projects/bear-sheep/occurrences/CVE-UH-OH",
		err:  true,
	}, {
		name: "projects//notes/CVE-UH-OH",
		err:  true,
	}, {
		name: "projects/bear-sheep/notes",
		err:  true,
	}, {
		name: "projects/bear-sheep/notes/CVE-UH-OH",
		pID:  "bear-sheep",
		nID:  "CVE-UH-OH",
	}}

	for _, tt := range tests {
		pID, nID, err := ParseNote(tt.name)
		if err != nil {
			if !tt.err {
				t.Errorf("Got err when parsing note name %q: %v, want success", tt.name, err)
			}
		} else if tt.err {
			t.Errorf("Got success when parsing note name %q, want error", tt.name)
		}

		if pID != tt.pID {
			t.Errorf("Got project ID %q, want %q", pID, tt.pID)
		}
		if nID != tt.nID {
			t.Errorf("Got note ID %q, want %q", nID, tt.nID)
		}
	}
}

func TestParseOccurrence(t *testing.T) {
	tests := []struct {
		name string
		pID  string
		oID  string
		err  bool
	}{{
		name: "",
		err:  true,
	}, {
		name: "asdf/bear-sheep",
		err:  true,
	}, {
		name: "projects/bear-sheep/occurrences/1234-asdf-5678/foo/bar",
		err:  true,
	}, {
		name: "projects/",
		err:  true,
	}, {
		name: "asdf/bear-sheep/occurrences/1234-asdf-5678",
		err:  true,
	}, {
		name: "projects/bear-sheep/notes/1234-asdf-5678",
		err:  true,
	}, {
		name: "projects//occurrences/1234-asdf-5678",
		err:  true,
	}, {
		name: "projects/bear-sheep/occurrences",
		err:  true,
	}, {
		name: "projects/bear-sheep/occurrences/1234-asdf-5678",
		pID:  "bear-sheep",
		oID:  "1234-asdf-5678",
	}}

	for _, tt := range tests {
		pID, oID, err := ParseOccurrence(tt.name)
		if err != nil {
			if !tt.err {
				t.Errorf("Got err when parsing occurrence name %q: %v, want success", tt.name, err)
			}
		} else if tt.err {
			t.Errorf("Got success when parsing occurrence name %q, want error", tt.name)
		}

		if pID != tt.pID {
			t.Errorf("Got project ID %q, want %q", pID, tt.pID)
		}
		if oID != tt.oID {
			t.Errorf("Got occurrence ID %q, want %q", oID, tt.oID)
		}
	}
}

func TestFormatProject(t *testing.T) {
	tests := []struct {
		pID  string
		name string
	}{{
		pID:  "",
		name: "projects/",
	}, {
		pID:  "bear-sheep",
		name: "projects/bear-sheep",
	}}

	for _, tt := range tests {
		name := FormatProject(tt.pID)
		if name != tt.name {
			t.Errorf("Got project name %q, want %q", name, tt.name)
		}
	}
}

func TestFormatNote(t *testing.T) {
	tests := []struct {
		pID  string
		nID  string
		name string
	}{{
		pID:  "",
		nID:  "",
		name: "projects//notes/",
	}, {
		pID:  "bear-sheep",
		nID:  "",
		name: "projects/bear-sheep/notes/",
	}, {
		pID:  "",
		nID:  "CVE-UH-OH",
		name: "projects//notes/CVE-UH-OH",
	}, {
		pID:  "bear-sheep",
		nID:  "CVE-UH-OH",
		name: "projects/bear-sheep/notes/CVE-UH-OH",
	}}

	for _, tt := range tests {
		name := FormatNote(tt.pID, tt.nID)
		if name != tt.name {
			t.Errorf("Got note name %q, want %q", name, tt.name)
		}
	}
}

func TestFormatOccurrence(t *testing.T) {
	tests := []struct {
		pID  string
		oID  string
		name string
	}{{
		pID:  "",
		oID:  "",
		name: "projects//occurrences/",
	}, {
		pID:  "bear-sheep",
		oID:  "",
		name: "projects/bear-sheep/occurrences/",
	}, {
		pID:  "",
		oID:  "1234-asdf-5678",
		name: "projects//occurrences/1234-asdf-5678",
	}, {
		pID:  "bear-sheep",
		oID:  "1234-asdf-5678",
		name: "projects/bear-sheep/occurrences/1234-asdf-5678",
	}}

	for _, tt := range tests {
		name := FormatOccurrence(tt.pID, tt.oID)
		if name != tt.name {
			t.Errorf("Got occurrence name %q, want %q", name, tt.name)
		}
	}
}
