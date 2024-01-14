package functools

type ReduceFunc[A any] func(A, A) A

func Reduce[A any](f ReduceFunc[A], list []A, init A) A {
	var result A = init
	for _, data := range list {
		result = f(result, data)
	}
	return result
}
