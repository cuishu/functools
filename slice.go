package functools

// Reverse 函数反转切片，返回一个新的切片。
func Reverse[T any](list []T) []T {
	length := len(list)
	n := length - 1
	reverseList := make([]T, length)
	for i, item := range list {
		reverseList[n-i] = item
	}
	return reverseList
}

// ReverseInPlace 函数原地反转切片，无返回值。
func ReverseInPlace[T any](list []T) {
	length := len(list)
	n := length - 1
	for i := 0; i < length/2; i++ {
		list[i], list[n-i] = list[n-i], list[i]
	}
}

// Diff 返回两个切片中不共有的元素（对称差集），结果去重。
func Diff[T comparable](a, b []T) []T {
	setA := make(map[T]struct{})
	setB := make(map[T]struct{})

	for _, v := range a {
		setA[v] = struct{}{}
	}
	for _, v := range b {
		setB[v] = struct{}{}
	}

	var diff []T
	for v := range setA {
		if _, ok := setB[v]; !ok {
			diff = append(diff, v)
		}
	}
	for v := range setB {
		if _, ok := setA[v]; !ok {
			diff = append(diff, v)
		}
	}
	return diff
}

// DiffBy 返回两个切片中不共有的元素（对称差集），基于 selector 提取的键去重。
// 结果中的元素取自首次出现的位置。
func DiffBy[T any, K comparable](a, b []T, selector func(T) K) []T {
	mapA := make(map[K]T)
	mapB := make(map[K]T)

	for _, v := range a {
		key := selector(v)
		if _, ok := mapA[key]; !ok {
			mapA[key] = v
		}
	}
	for _, v := range b {
		key := selector(v)
		if _, ok := mapB[key]; !ok {
			mapB[key] = v
		}
	}

	var diff []T
	for key, val := range mapA {
		if _, ok := mapB[key]; !ok {
			diff = append(diff, val)
		}
	}
	for key, val := range mapB {
		if _, ok := mapA[key]; !ok {
			diff = append(diff, val)
		}
	}
	return diff
}

// DiffSeparate 返回 (只在a中, 只在b中) 的两个切片，均去重。
func DiffSeparate[T comparable](a, b []T) (onlyA, onlyB []T) {
	setA := make(map[T]struct{})
	setB := make(map[T]struct{})

	for _, v := range a {
		setA[v] = struct{}{}
	}
	for _, v := range b {
		setB[v] = struct{}{}
	}

	for v := range setA {
		if _, ok := setB[v]; !ok {
			onlyA = append(onlyA, v)
		}
	}
	for v := range setB {
		if _, ok := setA[v]; !ok {
			onlyB = append(onlyB, v)
		}
	}
	return onlyA, onlyB
}

// DiffSeparateBy 返回两个切片：(只在a中, 只在b中)，基于 selector 提取的键去重。
func DiffSeparateBy[T any, K comparable](a, b []T, selector func(T) K) (onlyA, onlyB []T) {
	mapA := make(map[K]T)
	mapB := make(map[K]T)

	for _, v := range a {
		key := selector(v)
		if _, ok := mapA[key]; !ok {
			mapA[key] = v
		}
	}
	for _, v := range b {
		key := selector(v)
		if _, ok := mapB[key]; !ok {
			mapB[key] = v
		}
	}

	for key, val := range mapA {
		if _, ok := mapB[key]; !ok {
			onlyA = append(onlyA, val)
		}
	}
	for key, val := range mapB {
		if _, ok := mapA[key]; !ok {
			onlyB = append(onlyB, val)
		}
	}
	return onlyA, onlyB
}
