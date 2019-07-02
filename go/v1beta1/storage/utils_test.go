package storage_test

import (
	"testing"

	cpb "github.com/grafeas/grafeas/proto/v1beta1/common_go_proto"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"

	"github.com/google/go-cmp/cmp"
	"github.com/grafeas/grafeas/go/v1beta1/storage"
	"google.golang.org/genproto/protobuf/field_mask"
)

var (
	opts = cmp.FilterPath(
		func(p cmp.Path) bool {
			ignoreCreate := p.String() == "CreateTime"
			ignoreUpdate := p.String() == "UpdateTime"
			return ignoreCreate || ignoreUpdate
		}, cmp.Ignore())
)

func TestCreateFieldMask(t *testing.T) {
	pathTests := []struct {
		testName string
		paths    []string
	}{
		{
			testName: "Test non-zero number of paths",
			paths:    []string{"Hello.test", "a.b.v.f"},
		},
		{
			testName: "Test zero number of paths",
			paths:    []string{},
		},
		{
			testName: "Test nil paths",
			paths:    nil,
		},
	}

	for _, test := range pathTests {
		t.Run(test.testName, func(t *testing.T) {
			var pathsToTest []string
			outputMask := storage.CreateFieldMask(test.paths)

			if test.paths != nil {
				pathsToTest = test.paths
			} else {
				pathsToTest = []string{}
			}

			if diff := cmp.Diff(outputMask.GetPaths(), pathsToTest, nil); diff != "" {
				t.Errorf("CreateFieldMask returned diff (want -> got):\n%s", diff)
			}
		})
	}
}

func TestApplyUpdateOnOccurrence(t *testing.T) {
	occurrenceTests := []struct {
		testName    string
		inputOcc    *pb.Occurrence
		updateOcc   *pb.Occurrence
		updateMask  *field_mask.FieldMask
		expectedOcc *pb.Occurrence
	}{
		{
			testName: "Test single valid update",
			inputOcc: storage.TestOccurrence("project1", "CVE-2014-9911"),
			updateOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project1", "CVE-2014-9911")
				o.GetVulnerability().CvssScore = 1.5
				return o
			}(),
			updateMask: storage.CreateFieldMask([]string{"Details.Vulnerability.CvssScore"}),
			expectedOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project1", "CVE-2014-9911")
				o.GetVulnerability().CvssScore = 1.5
				return o
			}(),
		},
		{
			testName: "Test multiple valid updates",
			inputOcc: storage.TestOccurrence("project2", "CVE-2014-9911"),
			updateOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project2", "CVE-2014-9911")
				o.GetVulnerability().CvssScore = 1.5
				o.Kind = cpb.NoteKind_IMAGE
				return o
			}(),
			updateMask: storage.CreateFieldMask([]string{"Details.Vulnerability.CvssScore", "Kind"}),
			expectedOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project2", "CVE-2014-9911")
				o.GetVulnerability().CvssScore = 1.5
				o.Kind = cpb.NoteKind_IMAGE
				return o
			}(),
		},
		{
			testName:    "Test single invalid update",
			inputOcc:    storage.TestOccurrence("project3", "CVE-2014-9911"),
			updateOcc:   storage.TestOccurrence("project3", "CVE-2014-9911"),
			updateMask:  storage.CreateFieldMask([]string{"iAmInvalid"}),
			expectedOcc: storage.TestOccurrence("project3", "CVE-2014-9911"),
		},
		{
			testName: "Test 1 valid and 1 invalid update",
			inputOcc: storage.TestOccurrence("project4", "CVE-2014-9911"),
			updateOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project4", "CVE-2014-9911")
				o.Kind = cpb.NoteKind_IMAGE
				return o
			}(),
			updateMask: storage.CreateFieldMask([]string{"iAmInvalid.invalid", "Kind"}),
			expectedOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project4", "CVE-2014-9911")
				o.Kind = cpb.NoteKind_IMAGE
				return o
			}(),
		},
		{
			testName:    "Test no update",
			inputOcc:    storage.TestOccurrence("project5", "CVE-2014-9911"),
			updateOcc:   storage.TestOccurrence("project5", "CVE-2014-9911"),
			updateMask:  storage.CreateFieldMask([]string{}),
			expectedOcc: storage.TestOccurrence("project5", "CVE-2014-9911"),
		},
		{
			testName: "Test single valid update with updates not in mask",
			inputOcc: storage.TestOccurrence("project6", "CVE-2014-9911"),
			updateOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project6", "CVE-2015-1199")
				o.GetVulnerability().CvssScore = 1.5
				return o
			}(),
			updateMask: storage.CreateFieldMask([]string{"Details.Vulnerability.CvssScore"}),
			expectedOcc: func() *pb.Occurrence {
				o := storage.TestOccurrence("project6", "CVE-2014-9911")
				o.GetVulnerability().CvssScore = 1.5
				return o
			}(),
		},
	}

	for _, test := range occurrenceTests {
		t.Run(test.testName, func(t *testing.T) {
			testOcc, err := storage.ApplyUpdateOnOccurrence(test.inputOcc, test.updateOcc, test.updateMask)
			if err != nil {
				t.Fatalf("ApplyUpdateOnOccurrence got %v, want success", err)
			}
			if diff := cmp.Diff(testOcc, test.expectedOcc, opts); diff != "" {
				t.Errorf("ApplyUpdateOnOccurrence returned diff (want -> got):\n%s", diff)
			}
		})
	}
}

func TestApplyUpdateOnNote(t *testing.T) {
	noteTests := []struct {
		testName     string
		inputNote    *pb.Note
		updateNote   *pb.Note
		updateMask   *field_mask.FieldMask
		expectedNote *pb.Note
	}{
		{
			testName:  "Test single valid update",
			inputNote: storage.TestNote("project1"),
			updateNote: func() *pb.Note {
				n := storage.TestNote("project1")
				n.GetVulnerability().CvssScore = 1.5
				return n
			}(),
			updateMask: storage.CreateFieldMask([]string{"Type.Vulnerability.CvssScore"}),
			expectedNote: func() *pb.Note {
				n := storage.TestNote("project1")
				n.GetVulnerability().CvssScore = 1.5
				return n
			}(),
		},
		{
			testName:  "Test multiple valid updates",
			inputNote: storage.TestNote("project2"),
			updateNote: func() *pb.Note {
				n := storage.TestNote("project2")
				n.GetVulnerability().CvssScore = 1.5
				n.Kind = cpb.NoteKind_PACKAGE
				return n
			}(),
			updateMask: storage.CreateFieldMask([]string{"Type.Vulnerability.CvssScore", "Kind"}),
			expectedNote: func() *pb.Note {
				n := storage.TestNote("project2")
				n.GetVulnerability().CvssScore = 1.5
				n.Kind = cpb.NoteKind_PACKAGE
				return n
			}(),
		},
		{
			testName:     "Test single invalid update",
			inputNote:    storage.TestNote("project3"),
			updateNote:   storage.TestNote("project3"),
			updateMask:   storage.CreateFieldMask([]string{"iAmInvalid"}),
			expectedNote: storage.TestNote("project3"),
		},
		{
			testName:  "Test 1 valid and 1 invalid update",
			inputNote: storage.TestNote("project4"),
			updateNote: func() *pb.Note {
				n := storage.TestNote("project4")
				n.Kind = cpb.NoteKind_DEPLOYMENT
				return n
			}(),
			updateMask: storage.CreateFieldMask([]string{"iAmInvalid.invalid", "Kind"}),
			expectedNote: func() *pb.Note {
				n := storage.TestNote("project4")
				n.Kind = cpb.NoteKind_DEPLOYMENT
				return n
			}(),
		},
		{
			testName:     "Test no update",
			inputNote:    storage.TestNote("project5"),
			updateNote:   storage.TestNote("project5"),
			updateMask:   storage.CreateFieldMask([]string{}),
			expectedNote: storage.TestNote("project5"),
		},
		{
			testName:  "Test single valid update with updates not in mask",
			inputNote: storage.TestNote("project6"),
			updateNote: func() *pb.Note {
				n := storage.TestNote("project6")
				n.GetVulnerability().CvssScore = 1.5
				n.GetVulnerability().Severity = 342
				return n
			}(),
			updateMask: storage.CreateFieldMask([]string{"Type.Vulnerability.CvssScore"}),
			expectedNote: func() *pb.Note {
				n := storage.TestNote("project6")
				n.GetVulnerability().CvssScore = 1.5
				return n
			}(),
		},
	}

	for _, test := range noteTests {
		t.Run(test.testName, func(t *testing.T) {
			testNote, err := storage.ApplyUpdateOnNote(test.inputNote, test.updateNote, test.updateMask)
			if err != nil {
				t.Fatalf("ApplyUpdateOnNote got %v, want success", err)
			}
			if diff := cmp.Diff(testNote, test.expectedNote, opts); diff != "" {
				t.Errorf("ApplyUpdateOnNote returned diff (want -> got):\n%s", diff)
			}
		})
	}
}
