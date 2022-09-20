package functools

type Tuple[A any, B any] struct {
	V1 A
	V2 B
}

func NewTuple[A any, B any](a A, b B) Tuple[A, B] {
	return Tuple[A, B]{
		V1: a,
		V2: b,
	}
}

func (tuple Tuple[A, B]) Values() (A, B) {
	return tuple.V1, tuple.V2
}
