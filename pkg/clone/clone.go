package clone

type Clone[T any] interface {
	Clone() T
}
