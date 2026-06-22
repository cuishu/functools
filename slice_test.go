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

func TestDiff(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"banana", "date", "elderberry"}
	diff := Diff(a, b)
	t.Log(diff)
	t.Fail()
}
func TestDiffSeparate(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"banana", "date", "elderberry"}
	onlyA, onlyB := DiffSeparate(a, b)
	t.Log(onlyA)
	t.Log(onlyB)
	t.Fail()
}

func TestDiffBy(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"banana", "date", "elderberry"}
	diff := DiffBy(a, b, func(s string) string { return s })
	t.Log(diff)
	t.Fail()
}

func TestDiffSeparateBy(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"banana", "date", "elderberry"}
	onlyA, onlyB := DiffSeparateBy(a, b, func(s string) string { return s })
	t.Log(onlyA)
	t.Log(onlyB)
	t.Fail()
}
