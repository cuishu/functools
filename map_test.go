package functools

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	a = Map(a, func(i int) int {
		return i + 1
	})
	for i := 0; i < 100; i++ {
		if a[i] != i+1 {
			t.FailNow()
		}
	}
}

func BenchmarkMap(b *testing.B) {
	a := Range(0, 100, 1)
	for i := 0; i < b.N; i++ {
		Map(a, func(i int) int {
			return i + 1
		})
	}
}

func TestForEach(t *testing.T) {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	ForEach(a, func(i int) {
		fmt.Println(i)
	})
}
