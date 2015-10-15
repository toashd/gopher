package gopher

import (
	"reflect"
	"testing"
)

// TestNew verifies that the returned instance is of the proper type.
func TestNew(t *testing.T) {
	g := New()
	if reflect.TypeOf(g).String() != "*gopher.Gopher" {
		t.Error("New returned incorrect type")
	}
}
