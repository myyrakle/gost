package gost

type Shl[T any] interface {
	Shl(lhs T) T
}

type Shr[T any] interface {
	Shr(lhs T) T
}
