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
	if len(a) == 0 {
		return Unique(b)
	}
	if len(b) == 0 {
		return Unique(a)
	}

	// 预分配一半容量（避免浪费，最坏扩容一次）
	cap := (len(a) + len(b)) / 2
	if cap < 1 {
		cap = 1
	}
	diff := make([]T, 0, cap)

	// 构建 b 的 map（键 -> 元素，去重）
	bMap := make(map[T]T, len(b))
	for _, v := range b {
		if _, ok := bMap[v]; !ok {
			bMap[v] = v
		}
	}

	// 去重 a 中已处理的键
	seenA := make(map[T]struct{}, len(a))

	for _, v := range a {
		if _, ok := seenA[v]; ok {
			continue
		}
		seenA[v] = struct{}{}

		if _, ok := bMap[v]; ok {
			// 该键在两边都存在，不属于差集，从 bMap 中移除
			delete(bMap, v)
		} else {
			// 仅在 a 中，加入结果
			diff = append(diff, v)
		}
	}

	// bMap 中剩余的元素仅在 b 中，追加到结果
	for _, v := range bMap {
		diff = append(diff, v)
	}

	return diff
}

// DiffBy 返回两个切片的对称差集（只存在于其中一个切片中的元素），
// 基于 selector 提取的键进行去重。结果不保证顺序。
func DiffBy[T any, K comparable](a, b []T, selector func(T) K) []T {
	if len(a) == 0 {
		return UniqueBy(b, selector)
	}
	if len(b) == 0 {
		return UniqueBy(a, selector)
	}

	// 预分配一半容量（避免浪费，最坏扩容一次）
	cap := (len(a) + len(b)) / 2
	if cap < 1 {
		cap = 1
	}
	diff := make([]T, 0, cap)

	// 构建 b 的 map（键 -> 元素，去重）
	bMap := make(map[K]T, len(b))
	for _, v := range b {
		key := selector(v)
		if _, ok := bMap[key]; !ok {
			bMap[key] = v
		}
	}

	// 去重 a 中已处理的键
	seenA := make(map[K]struct{}, len(a))

	for _, v := range a {
		key := selector(v)
		if _, ok := seenA[key]; ok {
			continue
		}
		seenA[key] = struct{}{}

		if _, ok := bMap[key]; ok {
			// 该键在两边都存在，不属于差集，从 bMap 中移除
			delete(bMap, key)
		} else {
			// 仅在 a 中，加入结果
			diff = append(diff, v)
		}
	}

	// bMap 中剩余的元素仅在 b 中，追加到结果
	for _, v := range bMap {
		diff = append(diff, v)
	}

	return diff
}

// calcHalfCap 计算半容量，至少为1
func calcHalfCap(length int) int {
	cap := length / 2
	if cap < 1 {
		cap = 1
	}
	return cap
}

// DiffSeparate 返回 (只在a中, 只在b中) 的两个切片，均去重。
func DiffSeparate[T comparable](a, b []T) (onlyA, onlyB []T) {
	// 快速路径：空切片
	if len(a) == 0 {
		return nil, Unique(b)
	}
	if len(b) == 0 {
		return Unique(a), nil
	}

	// 预分配容量
	onlyA = make([]T, 0, calcHalfCap(len(a)))
	onlyB = make([]T, 0, calcHalfCap(len(b)))

	// 构建 b 的 map（去重）
	bMap := make(map[T]struct{}, len(b))
	for _, v := range b {
		bMap[v] = struct{}{}
	}

	// 用于去重 a 中已处理的元素
	seenA := make(map[T]struct{}, len(a))

	for _, v := range a {
		if _, ok := seenA[v]; ok {
			continue
		}
		seenA[v] = struct{}{}

		if _, ok := bMap[v]; ok {
			delete(bMap, v) // 该键在两边都存在，从 bMap 中移除
		} else {
			onlyA = append(onlyA, v)
		}
	}

	// bMap 中剩余的元素就是仅存在于 b 的
	for v := range bMap {
		onlyB = append(onlyB, v)
	}

	return onlyA, onlyB
}

// DiffSeparateBy 返回两个切片：(只在a中, 只在b中)，基于 selector 提取的键去重。
func DiffSeparateBy[T any, K comparable](a, b []T, selector func(T) K) (onlyA, onlyB []T) {
	// 快速路径：空切片
	if len(a) == 0 {
		return nil, UniqueBy(b, selector)
	}
	if len(b) == 0 {
		return UniqueBy(a, selector), nil
	}

	// 预分配容量
	onlyA = make([]T, 0, calcHalfCap(len(a)))
	onlyB = make([]T, 0, calcHalfCap(len(b)))

	// 构建 b 的映射：键 -> 元素（遇到重复键保留首次出现）
	bMap := make(map[K]T, len(b))
	for _, v := range b {
		key := selector(v)
		if _, ok := bMap[key]; !ok {
			bMap[key] = v
		}
	}

	// 用于去重 a 中已处理过的键
	seenA := make(map[K]struct{}, len(a))

	for _, v := range a {
		key := selector(v)
		if _, ok := seenA[key]; ok {
			continue
		}
		seenA[key] = struct{}{}

		if _, ok := bMap[key]; ok {
			delete(bMap, key) // 键在两边都存在，移除
		} else {
			onlyA = append(onlyA, v)
		}
	}

	// bMap 中剩余的键值对即为仅存在于 b 的元素
	for _, v := range bMap {
		onlyB = append(onlyB, v)
	}

	return onlyA, onlyB
}
