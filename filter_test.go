package functools

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var a = Range(0, 100, 1)
	a = Filter(func(i int) bool {
		return i < 50
	}, a)
	if len(a) > 50 {
		t.FailNow()
	}
}

func BenchmarkFilter(b *testing.B) {
	var a = Range(0, 100, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(func(i int) bool {
			return i < 50
		}, a)
	}
}
