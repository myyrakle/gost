package vec

import "github.com/myyrakle/gost/pkg/option"

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

// Appends an element to the back of a collection.
func (self *Vec[T]) Push(value T) {
	self.data = append(self.data, value)
}

// Removes the last element from a vector and returns it, or None if it is empty.
func (self *Vec[T]) Pop() option.Option[T] {
	if len(self.data) == 0 {
		return option.None[T]()
	} else {
		value := self.data[len(self.data)-1]
		self.data = self.data[:len(self.data)-1]
		return option.Some[T](value)
	}
}

// Extracts a slice containing the entire vector.
func (self Vec[T]) AsSlice() []T {
	return self.data
}

// Inserts an element at position index within the vector, shifting all elements after it to the right.
func (self *Vec[T]) Insert(index int, value T) {
	self.data = append(self.data, value)
	copy(self.data[index+1:], self.data[index:])
	self.data[index] = value
}

// // Removes and returns the element at position index within the vector, shifting all elements after it to the left.
// func (self *Vec[T]) Remove(index int) T {
// 	value := self.data[index]
// 	copy(self.data[index:], self.data[index+1:])
// 	self.data[len(self.data)-1] = nil
// 	self.data = self.data[:len(self.data)-1]
// 	return value
// }
