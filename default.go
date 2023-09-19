package gost

// A trait for giving a type a useful default value.
type Default[T any] interface {
	Default() T
}
