package functools

type Pair[A any, B any] struct {
	First A
	Second B
}

func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{
		First: a,
		Second: b,
	}
}

func (pair Pair[A, B]) Values() (A, B) {
	return pair.First, pair.Second
}
