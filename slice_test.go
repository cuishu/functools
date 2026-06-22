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

// BenchmarkDiffBy-8   	     100	  20793552 ns/op	 1540616 B/op	     174 allocs/op
func BenchmarkDiffBy(b *testing.B) {
	m := Range(0, 10000, 1)
	n := Range(5000, 15000, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffBy(m, n, func(i int) int { return i })
	}
}

func TestDiffSeparateBy(t *testing.T) {
	a := []string{"apple", "banana", "cherry"}
	b := []string{"banana", "date", "elderberry"}
	onlyA, onlyB := DiffSeparateBy(a, b, func(s string) string { return s })
	t.Log(onlyA)
	t.Log(onlyB)
	t.Fail()
}

func BenchmarkDiffSeparateBy(b *testing.B) {
	m := Range(0, 10000, 1)
	n := Range(5000, 15000, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffSeparateBy(m, n, func(i int) int { return i })
	}
}
