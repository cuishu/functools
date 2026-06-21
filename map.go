package functools

type MapFunc[A any, B any] func(A) B

func Map[A any, B any](list []A, f MapFunc[A, B]) []B {
	var result []B = make([]B, len(list))
	for i, data := range list {
		result[i] = f(data)
	}
	return result
}

func ForEach[A any](list []A, f func(A)) {
	for _, data := range list {
		f(data)
	}
}
