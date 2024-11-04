package functools

import (
	"testing"
)

func TestUnique(t *testing.T) {
	var a = Range(0, 100, 1)

	if len(Unique(a)) != 100 {
		t.Fail()
	}

	var b = []string{"a", "a", "b", "c", "c", "c"}
	if len(Unique(b)) != 3 {
		t.Fail()
	}
}
