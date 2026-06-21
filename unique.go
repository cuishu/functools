package functools

func Unique[T comparable](list []T) []T {
	var result []T = make([]T, 0, len(list))
	var set map[T]struct{} = make(map[T]struct{}, len(list))
	for _, data := range list {
		if _, ok := set[data]; !ok {
			result = append(result, data)
			set[data] = struct{}{}
		}
	}
	return result
}

func UniqueBy[A any, B comparable](list []A, selector func(A) B) []A {
	var result []A = make([]A, 0, len(list))
	var set map[B]struct{} = make(map[B]struct{}, len(list))
	for _, data := range list {
		key := selector(data)
		if _, ok := set[key]; !ok {
			result = append(result, data)
			set[key] = struct{}{}
		}
	}
	return result
}