package functools

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Int | Uint | Float
}

func Sum[T Number](slice []T) T {
	var n T
	return Reduce(func(x, y T) T {
		return x + y
	}, slice, n)
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Range[T Int | Uint](start, stop, step T) []T {
	var result []T = make([]T, (stop-start)/step)
	n := 0
	for i := start; i < stop; i += step {
		result[n] = i
		n++
	}
	return result
}
