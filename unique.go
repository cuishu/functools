package functools

func Unique[T Number | string](list []T) []T {
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
