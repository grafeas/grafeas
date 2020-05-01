package storage_test

import (
	"testing"
	"fmt"

	"github.com/grafeas/grafeas/go/v1beta1/storage"
)

var myFilter storage.MysqlFilterSql

func TestParseFilter(t *testing.T) {
	filter := `note_name="test_note_1"`
	actual := myFilter.ParseFilter(filter)
	expected := `(data->'$.note_name' = "test_note_1")`
	fmt.Println(actual)	
	if actual != expected {
		t.Errorf("Expecting: " + expected + "\nGet: " + actual)
	}
} 

