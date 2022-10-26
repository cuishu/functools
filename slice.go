package functools

func Reverse[T any](list []T) []T {
	length := len(list)
	n := length - 1
	reverseList := make([]T, length)
	for i, item := range list {
		reverseList[n-i] = item
	}
	return reverseList
}
