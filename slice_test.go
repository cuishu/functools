package functools

import (
	"testing"
)

func TestSlice(t *testing.T) {
	var list []int
	for i := 0; i < 100; i++ {
		list = append(list, i)
	}
	list = Reverse(list)
	t.Log(list)
}

func TestReverseInPlace(t *testing.T) {
	var list []int
	for i := 0; i < 100; i++ {
		list = append(list, i)
	}
	t.Log(list)
	ReverseInPlace(list)
	t.Log(list)
}
