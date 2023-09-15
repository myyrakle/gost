package trait

type Clone[T any] interface {
	Clone() T
}
