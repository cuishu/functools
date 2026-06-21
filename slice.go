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
