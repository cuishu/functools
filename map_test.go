package functools

import (
	"testing"
)

func TestMap(t *testing.T) {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	a = Map(func(i int) int {
		return i + 1
	}, a)
	for i := 0; i < 100; i++ {
		if a[i] != i+1 {
			t.FailNow()
		}
	}
}

func BenchmarkMap(b *testing.B) {
	a := Range(0, 100, 1)
	for i := 0; i < b.N; i++ {
		Map(func(i int) int {
			return i + 1
		}, a)
	}
}
