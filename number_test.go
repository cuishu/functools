package functools

import (
	"testing"
)

func TestSum(t *testing.T) {
	var a []int
	for i := 1; i <= 100; i++ {
		a = append(a, i)
	}

	n := Sum(a)
	if n != 5050 {
		t.Fail()
	}
}

func BenchmarkSum(b *testing.B) {
	data := Range(1, 101, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data)
	}
}

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Fail()
	}

	if Min(2, 1) != 1 {
		t.Fail()
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(1, 2)
	}
}

func TestMax(t *testing.T) {
	if Max(3, 4) != 4 {
		t.Fail()
	}

	if Max(4, 3) != 4 {
		t.Fail()
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(1, 2)
	}
}

func TestRange(t *testing.T) {
	data := Range(0, 100, 1)
	if len(data) != 100 {
		t.Fail()
	}
	data = Range(0, 100, 2)
	if len(data) != 50 {
		t.Fail()
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Range(1, 100, 1)
	}
}
