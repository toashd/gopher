package gopher

import (
	"reflect"
	"testing"
)

// TestNew verifies that the returned instance is of the proper type.
func TestNew(t *testing.T) {
	s := New()
	if reflect.TypeOf(s).String() != "*gopher.Gopher" {
		t.Error("New returned incorrect type")
	}
}
