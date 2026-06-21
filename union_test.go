package functools

import (
	"testing"
)

func TestUnion(t *testing.T) {
	var a = Range(0, 100, 1)
	var b = Range(50, 150, 1)
	var c = Union(a, b)
	t.Log(c)
	t.Fail()
}
