package functools

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var a = Range(0, 100, 1)
	a = Filter(a, func(i int) bool {
		return i < 50
	})
	if len(a) > 50 {
		t.FailNow()
	}
}

func BenchmarkFilter(b *testing.B) {
	var a = Range(0, 100, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(a, func(i int) bool {
			return i < 50
		})
	}
}
