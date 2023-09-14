package vec

type Vec[T any] struct {
	data []T
}

// Constructs a new, empty Vec<T>.
// The vector will not allocate until elements are pushed onto it.
func New[T any]() Vec[T] {
	return Vec[T]{data: make([]T, 0)}
}

// Constructs a new, empty Vec<T> with at least the specified capacity.
func WithCapacity[T any](capacity int) Vec[T] {
	return Vec[T]{data: make([]T, 0, capacity)}
}
