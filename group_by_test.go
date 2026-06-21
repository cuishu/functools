package functools

import "testing"

func TestGroupBy(t *testing.T) {
	var list []int
	for i := 0; i < 100; i++ {
		list = append(list, i)
	}
	result := GroupBy(list, func(i int) int {
		return i % 2
	})
	t.Log(result)
}
