package vec

import (
	"reflect"

	"github.com/myyrakle/gost/pkg/option"
)

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

// Returns the number of elements in the vector, also referred to as its ‘length’.
func (self Vec[T]) Len() int {
	return len(self.data)
}

// Returns true if the vector contains no elements.
func (self Vec[T]) IsEmpty() bool {
	return self.Len() == 0
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

// Moves all the elements of other into self, leaving other empty.
func (self *Vec[T]) Append(other *Vec[T]) {
	self.data = append(self.data, other.data...)
	other.data = make([]T, 0)
}

// Clears the vector, removing all values.
func (self *Vec[T]) Clear() {
	self.data = make([]T, 0)
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

// Retains only the elements specified by the predicate.
func (self *Vec[T]) Retain(predicate func(T) bool) {
	newData := make([]T, 0, len(self.data))
	for _, value := range self.data {
		if predicate(value) {
			newData = append(newData, value)
		}
	}
	self.data = newData
}

// Removes all but the first of consecutive elements in the vector that resolve to the same key.
// If the vector is sorted, this removes all duplicates.
func (self *Vec[T]) DedupByKey(key func(T) any) {
	newData := make([]T, 0, len(self.data))
	seen := make(map[interface{}]bool)
	for _, value := range self.data {
		k := key(value)
		if !seen[k] {
			seen[k] = true
			newData = append(newData, value)
		}
	}
	self.data = newData
}

// Removes consecutive repeated elements in the vector according to the PartialEq trait implementation.
// If the vector is sorted, this removes all duplicates.
func (self *Vec[T]) Dedup() {
	if len(self.data) <= 1 {
		return
	}

	prev := 0

	for i := 1; i < len(self.data); i++ {
		if !reflect.DeepEqual(self.data[prev], self.data[i]) {
			prev++
			self.data[prev] = self.data[i]
		}
	}

	self.data = self.data[:prev+1]
}

// Returns a reference to an element or subslice depending on the type of index.
// If given a position, returns a reference to the element at that position or None if out of bounds.
// If given a range, returns the subslice corresponding to that range, or None if out of bounds.
func (self Vec[T]) Get(index int) option.Option[T] {
	if index < 0 || index >= len(self.data) {
		return option.None[T]()
	} else {
		return option.Some[T](self.data[index])
	}
}
