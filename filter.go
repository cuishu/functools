package functools

type FilterFunc[A any] func(A) bool

func Filter[A any](list []A, f FilterFunc[A]) []A {
	var result []A = make([]A, 0, len(list))
	for _, data := range list {
		if f(data) {
			result = append(result, data)
		}
	}
	return result
}
