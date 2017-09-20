package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"errors"
	"fmt"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
)

func TestGrafeas_CreateNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := note()
	if err := createNote(n, g); err != nil {
		t.Errorf("%v", err)
	}
}

func TestGrafeas_CreateOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := note()
	if err := createNote(n, g); err != nil {
		t.Errorf("Error creating note: %v")
	}
	o := occurrence(n.Name)
	if err := createOccurrence(o, g); err != nil {
		t.Errorf("%v", err)
	}
}

func TestGrafeas_CreateOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	o := operation()
	if err := createOperation(o, g); err != nil {
		t.Errorf("%v", err)
	}
}

func createOccurrence(o swagger.Occurrence, g Grafeas) error {
	rawOcc, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOcc)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/test-project/occurrences", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOccurrence(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateOccurrence(%v) got %v want 200", o, w.Code))
	}
	return nil
}

func createNote(n swagger.Note, g Grafeas) error {
	rawNote, err := json.Marshal(&n)
	reader := bytes.NewReader(rawNote)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/vulnerability-scanner-a/notes?note_id=CVE-1999-0710", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateNote(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateNote(%v) got %v want 200", n, w.Code))
	}
	return nil
}

func createOperation(o swagger.Operation, g Grafeas) error {
	rawOp, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOp)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/vulnerability-scanner-a/operations", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOperation(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateNote(%v) got %v want 200", o, w.Code))
	}
	return nil
}

func operation() swagger.Operation {
	return swagger.Operation{
		Name:     "projects/vulnerability-scanner-a/operations/foo",
		Metadata: map[string]string{"StartTime": "0916162344"},
		Done:     true,
	}
}

func occurrence(noteName string) swagger.Occurrence {
	return swagger.Occurrence{
		Name:        "projects/test-project/occurrences/134",
		ResourceUrl: "gcr.io/foo/bar",
		NoteName:    noteName,
		Kind:        "PACKAGE_VULNERABILITY",
		VulnerabilityDetails: swagger.VulnerabilityDetails{
			Severity:  "HIGH",
			CvssScore: 7.5,
			PackageIssue: []swagger.PackageIssue{
				swagger.PackageIssue{
					SeverityName: "HIGH",
					AffectedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "52.1",
							Revision: "8+deb8u3",
						},
					},
					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
			},
		},
	}
}

func note() swagger.Note {
	return swagger.Note{
		Name:             "projects/vulnerability-scanner-a/notes/CVE-1999-0710",
		ShortDescription: "CVE-2014-9911",
		LongDescription:  "NIST vectors: AV:N/AC:L/Au:N/C:P/I:P",
		Kind:             "PACKAGE_VULNERABILITY",
		VulnerabilityType: swagger.VulnerabilityType{
			CvssScore: 7.5,
			Severity:  "HIGH",
			Details: []swagger.Detail{
				swagger.Detail{
					CpeUri:  "cpe:/o:debian:debian_linux:7",
					Package: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:7",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "4.8.1.1",
							Revision: "12+deb7u6",
						},
					},
				},
				swagger.Detail{
					CpeUri:  "cpe:/o:debian:debian_linux:8",
					Package: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:8",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "52.1",
							Revision: "8+deb8u4",
						},
					},
				},
				swagger.Detail{
					CpeUri:  "cpe:/o:debian:debian_linux:9",
					Package: "icu",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "HIGH",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:debian:debian_linux:9",
						Package_: "icu",
						Version: swagger.Version{
							Name:     "55.1",
							Revision: "3",
						},
					},
				},
				swagger.Detail{
					CpeUri:  "cpe:/o:canonical:ubuntu_linux:14.04",
					Package: "andriod",
					Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
						"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
						"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
					MinAffectedVersion: swagger.Version{
						Kind: "MINIMUM",
					},
					SeverityName: "MEDIUM",

					FixedLocation: swagger.VulnerabilityLocation{
						CpeUri:   "cpe:/o:canonical:ubuntu_linux:14.04",
						Package_: "andriod",
						Version: swagger.Version{
							Kind: "MAXIMUM",
						},
					},
				},
			},
		},
		RelatedUrl: []swagger.RelatedUrl{
			swagger.RelatedUrl{
				Url:   "https://security-tracker.debian.org/tracker/CVE-2014-9911",
				Label: "More Info",
			},
			swagger.RelatedUrl{
				Url:   "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9911",
				Label: "More Info",
			},
		},
	}
}
