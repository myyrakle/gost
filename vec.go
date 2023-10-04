package gost

import (
	"fmt"
	"reflect"
	"sort"
)

type Vec[T any] struct {
	data []T
}

// Constructs a new, empty Vec<T>.
// The vector will not allocate until elements are pushed onto it.
func VecNew[T any]() Vec[T] {
	return Vec[T]{data: make([]T, 0)}
}

// Constructs a new, empty Vec<T> with at least the specified capacity.
func VecWithCapacity[T any](capacity ISize) Vec[T] {
	return Vec[T]{data: make([]T, 0, capacity)}
}

// Constructs a new, empty Vec<T> with at least the specified capacity.
func VecWithLen[T any](len ISize) Vec[T] {
	return Vec[T]{data: make([]T, len)}
}

// Returns the total number of elements the vector can hold without reallocating.
func (self Vec[T]) Capacity() ISize {
	return ISize(cap(self.data))
}

// Returns the number of elements in the vector, also referred to as its ‘length’.
func (self Vec[T]) Len() ISize {
	return ISize(len(self.data))
}

// Returns true if the vector contains no elements.
func (self Vec[T]) IsEmpty() Bool {
	return self.Len() == 0
}

// Reserves capacity for at least additional more elements to be inserted in the given Vec<T>. The collection may reserve more space to speculatively avoid frequent reallocations. After calling reserve, capacity will be greater than or equal to self.len() + additional. Does nothing if capacity is already sufficient.
func (self *Vec[T]) Reserve(additional ISize) {
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

// Moves all the elements of other ISizeo self, leaving other empty.
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
func (self *Vec[T]) Insert(index ISize, value T) {
	self.data = append(self.data, value)
	copy(self.data[index+1:], self.data[index:])
	self.data[index] = value
}

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
func (self Vec[T]) Get(index ISize) Option[T] {
	if index < 0 || index >= self.Len() {
		return None[T]()
	} else {
		return Some[T](self.data[index])
	}
}

// Returns a reference to an element or subslice, without doing bounds checking.
// For a safe alternative see get.
func (self Vec[T]) GetUnchecked(index ISize) T {
	return self.data[index]
}

// Swaps two elements in the slice.
// If a equals to b, it’s guaranteed that elements won’t change value.
func (self *Vec[T]) Swap(a, b ISize) {
	self.data[a], self.data[b] = self.data[b], self.data[a]
}

// Reverses the order of elements in the slice, in place.
func (self *Vec[T]) Reverse() {
	for i := 0; i < len(self.data)/2; i++ {
		self.Swap(ISize(i), ISize(len(self.data)-1-i))
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
func (self Vec[T]) BinarySearch(value T) Option[ISize] {
	low := 0
	high := len(self.data) - 1

	for low <= high {
		mid := (low + high) / 2

		midData := castToOrd(self.data[mid]).Unwrap()
		ordering := midData.Cmp(value)

		switch ordering {
		case OrderingLess:
			{
				low = mid + 1
			}
		case OrderingGreater:
			{
				high = mid - 1
			}
		case OrderingEqual:
			{
				return Some[ISize](ISize(mid))
			}
		}
	}

	return None[ISize]()
}

// Binary searches this slice with a comparator function.
// The comparator function should return an order code that indicates whether its argument is Less, Equal or Greater the desired target. If the slice is not sorted or if the comparator function does not implement an order consistent with the sort order of the underlying slice, the returned result is unspecified and meaningless.
func (self Vec[T]) BinarySearchBy(f func(T) Ordering) Option[ISize] {
	low := 0
	high := len(self.data) - 1

	for low <= high {
		mid := (low + high) / 2
		ordering := f(self.data[mid])

		switch ordering {
		case OrderingLess:
			{
				low = mid + 1
			}
		case OrderingGreater:
			{
				high = mid - 1
			}
		case OrderingEqual:
			{
				return Some[ISize](ISize(mid))
			}
		}
	}

	return None[ISize]()
}

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
func (self *Vec[T]) Sort() {
	sort.SliceStable(self.data, func(i, j int) bool {
		lhs := castToOrd(self.data[i]).Unwrap()
		rhs := self.data[j]

		return lhs.Cmp(rhs) == OrderingLess
	})
}

// Sorts the slice with a comparator function.
// This sort is stable (i.e., does not reorder equal elements) and O(n * log(n)) worst-case.
// The comparator function must define a total ordering for the elements in the slice. If the ordering is not total, the order of the elements is unspecified. An order is a total order if it is (for all a, b and c):
// - total and antisymmetric: exactly one of a < b, a == b or a > b is true, and
// - transitive, a < b and b < c implies a < c. The same must hold for both == and >.
func (self *Vec[T]) SortBy(compare func(T, T) Ordering) {
	sort.SliceStable(self.data, func(i, j int) bool {
		lhs := self.data[i]
		rhs := self.data[j]

		return compare(lhs, rhs) == OrderingLess
	})
}

// Sorts the slice, but might not preserve the order of equal elements.
// This sort is unstable (i.e., may reorder equal elements), in-place (i.e., does not allocate), and O(n * log(n)) worst-case.
func (self *Vec[T]) SortUnstable() {
	sort.Slice(self.data, func(i, j int) bool {
		lhs := castToOrd(self.data[i]).Unwrap()
		rhs := self.data[j]

		return lhs.Cmp(rhs) == OrderingLess
	})
}

// Sorts the slice with a comparator function, but might not preserve the order of equal elements.
// This sort is unstable (i.e., may reorder equal elements), in-place (i.e., does not allocate), and O(n * log(n)) worst-case.
// The comparator function must define a total ordering for the elements in the slice. If the ordering is not total, the order of the elements is unspecified. An order is a total order if it is (for all a, b and c):
// - total and antisymmetric: exactly one of a < b, a == b or a > b is true, and
// - transitive, a < b and b < c implies a < c. The same must hold for both == and >.
func (self *Vec[T]) SortUnstableBy(compare func(T, T) Ordering) {
	sort.Slice(self.data, func(i, j int) bool {
		lhs := self.data[i]
		rhs := self.data[j]

		return compare(lhs, rhs) == OrderingLess
	})
}

type VecIter[T any] struct {
	vec      Vec[T]
	position ISize
}

// into_iter
func (self Vec[T]) IntoIter() Iterator[T] {
	return &VecIter[T]{vec: self, position: 0}
}

// next
func (self *VecIter[T]) Next() Option[T] {
	if self.position >= self.vec.Len() {
		return None[T]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[T](value)
}

// map
func (self VecIter[T]) Map(f func(T) T) Iterator[T] {
	newVec := VecNew[T]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.Push(f(value.Unwrap()))
	}
}

// filter
func (self VecIter[T]) Filter(f func(T) Bool) Iterator[T] {
	newVec := VecNew[T]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}

		unwraped := value.Unwrap()
		if f(unwraped) {
			newVec.Push(unwraped)
		}
	}
}

// fold
func (self VecIter[T]) Fold(init T, f func(T, T) T) T {
	for {
		value := self.Next()

		if value.IsNone() {
			return init
		}

		init = f(init, value.Unwrap())
	}
}

// rev
func (self VecIter[T]) Rev() Iterator[T] {
	newVec := VecWithLen[T](self.vec.Len())
	i := self.vec.Len() - 1

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.AsSlice()[i] = value.Unwrap()
		i--
	}
}

func (self VecIter[T]) CollectToVec() Vec[T] {
	vec := Vec[T]{}
	for {
		value := self.Next()
		if value.IsNone() {
			return vec
		}
		vec.Push(value.Unwrap())
	}
}

func (self VecIter[T]) CollectToLinkedList() LinkedList[T] {
	list := LinkedListNew[T]()
	for {
		value := self.Next()
		if value.IsNone() {
			return list
		}
		list.PushBack(value.Unwrap())
	}
}

// impl Display for Vec
func (self Vec[T]) Display() String {
	buffer := String("")
	buffer += "Vec["

	for i := 0; i < len(self.data); i++ {
		e := self.data[i]

		display := castToDisplay(e)
		if display.IsSome() {
			buffer += display.Unwrap().Display()
		} else {
			typeName := getTypeName(e)

			panic(fmt.Sprintf("'%s' does not implement Display[%s]", typeName, typeName))
		}

		if i != len(self.data)-1 {
			buffer += ", "
		}
	}

	buffer += "]"

	return String(buffer)
}

// impl Debug for Vec
func (self Vec[T]) Debug() String {
	return self.Display()
}
