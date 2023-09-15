package trait

// A trait for converting a value to a String.
type ToString[T any] interface {
	ToString() string
}
