package functools

type MapFunc[A any, B any] func(A) B

func Map[A any, B any](f MapFunc[A, B], list []A) []B {
	var result []B = make([]B, len(list))
	for i, data := range list {
		result[i] = f(data)
	}
	return result
}
