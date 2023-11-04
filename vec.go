package gost

import (
	"fmt"
	"reflect"
	"sort"
)

// A contiguous growable array type with heap-allocated contents, written Vec<T>.
// Vectors have O(1) indexing, amortized O(1) push (to the end) and O(1) pop (from the end).
type Vec[T any] struct {
	data []T
}

// Constructs a new, empty Vec<T>.
// The vector will not allocate until elements are pushed onto it.
func VecNew[T any]() Vec[T] {
	return Vec[T]{data: make([]T, 0)}
}

// Constructs a new, empty Vec<T> with at least the specified capacity.
func VecWithCapacity[T any](capacity USize) Vec[T] {
	return Vec[T]{data: make([]T, 0, capacity)}
}

// Constructs a new, empty Vec<T> with at least the specified len.
func VecWithLen[T any](len USize) Vec[T] {
	return Vec[T]{data: make([]T, len)}
}

// Returns the total number of elements the vector can hold without reallocating.
//
//	vec := gost.VecNew[gost.I32]()
//	gost.AssertEq(vec.Capacity(), gost.USize(0))
//
//	vec.Reserve(10)
//	gost.AssertEq(vec.Capacity(), gost.USize(10))
func (self Vec[T]) Capacity() USize {
	return USize(cap(self.data))
}

// Returns the number of elements in the vector, also referred to as its ‘length’.
//
//	vec := gost.VecNew[gost.I32]()
//	gost.AssertEq(vec.Len(), gost.USize(0))
//
//	vec.Push(1)
//	gost.AssertEq(vec.Len(), gost.USize(1))
func (self Vec[T]) Len() USize {
	return USize(len(self.data))
}

// Returns true if the vector contains no elements.
//
//	vec := gost.VecNew[gost.I32]()
//	gost.Assert(vec.IsEmpty())
//
//	vec.Push(1)
//	gost.Assert(!vec.IsEmpty())
func (self Vec[T]) IsEmpty() Bool {
	return self.Len() == 0
}

// Reserves capacity for at least additional more elements to be inserted in the given Vec<T>. The collection may reserve more space to speculatively avoid frequent reallocations. After calling reserve, capacity will be greater than or equal to self.len() + additional. Does nothing if capacity is already sufficient.
//
//	vec := gost.VecNew[gost.I32]()
//	gost.AssertEq(vec.Capacity(), gost.USize(0))
//
//	vec.Reserve(10)
//	gost.AssertEq(vec.Capacity(), gost.USize(10))
func (self *Vec[T]) Reserve(additional USize) {
	if self.Capacity() < self.Len()+additional {
		self.data = append(self.data, make([]T, additional)...)
	}
}

// Appends an element to the back of a collection.
//
//	vec := gost.VecNew[gost.I32]()
//	gost.AssertEq(vec.Len(), gost.USize(0))
//
//	vec.Push(1)
//	gost.AssertEq(vec.Len(), gost.USize(1))
func (self *Vec[T]) Push(value T) {
	self.data = append(self.data, value)
}

// Removes the last element from a vector and returns it, or None if it is empty.
//
//	vec := gost.VecNew[gost.I32]()
//	gost.AssertEq(vec.Pop().IsNone(), true)
//
//	vec.Push(1)
//	gost.AssertEq(vec.Pop().IsSome(), true)
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
//
//	vec1 := gost.VecNew[gost.I32]()
//	vec2 := gost.VecNew[gost.I32]()
//	vec1.Push(1)
//	vec2.Push(2)
//	vec1.Append(&vec2)
//	gost.AssertEq(vec1.Len(), gost.USize(2))
func (self *Vec[T]) Append(other *Vec[T]) {
	self.data = append(self.data, other.data...)
	other.data = make([]T, 0)
}

// Clears the vector, removing all values.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Clear()
//	gost.AssertEq(vec.Len(), gost.USize(0))
func (self *Vec[T]) Clear() {
	self.data = make([]T, 0)
}

// Extracts a slice containing the entire vector.
func (self Vec[T]) AsSlice() []T {
	return self.data
}

// Inserts an element at position index within the vector, shifting all elements after it to the right.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Insert(1, 3)
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
//	gost.AssertEq(vec.GetUnchecked(1), gost.I32(3))
func (self *Vec[T]) Insert(index USize, value T) {
	self.data = append(self.data, value)
	copy(self.data[index+1:], self.data[index:])
	self.data[index] = value
}

// Retains only the elements specified by the predicate.
// In other words, remove all elements e such that f(e) returns false.
// This method operates in place and preserves the order of the retained elements.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	vec.Retain(func(e gost.I32) gost.Bool {
//		return e == 2
//	})
//	gost.AssertEq(vec.Len(), gost.USize(1))
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
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(2)
//	vec.Push(3)
//	vec.DedupByKey(func(e gost.I32) gost.I32 {
//		return e
//	}
//	gost.AssertEq(vec.Len(), gost.USize(3))
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
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(2)
//	vec.Push(3)
//	vec.Dedup()
//	gost.AssertEq(vec.Len(), gost.USize(3))
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
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	gost.AssertEq(vec.Get(0).Unwrap(), gost.I32(1))
//	gost.AssertEq(vec.Get(1).Unwrap(), gost.I32(2))
func (self Vec[T]) Get(index USize) Option[T] {
	if index < 0 || index >= self.Len() {
		return None[T]()
	} else {
		return Some[T](self.data[index])
	}
}

// Returns a reference to an element or subslice, without doing bounds checking.
// For a safe alternative see get.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
func (self Vec[T]) GetUnchecked(index USize) T {
	return self.data[index]
}

// Set value
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	vec.Set(0, 4)
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(4))
func (self *Vec[T]) Set(index USize, value T) Option[T] {
	if index < 0 || index >= self.Len() {
		return Some[T](self.data[index])
	} else {
		self.data[index] = value
		return None[T]()
	}
}

// Set unchecked
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	vec.SetUnchecked(0, 4)
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(4))
func (self *Vec[T]) SetUnchecked(index USize, value T) {
	self.data[index] = value
}

// Swaps two elements in the slice.
// If a equals to b, it’s guaranteed that elements won’t change value.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	vec.Swap(0, 2)
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(3))
//	gost.AssertEq(vec.GetUnchecked(2), gost.I32(1))
func (self *Vec[T]) Swap(a, b USize) {
	self.data[a], self.data[b] = self.data[b], self.data[a]
}

// Reverses the order of elements in the slice, in place.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	vec.Reverse()
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(3))
func (self *Vec[T]) Reverse() {
	for i := 0; i < len(self.data)/2; i++ {
		self.Swap(USize(i), USize(len(self.data)-1-i))
	}
}

// Returns true if the slice contains an element with the given value.
// This operation is O(n).
// Note that if you have a sorted slice, binary_search may be faster.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	gost.Assert(vec.Contains(2))
//	gost.Assert(!vec.Contains(4))
func (self Vec[T]) Contains(value T) Bool {
	for _, v := range self.data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Binary searches this slice for a given element. If the slice is not sorted, the returned result is unspecified and meaningless.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	gost.Assert(vec.BinarySearch(2).IsSome())
//	gost.Assert(vec.BinarySearch(4).IsNone())
func (self Vec[T]) BinarySearch(value T) Option[USize] {
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
				return Some[USize](USize(mid))
			}
		}
	}

	return None[USize]()
}

// Binary searches this slice with a comparator function.
// The comparator function should return an order code that indicates whether its argument is Less, Equal or Greater the desired target. If the slice is not sorted or if the comparator function does not implement an order consistent with the sort order of the underlying slice, the returned result is unspecified and meaningless.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(1)
//	vec.Push(2)
//	vec.Push(3)
//	gost.Assert(vec.BinarySearchBy(func(e gost.I32) gost.Ordering {
//		if e < 2 {
//			return gost.OrderingLess
//		} else if e > 2 {
//			return gost.OrderingGreater
//		} else {
//			return gost.OrderingEqual
//		}
//	}).IsSome())
func (self Vec[T]) BinarySearchBy(f func(T) Ordering) Option[USize] {
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
				return Some[USize](USize(mid))
			}
		}
	}

	return None[USize]()
}

// Fills self with elements by cloning value.
//
//	vec := gost.VecWithLen[gost.I32](gost.USize(5))
//	vec.Fill(gost.I32(1))
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
func (self *Vec[T]) Fill(value T) {
	for i := 0; i < len(self.data); i++ {
		self.data[i] = value
	}
}

// Fills self with elements returned by calling a closure repeatedly.
//
//	vec := gost.VecWithLen[gost.I32](gost.USize(5))
//	vec.FillWith(func() gost.I32 {
//		return 1
//	})
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
func (self *Vec[T]) FillWith(f func() T) {
	for i := 0; i < len(self.data); i++ {
		self.data[i] = f()
	}
}

// Sorts the slice.
// This sort is stable (i.e., does not reorder equal elements) and O(n * log(n)) worst-case.
// When applicable, unstable sorting is preferred because it is generally faster than stable sorting and it doesn’t allocate auxiliary memory. See sort_unstable.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(3)
//	vec.Push(2)
//	vec.Push(1)
//	vec.Sort()
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
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
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(3)
//	vec.Push(2)
//	vec.Push(1)
//	vec.SortBy(func(lhs, rhs gost.I32) gost.Ordering {
//		if lhs < rhs {
//			return gost.OrderingLess
//		} else if lhs > rhs {
//			return gost.OrderingGreater
//		} else {
//			return gost.OrderingEqual
//		}
//	})
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
func (self *Vec[T]) SortBy(compare func(T, T) Ordering) {
	sort.SliceStable(self.data, func(i, j int) bool {
		lhs := self.data[i]
		rhs := self.data[j]

		return compare(lhs, rhs) == OrderingLess
	})
}

// Sorts the slice, but might not preserve the order of equal elements.
// This sort is unstable (i.e., may reorder equal elements), in-place (i.e., does not allocate), and O(n * log(n)) worst-case.
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(3)
//	vec.Push(2)
//	vec.Push(1)
//	vec.SortUnstable()
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
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
//
//	vec := gost.VecNew[gost.I32]()
//	vec.Push(3)
//	vec.Push(2)
//	vec.Push(1)
//	vec.SortUnstableBy(func(lhs, rhs gost.I32) gost.Ordering {
//		if lhs < rhs {
//			return gost.OrderingLess
//		} else if lhs > rhs {
//			return gost.OrderingGreater
//		} else {
//			return gost.OrderingEqual
//		}
//	})
//	gost.AssertEq(vec.GetUnchecked(0), gost.I32(1))
func (self *Vec[T]) SortUnstableBy(compare func(T, T) Ordering) {
	sort.Slice(self.data, func(i, j int) bool {
		lhs := self.data[i]
		rhs := self.data[j]

		return compare(lhs, rhs) == OrderingLess
	})
}

// iterator for Vec
type VecIter[T any] struct {
	vec      Vec[T]
	position USize
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

// collect to Vec
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

// collect to LinkedList
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

// impl AsRef for Vec
func (self Vec[T]) AsRef() *Vec[T] {
	return &self
}

// impl Clone for Vec
func (self Vec[T]) Clone() Vec[T] {
	cloned := VecWithLen[T](self.Len())

	for i := USize(0); i < self.Len(); i++ {
		e := castToClone(self.data[i]).Unwrap()
		cloned.SetUnchecked(USize(i), e.Clone())
	}

	return cloned
}

// impl Eq for Vec
func (self Vec[T]) Eq(other Vec[T]) Bool {
	if self.Len() != other.Len() {
		return false
	}

	for i := USize(0); i < self.Len(); i++ {
		if !castToEq(self.data[i]).Unwrap().Eq(other.data[i]) {
			return false
		}
	}

	return true
}
