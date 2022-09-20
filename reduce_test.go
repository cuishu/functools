package functools

import (
	"testing"
)

func TestReduce(t *testing.T) {
	var a []int
	for i := 1; i <= 100; i++ {
		a = append(a, i)
	}
	n := Reduce(func(x, y int) int {
		return x + y
	}, a, 0)
	if n != 5050 {
		t.FailNow()
	}
}

func BenchmarkReduce(b *testing.B) {
	data := Range(1, 101, 1)
	for i := 0; i < b.N; i++ {
		Reduce(func(x, y int) int {
			return x + y
		}, data, 0)
	}
}
