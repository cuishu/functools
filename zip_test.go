package functools

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	var b = []string{"a", "b", "c", "d", "e"}

	data := Zip(a, b)
	for _, tuple := range data {
		x, y := tuple.Values()
		fmt.Println(x, y)
	}
}

func TestZip1(t *testing.T) {
	var a = []int{1, 2, 3}
	var b = []string{"a", "b", "c", "d", "e"}

	data := Zip(a, b)
	for _, tuple := range data {
		x, y := tuple.Values()
		fmt.Println(x, y)
	}
}

func BenchmarkZip(b *testing.B) {
	data := Range(0, 100, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Zip(data, data)
	}
}
