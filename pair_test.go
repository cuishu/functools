package functools

import (
	"testing"
)

func TestTuple(t *testing.T) {
	tuple := NewPair(1, "a")
	if tuple.First != 1 || tuple.Second != "a" {
		t.Fail()
	}
}

func TestTuple1(t *testing.T) {
	tuple := NewPair(1, "a")
	a, b := tuple.Values()
	if a != 1 || b != "a" {
		t.Fail()
	}
}
