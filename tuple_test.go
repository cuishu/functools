package functools

import (
	"testing"
)

func TestTuple(t *testing.T) {
	tuple := NewTuple(1, "a")
	if tuple.V1 != 1 || tuple.V2 != "a" {
		t.Fail()
	}
}

func TestTuple1(t *testing.T) {
	tuple := NewTuple(1, "a")
	a, b := tuple.Values()
	if a != 1 || b != "a" {
		t.Fail()
	}
}
