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

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Fail()
	}

	if Min(2, 1) != 1 {
		t.Fail()
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
