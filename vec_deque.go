package gost

type VecDeque[T any] struct {
	buffer []T
	head   uint
	tail   uint
}
