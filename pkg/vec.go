package gost

import (
	"reflect"
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
func WithCapacity[T any](capacity Int) Vec[T] {
	return Vec[T]{data: make([]T, 0, capacity)}
}

// Returns the total number of elements the vector can hold without reallocating.
func (self Vec[T]) Capacity() Int {
	return Int(cap(self.data))
}

// Returns the number of elements in the vector, also referred to as its ‘length’.
func (self Vec[T]) Len() Int {
	return Int(len(self.data))
}

// Returns true if the vector contains no elements.
func (self Vec[T]) IsEmpty() Bool {
	return self.Len() == 0
}

// Reserves capacity for at least additional more elements to be inserted in the given Vec<T>. The collection may reserve more space to speculatively avoid frequent reallocations. After calling reserve, capacity will be greater than or equal to self.len() + additional. Does nothing if capacity is already sufficient.
func (self *Vec[T]) Reserve(additional Int) {
	if self.Capacity() < self.Len()+additional {
		self.data = append(self.data, make([]T, additional)...)
	}
}

// Appends an element to the back of a collection.
func (self *Vec[T]) Push(value T) {
	self.data = append(self.data, value)
}

// Removes the last element from a vector and returns it, or None if it is empty.
func (self *Vec[T]) Pop() Option[T] {
	if len(self.data) == 0 {
		return None[T]()
	} else {
		value := self.data[len(self.data)-1]
		self.data = self.data[:len(self.data)-1]
		return Some[T](value)
	}
}

// Moves all the elements of other Into self, leaving other empty.
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
func (self *Vec[T]) Insert(index Int, value T) {
	self.data = append(self.data, value)
	copy(self.data[index+1:], self.data[index:])
	self.data[index] = value
}

// Retains only the elements specified by the predicate.
func (self *Vec[T]) Retain(predicate func(T) Bool) {
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
	seen := make(map[any]bool)
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
func (self Vec[T]) Get(index Int) Option[T] {
	if index < 0 || index >= self.Len() {
		return None[T]()
	} else {
		return Some[T](self.data[index])
	}
}

// Returns a reference to an element or subslice, without doing bounds checking.
// For a safe alternative see get.
func (self Vec[T]) GetUnchecked(index Int) T {
	return self.data[index]
}

// Swaps two elements in the slice.
// If a equals to b, it’s guaranteed that elements won’t change value.
func (self *Vec[T]) Swap(a, b Int) {
	self.data[a], self.data[b] = self.data[b], self.data[a]
}

// Reverses the order of elements in the slice, in place.
func (self *Vec[T]) Reverse() {
	for i := 0; i < len(self.data)/2; i++ {
		self.Swap(Int(i), Int(len(self.data)-1-i))
	}
}

// Returns true if the slice contains an element with the given value.
// This operation is O(n).
// Note that if you have a sorted slice, binary_search may be faster.
func (self Vec[T]) Contains(value T) Bool {
	for _, v := range self.data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Binary searches this slice for a given element. If the slice is not sorted, the returned result is unspecified and meaningless.
// If the value is found then Result::Ok is returned, containing the index of the matching element. If there are multiple matches, then any one of the matches could be returned. The index is chosen deterministically, but is subject to change in future versions of Rust. If the value is not found then Result::Err is returned, containing the index where a matching element could be inserted while maIntaining sorted order.
// func (self Vec[T]) BinarySearch(value T) Option[Int] {
// 	low := 0
// 	high := len(self.data) - 1

// 	for low <= high {
// 		mid := (low + high) / 2
// 		if self.data[mid] < value {
// 			low = mid + 1
// 		} else if self.data[mid] > value {
// 			high = mid - 1
// 		} else {
// 			return Some[Int](mid)
// 		}
// 	}

// 	return None[Int]()
// }

// Fills self with elements by cloning value.
func (self *Vec[T]) Fill(value T) {
	for i := 0; i < len(self.data); i++ {
		self.data[i] = value
	}
}

// Fills self with elements returned by calling a closure repeatedly.
func (self *Vec[T]) FillWith(f func() T) {
	for i := 0; i < len(self.data); i++ {
		self.data[i] = f()
	}
}

// Sorts the slice.
// This sort is stable (i.e., does not reorder equal elements) and O(n * log(n)) worst-case.
// When applicable, unstable sorting is preferred because it is generally faster than stable sorting and it doesn’t allocate auxiliary memory. See sort_unstable.
// func (self *Vec[T]) Sort() {
// 	// type check
// 	sort.SliceStable(self.data, func(i, j Int) { return self.data[i] < self.data[j] })
// }
