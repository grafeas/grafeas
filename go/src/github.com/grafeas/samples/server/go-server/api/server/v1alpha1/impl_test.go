package v1alpha1

import (
	"github.com/grafeas/samples/server/go-server/api"
	"github.com/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/samples/server/go-server/api/server/storage"
	"github.com/grafeas/samples/server/go-server/api/server/testing"
	"net/http"
	"reflect"
	"testing"
	"fmt"
)

func TestGrafeas_CreateOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	op := swagger.Operation{}
	if err := g.CreateOperation(&op); err == nil {
		t.Error("CreateOperation(empty operation): got success, want error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOperation(empty operation): got %v, want BadRequest(400)", err.StatusCode)
	}
	op = testutil.Operation()
	if err := g.CreateOperation(&op); err != nil {
		t.Errorf("CreateOperation(%v) got %v, want success", op, err)
	}
}

func TestGrafeas_CreateOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	o := swagger.Occurrence{}
	if err := g.CreateOccurrence(&o); err == nil {
		t.Error("CreateOccurrence(empty occ): got success, want error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOccurrence(empty occ): got %v, want BadRequest(400)", err.StatusCode)
	}
	o = testutil.Occurrence(n.Name)
	if err := g.CreateOccurrence(&o); err != nil {
		t.Errorf("CreateOccurrence(%v) got %v, want success", n, err)
	}

	// Try to insert an occurrence for a note that does not exist.
	o.Name = "projects/testproject/occurrences/nonote"
	o.NoteName = "projects/scan-provider/notes/notthere"
	if err := g.CreateOccurrence(&o); err == nil {
		t.Errorf("CreateOccurrence got success, want Error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOccurrence got code %v want %v", err.StatusCode, http.StatusBadRequest)
	}

}

func TestGrafeas_CreateNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := swagger.Note{}
	if err := g.CreateNote(&n); err == nil {
		t.Error("CreateNote(empty note): got success, want error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateNote(empty note): got %v, want %v", err.StatusCode, http.StatusBadRequest)
	}
	n = testutil.Note()
	if err := g.CreateNote(&n); err != nil {
		t.Errorf("CreateNote(%v) got %v, want success", n, err)
	}
}

func TestGrafeas_DeleteNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if err := g.DeleteNote(pID, nID); err == nil {
		t.Error("DeleteNote that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.DeleteNote(pID, nID); err != nil {
		t.Errorf("DeleteNote  got %v, want success", err)
	}
}

func TestGrafeas_DeleteOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	// CreateNote so we can create an occurrence
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	o := testutil.Occurrence(n.Name)
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if err := g.DeleteNote(pID, oID); err == nil {
		t.Error("DeleteNote that doesn't exist got success, want err")
	}
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.DeleteOccurrence(pID, oID); err != nil {
		t.Errorf("DeleteNote  got %v, want success", err)
	}
}

func TestGrafeas_DeleteOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	o := testutil.Operation()
	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if err := g.DeleteOperation(pID, oID); err == nil {
		t.Error("DeleteOperation that doesn't exist got success, want err")
	}
	if err := g.CreateOperation(&o); err != nil {
		t.Fatalf("CreateOperation(%v) got %v, want success", o, err)
	}
	if err := g.DeleteOperation(pID, oID); err != nil {
		t.Errorf("DeleteOperation  got %v, want success", err)
	}
}

func TestGrafeas_GetNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if _, err := g.GetNote(pID, nID); err == nil {
		t.Error("GetNote that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if got, err := g.GetNote(pID, nID); err != nil {
		t.Fatalf("GetNote(%v) got %v, want success", n, err)
	} else if n.Name != got.Name || !reflect.DeepEqual(n.VulnerabilityType, got.VulnerabilityType) {
		t.Errorf("GetNote got %v, want %v", *got, n)
	}
}

func TestGrafeas_GetOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()

	o := testutil.Occurrence(n.Name)

	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if _, err := g.GetOccurrence(pID, oID); err == nil {
		t.Error("GetOccurrence that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence(%v) got %v, want success", n, err)
	}
	if got, err := g.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence(%v) got %v, want success", o, err)
	} else if o.Name != got.Name || !reflect.DeepEqual(o.VulnerabilityDetails, got.VulnerabilityDetails) {
		t.Errorf("GetOccurrence got %v, want %v", *got, o)
	}
}

func TestGrafeas_GetOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	o := testutil.Operation()
	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if _, err := g.GetOperation(pID, oID); err == nil {
		t.Error("GetOperation that doesn't exist got success, want err")
	}
	if err := g.CreateOperation(&o); err != nil {
		t.Fatalf("CreateOperation(%v) got %v, want success", o, err)
	}
	if got, err := g.GetOperation(pID, oID); err != nil {
		t.Fatalf("GetOperation(%v) got %v, want success", o, err)
	} else if o.Name != got.Name || !reflect.DeepEqual(*got, o) {
		t.Errorf("GetOperation got %v, want %v", *got, o)
	}
}

func TestGrafeas_GetOccurrenceNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()

	o := testutil.Occurrence(n.Name)

	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if _, err := g.GetOccurrenceNote(pID, oID); err == nil {
		t.Error("GetOccurrenceNote that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence(%v) got %v, want success", n, err)
	}
	if got, err := g.GetOccurrenceNote(pID, oID); err != nil {
		t.Fatalf("GetOccurrenceNote(%v) got %v, want success", n, err)
	} else if n.Name != got.Name || !reflect.DeepEqual(n.VulnerabilityType, got.VulnerabilityType) {
		t.Errorf("GetOccurrenceNote got %v, want %v", *got, n)
	}
}

func TestUpdateNote (t *testing.T) {
	// Update Note that doesn't exist
	updateDesc := "this is a new description"
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	update := testutil.Note()
	update.LongDescription = updateDesc
	if _, err := g.UpdateNote(pID, nID, &update); err == nil {
		t.Error("UpdateNote that doesn't exist got success, want err")
	}

	// Actually create note
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}

	// Update Note name and fail
	update.Name = "New name"
	if _, err := g.UpdateNote(pID, nID, &update); err == nil {
		t.Error("UpdateNote that with name change got success, want err")
	}

	// Update Note and verify that update worked.
	update = testutil.Note()
	update.LongDescription = updateDesc
	if got, err := g.UpdateNote(pID, nID, &update); err != nil {
		t.Errorf("UpdateNote got %v, want success", err)
	} else if updateDesc != update.LongDescription {
		t.Errorf("UpdateNote got %v, want %v",
			got.LongDescription, updateDesc)
	}
	if got, err := g.GetNote(pID, nID); err != nil {
		t.Fatalf("GetNote(%v) got %v, want success", n, err)
	} else if updateDesc != got.LongDescription  {
		t.Errorf("GetNote got %v, want %v", got.LongDescription, updateDesc)
	}
}

func TestUpdateOccurrence(t *testing.T) {
	// Update occurrence that doesn't exist
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	o := testutil.Occurrence(n.Name)

	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if _, err := g.UpdateOccurrence(pID, oID, &o); err == nil {
		t.Error("UpdateOccurrence that doesn't exist got success, want err")
	}
	// Create an occurrence to update
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence(%v) got %v, want success", n, err)
	}
	// update occurrence name
	update := testutil.Occurrence(n.Name)
	update.Name = "New name"
	if _, err := g.UpdateOccurrence(pID, oID, &update); err == nil {
		t.Error("UpdateOccurrence that with name change got success, want err")
	}

	// update note name to a note that doesn't exist
	update = testutil.Occurrence("projects/p/notes/bar")
	if _, err := g.UpdateOccurrence(pID, oID, &update); err == nil {
		t.Error("UpdateOccurrence that with note name that doesn't exist" +
			" got success, want err")
	}

	// update note name to a note that does exist
	n = testutil.Note()
	newName := fmt.Sprintf("%v-new", n.Name)
	n.Name = newName
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	update = testutil.Occurrence(n.Name)
	if got, err := g.UpdateOccurrence(pID, oID, &update);  err != nil {
		t.Errorf("UpdateOccurrence got %v, want success", err)
	} else if n.Name != got.NoteName {
		t.Errorf("UpdateOccurrence got %v, want %v",
			got.NoteName, n.Name)
	}
	if got, err := g.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence(%v) got %v, want success", n, err)
	} else if n.Name != got.NoteName  {
		t.Errorf("GetOccurrence got %v, want %v",
			got.NoteName, n.Name)
	}
}