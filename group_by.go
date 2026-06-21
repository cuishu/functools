package functools

// GroupBy 函数将列表中的元素根据选择器函数的返回值进行分组，返回一个映射，键为选择器函数的返回值，值为对应的元素列表。
// 选择器函数的返回值应为可比较的类型。
func GroupBy[A any, B comparable](list []A, selector func(A) B) map[B][]A {
	var result map[B][]A = make(map[B][]A)
	for _, data := range list {
		key := selector(data)
		result[key] = append(result[key], data)
	}
	return result
}
