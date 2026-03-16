package functools

func Zip[A any, B any](a []A, b []B) []Pair[A, B] {
	n := Min(len(a), len(b))

	var result []Pair[A, B] = make([]Pair[A, B], n)
	for i := 0; i < n; i++ {
		result[i] = NewPair(a[i], b[i])
	}
	return result
}
