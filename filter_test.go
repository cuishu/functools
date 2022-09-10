package functools

import (
	"testing"
)

func TestFilter(t *testing.T) {
	var a []int
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}
	a = Filter(func(i int) bool {
		return i < 50
	}, a)
	if len(a) > 50 {
		t.FailNow()
	}
}
