package functools

func Zip[A any, B any](a []A, b []B) []Tuple[A, B] {
	n := Min(len(a), len(b))

	var result []Tuple[A, B] = make([]Tuple[A, B], n)
	for i := 0; i < n; i++ {
		result[i] = NewTuple(a[i], b[i])
	}
	return result
}
