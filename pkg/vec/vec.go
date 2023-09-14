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

// Returns the total number of elements the vector can hold without reallocating.
func (self Vec[T]) Capacity() int {
	return cap(self.data)
}

// Reserves capacity for at least additional more elements to be inserted in the given Vec<T>. The collection may reserve more space to speculatively avoid frequent reallocations. After calling reserve, capacity will be greater than or equal to self.len() + additional. Does nothing if capacity is already sufficient.
func (self *Vec[T]) Reserve(additional int) {
	if self.Capacity() < len(self.data)+additional {
		self.data = append(self.data, make([]T, additional)...)
	}
}

// Extracts a slice containing the entire vector.
func (self Vec[T]) AsSlice() []T {
	return self.data
}
