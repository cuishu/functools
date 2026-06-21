package functools

// Union 函数将多个列表中的元素去重，返回一个列表。
// 元素的顺序与第一个列表中的顺序相同。
func Union[T comparable](collections ...[]T) []T {
	set := make(map[T]struct{})
	result := make([]T, 0)
	for _, collection := range collections {
		for _, item := range collection {
			if _, ok := set[item]; !ok {
				result = append(result, item)
				set[item] = struct{}{}
			}
		}
	}
	return result
}

// UnionBy 函数将多个列表中的元素根据选择器函数的返回值进行去重，返回一个列表。
func UnionBy[A any, B comparable](selector func(A) B, collections ...[]A) []A {
	set := make(map[B]struct{})
	result := make([]A, 0)
	for _, collection := range collections {
		for _, item := range collection {
			key := selector(item)
			if _, ok := set[key]; !ok {
				result = append(result, item)
				set[key] = struct{}{}
			}
		}
	}
	return result
}
