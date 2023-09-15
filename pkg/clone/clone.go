package clone

type Clone[T any] interface {
	Clone() T
}

func DeepClone[T any](value T) T {
	// TODO: Implement deep clone
	return value
}
