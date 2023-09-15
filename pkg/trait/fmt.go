package trait

type Display[T any] interface {
	Display() string
}

type Debug[T any] interface {
	Debug() string
}
