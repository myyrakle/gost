package gost

import (
	"fmt"
	"sort"
)

const (
	_VECDEQUE_INITIAL_CAPACITY USize = 7 // 2^3 - 1
	_VECDEQUE_MINIMUM_CAPACITY USize = 1 // 2 - 1
)

// A double-ended queue implemented with a growable ring buffer.
type VecDeque[T any] struct {
	buffer []T
	len    USize
	head   USize
}

// Creates an empty VecDeque.
func VecDequeNew[T any]() VecDeque[T] {
	return VecDequeWithCapacity[T](_VECDEQUE_INITIAL_CAPACITY)
}

// Creates an empty VecDeque with at least the specified capacity.
func VecDequeWithCapacity[T any](capacity USize) VecDeque[T] {
	if capacity < _VECDEQUE_MINIMUM_CAPACITY {
		capacity = _VECDEQUE_MINIMUM_CAPACITY
	}

	return VecDeque[T]{buffer: make([]T, capacity), len: 0, head: 0}
}

// Constructs a new, empty Vec<T> with at least the specified capacity.
func VecDequeWithLen[T any](len USize) VecDeque[T] {
	deque := VecDequeWithCapacity[T](len)
	deque.len = len

	return deque
}

// Returns the number of elements in the vecdeque, also referred to as its ‘length’.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.Len(), gost.USize(2))
func (self VecDeque[T]) Len() USize {
	return USize(self.len)
}

// Returns true if the deque is empty.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.IsEmpty(), gost.Bool(false))
func (self VecDeque[T]) IsEmpty() Bool {
	return self.len == 0
}

// Returns the number of elements the deque can hold without reallocating.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	deque.Reserve(gost.USize(8))
//	gost.AssertEqual(deque.Capacity(), gost.USize(8))
//	gost.AssertEqual(deque.Len(), gost.USize(2))
func (self VecDeque[T]) Capacity() USize {
	return USize(len(self.buffer))
}

// Reserves capacity for at least additional more elements to be inserted in the given deque.
// The collection may reserve more space to speculatively avoid frequent reallocations.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	deque.Reserve(gost.USize(8))
//	gost.AssertEqual(deque.Capacity(), gost.USize(8))
func (self *VecDeque[T]) Reserve(additional USize) {
	// TODO: overflow check
	oldCapacity := self.Capacity()
	newCapacity := oldCapacity + additional

	if newCapacity > oldCapacity {
		newBuffer := make([]T, newCapacity)
		copy(newBuffer, self.buffer)
		self.buffer = newBuffer

		self._HandleCapacityIncrease(oldCapacity)
	}
}

// Prepends an element to the deque.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushFront(gost.I32(3))
//	deque.PushFront(gost.I32(4))
//	gost.AssertEqual(deque.Len(), gost.USize(2))
func (self *VecDeque[T]) PushFront(value T) {
	if self._IsFull() {
		self._Grow()
	}

	self.head = self._WrapSub(self.head, 1)
	self.buffer[self.head] = value
	self.len++
}

// Appends an element to the back of the deque.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.Len(), gost.USize(2))
func (self *VecDeque[T]) PushBack(value T) {
	if self._IsFull() {
		self._Grow()
	}

	self.buffer[self._ToPhysicalIndex(self.len)] = value
	self.len++
}

// Removes the first element and returns it, or `None` if the deque is empty.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.PopFront(), gost.Some[gost.I32](gost.I32(3)))
//	gost.AssertEqual(deque.PopFront(), gost.Some[gost.I32](gost.I32(4)))
//	gost.AssertEqual(deque.PopFront(), gost.None[gost.I32]())
func (self *VecDeque[T]) PopFront() Option[T] {
	if self.IsEmpty() {
		return None[T]()
	}

	oldHead := self.head
	self.head = self._ToPhysicalIndex(1)
	self.len--

	value := self.buffer[oldHead]

	return Some[T](value)
}

// Removes the last element from the deque and returns it, or `None` if it is empty.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.PopBack(), gost.Some[gost.I32](gost.I32(4)))
//	gost.AssertEqual(deque.PopBack(), gost.Some[gost.I32](gost.I32(3)))
//	gost.AssertEqual(deque.PopBack(), gost.None[gost.I32]())
func (self *VecDeque[T]) PopBack() Option[T] {
	if self.IsEmpty() {
		return None[T]()
	}

	self.len--
	return Some[T](self.buffer[self._ToPhysicalIndex(self.len)])
}

// Moves all the elements of other ISizeo self, leaving other empty.
//
//		deque1 := gost.VecDequeNew[gost.I32]()
//		deque2 := gost.VecDequeNew[gost.I32]()
//		deque1.PushBack(1)
//		deque2.PushBack(2)
//		deque1.Append(&deque2)
//		gost.AssertEq(deque1.Len(), gost.USize(2))
//	 gost.AssertEq(deque2.Len(), gost.USize(0))
func (self *VecDeque[T]) Append(other *VecDeque[T]) {
	self.Reserve(other.Len())

	for i := USize(0); i < other.Len(); i++ {
		self.PushBack(other.GetUnchecked(i))
	}

	other.Clear()
}

// Retains only the elements specified by the predicate.
// In other words, remove all elements e such that f(e) returns false.
// This method operates in place and preserves the order of the retained elements.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(1)
//	deque.Push(2)
//	deque.Push(3)
//	deque.Retain(func(e gost.I32) gost.Bool {
//		return e == 2
//	})
func (self *VecDeque[T]) Retain(f func(T) Bool) {
	newLen := USize(0)

	for i := USize(0); i < self.Len(); i++ {
		value := self.GetUnchecked(i)
		if f(value) {
			self.SetUnchecked(newLen, value)
			newLen++
		}
	}

	self.len = newLen
}

// Removes consecutive repeated elements in the vector according to the PartialEq trait implementation.
// If the vecdeque is sorted, this removes all duplicates.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(1)
//	deque.Push(2)
//	deque.Push(2)
//	deque.Push(3)
//	deque.Dedup()
//	gost.AssertEq(deque.Len(), gost.USize(3))
func (self *VecDeque[T]) Dedup(key func(T) any) {
	if self.IsEmpty() {
		return
	}

	newLen := USize(1)

	for i := USize(1); i < self.Len(); i++ {
		if !castToEq(self.GetUnchecked(i)).Unwrap().Eq(self.GetUnchecked(newLen - 1)) {
			self.SetUnchecked(newLen, self.GetUnchecked(i))
			newLen++
		}
	}

	self.len = newLen
}

// Swaps two elements in the slice.
// if a equals to b, it’s guaranteed that elements won’t change value.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(1)
//	deque.Push(2)
//	deque.Push(3)
//	deque.Swap(0, 2)
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(3))
//	gost.AssertEq(deque.GetUnchecked(2), gost.I32(1))
func (self *VecDeque[T]) Swap(a USize, b USize) {
	if a == b {
		return
	}

	a = self._ToPhysicalIndex(a)
	b = self._ToPhysicalIndex(b)

	self.buffer[a], self.buffer[b] = self.buffer[b], self.buffer[a]
}

// Provides a reference to the element at the given index.
// Element at index 0 is the front of the queue.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.Get(gost.USize(0)), gost.Some[gost.I32](gost.I32(3)))
func (self VecDeque[T]) Get(index USize) Option[T] {
	if index >= self.Len() {
		return None[T]()
	}

	return Some[T](self.buffer[uint(self._ToPhysicalIndex(index))])
}

// Returns a reference to an element or subslice, without doing bounds checking.
// For a safe alternative see get.
//
//		deque := gost.VecDequeNew[gost.I32]()
//		deque.PushBack(gost.I32(3))
//		deque.PushBack(gost.I32(4))
//	 gost.AssertEqual(deque.GetUnchecked(gost.USize(0)), gost.I32(3))
func (self VecDeque[T]) GetUnchecked(index USize) T {
	return self.buffer[uint(self._ToPhysicalIndex(index))]
}

// Set value at index.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.Set(gost.USize(0), gost.I32(4))
//	gost.AssertEqual(deque.Get(gost.USize(0)), gost.Some[gost.I32](gost.I32(4)))
func (self *VecDeque[T]) Set(index USize, value T) Option[T] {
	if index >= self.Len() {
		return Some[T](value)
	}

	self.buffer[self._ToPhysicalIndex(index)] = value
	return None[T]()
}

// Set value at index, without doing bounds checking.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.SetUnchecked(gost.USize(0), gost.I32(4))
//	gost.AssertEqual(deque.Get(gost.USize(0)), gost.Some[gost.I32](gost.I32(4)))
func (self *VecDeque[T]) SetUnchecked(index USize, value T) {
	self.buffer[self._ToPhysicalIndex(index)] = value
}

// Provides a reference to the back element, or None if the deque is empty.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.Back(), gost.Some[gost.I32](gost.I32(4)))
func (self VecDeque[T]) Back() Option[T] {
	return self.Get(self.Len().WrappingSub(1))
}

// Provides a reference to the front element, or None if the deque is empty.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.Front(), gost.Some[gost.I32](gost.I32(3)))
func (self VecDeque[T]) Front() Option[T] {
	return self.Get(0)
}

// Clears the deque, removing all values.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	deque.Clear()
//	gost.AssertEqual(deque.Len(), gost.USize(0))
func (self *VecDeque[T]) Clear() {
	self.len = 0
	self.head = 0
	self.buffer = make([]T, _VECDEQUE_INITIAL_CAPACITY)
}

// Extracts a slice containing the entire vecdeque.
func (self VecDeque[T]) AsSlice() []T {
	return self.buffer[:self.len]
}

// Require `impl Eq[T] for T`
// Returns true if the deque contains an element equal to the given value.
// This operation is O(n).
// Note that if you have a sorted VecDeque, binary_search may be faster.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.PushBack(gost.I32(3))
//	deque.PushBack(gost.I32(4))
//	gost.AssertEqual(deque.Contains(gost.I32(3)), gost.Bool(true))
//	gost.AssertEqual(deque.Contains(gost.I32(4)), gost.Bool(true))
//	gost.AssertEqual(deque.Contains(gost.I32(5)), gost.Bool(false))
func (self VecDeque[T]) Contains(value T) Bool {
	equalableValue := castToEq(value)

	if equalableValue.IsNone() {
		typeName := getTypeName(value)
		panic(fmt.Sprintf("'%s' does not implement Eq[%s]", typeName, typeName))
	}

	for i := USize(0); i < self.Len(); i++ {
		if equalableValue.Unwrap().Eq(self.Get(i).Unwrap()) {
			return true
		}
	}

	return false
}

// Fills self with elements by cloning value.
//
//	deque := gost.VecDequeWithLen[gost.I32](gost.USize(5))
//	deque.Fill(gost.I32(1))
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(1))
//	gost.AssertEq(deque.GetUnchecked(1), gost.I32(1))
func (self *VecDeque[T]) Fill(value T) {
	for i := USize(0); i < self.Len(); i++ {
		self.SetUnchecked(i, value)
	}
}

// Fills self with elements returned by calling a closure repeatedly.
//
//	deque := gost.VecDequeWithLen[gost.I32](gost.USize(5))
//	deque.FillWith(func() gost.I32 {
//		return 1
//	})
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(1))
//	gost.AssertEq(deque.GetUnchecked(1), gost.I32(1))
func (self *VecDeque[T]) FillWith(f func() T) {
	for i := USize(0); i < self.Len(); i++ {
		self.SetUnchecked(i, f())
	}
}

// Binary searches this slice for a given element. If the slice is not sorted, the returned result is unspecified and meaningless.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(1)
//	deque.Push(2)
//	deque.Push(3)
//	gost.Assert(deque.BinarySearch(2).IsSome())
//	gost.Assert(deque.BinarySearch(4).IsNone())
func (self VecDeque[T]) BinarySearch(value T) Option[USize] {
	low := 0
	high := int(self.len) - 1

	for low <= high {
		mid := (low + high) / 2

		midData := castToOrd[T](self.buffer[mid]).Unwrap()
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
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(1)
//	deque.Push(2)
//	deque.Push(3)
//	gost.Assert(deque.BinarySearchBy(func(e gost.I32) gost.Ordering {
//		if e < 2 {
//			return gost.OrderingLess
//		} else if e > 2 {
//			return gost.OrderingGreater
//		} else {
//			return gost.OrderingEqual
//		}
//	}).IsSome())
func (self VecDeque[T]) BinarySearchBy(f func(T) Ordering) Option[USize] {
	low := 0
	high := int(self.len) - 1

	for low <= high {
		mid := (low + high) / 2
		ordering := f(self.buffer[mid])

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

// Sorts the slice.
// This sort is stable (i.e., does not reorder equal elements) and O(n * log(n)) worst-case.
// When applicable, unstable sorting is preferred because it is generally faster than stable sorting and it doesn’t allocate auxiliary memory. See sort_unstable.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(3)
//	deque.Push(2)
//	deque.Push(1)
//	deque.Sort()
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(1))
//	gost.AssertEq(deque.GetUnchecked(1), gost.I32(2))
//	gost.AssertEq(deque.GetUnchecked(2), gost.I32(3))
func (self *VecDeque[T]) Sort() {
	sort.SliceStable(self.buffer[:self.len], func(i, j int) bool {
		if castToOrd[T](self.buffer[i]).IsNone() {
			typeName := getTypeName(self.buffer[i])
			panic(fmt.Sprintf("'%s' does not implement Ord[%s]", typeName, typeName))
		}

		lhs := castToOrd[T](self.buffer[i]).Unwrap()
		rhs := self.buffer[j]

		return lhs.Cmp(rhs) == OrderingLess
	})
}

// Sorts the slice with a comparator function.
// This sort is stable (i.e., does not reorder equal elements) and O(n * log(n)) worst-case.
// The comparator function must define a total ordering for the elements in the slice. If the ordering is not total, the order of the elements is unspecified. An order is a total order if it is (for all a, b and c):
// - total and antisymmetric: exactly one of a < b, a == b or a > b is true, and
// - transitive, a < b and b < c implies a < c. The same must hold for both == and >.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(3)
//	deque.Push(2)
//	deque.Push(1)
//	deque.SortBy(func(lhs, rhs gost.I32) gost.Ordering {
//		if lhs < rhs {
//			return gost.OrderingLess
//		} else if lhs > rhs {
//			return gost.OrderingGreater
//		} else {
//			return gost.OrderingEqual
//		}
//	})
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(1))
func (self *VecDeque[T]) SortBy(f func(T, T) Ordering) {
	sort.SliceStable(self.buffer[:self.len], func(i, j int) bool {
		return f(self.buffer[i], self.buffer[j]) == OrderingLess
	})
}

// Sorts the slice, but might not preserve the order of equal elements.
// This sort is unstable (i.e., may reorder equal elements), in-place (i.e., does not allocate), and O(n * log(n)) worst-case.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(3)
//	deque.Push(2)
//	deque.Push(1)
//	deque.SortUnstable()
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(1))
func (self *VecDeque[T]) SortUnstable() {
	sort.Slice(self.buffer[:self.len], func(i, j int) bool {
		if castToOrd[T](self.buffer[i]).IsNone() {
			typeName := getTypeName(self.buffer[i])
			panic(fmt.Sprintf("'%s' does not implement Ord[%s]", typeName, typeName))
		}

		lhs := castToOrd[T](self.buffer[i]).Unwrap()
		rhs := self.buffer[j]

		return lhs.Cmp(rhs) == OrderingLess
	})
}

// Sorts the slice with a comparator function, but might not preserve the order of equal elements.
// This sort is unstable (i.e., may reorder equal elements), in-place (i.e., does not allocate), and O(n * log(n)) worst-case.
// The comparator function must define a total ordering for the elements in the slice. If the ordering is not total, the order of the elements is unspecified. An order is a total order if it is (for all a, b and c):
// - total and antisymmetric: exactly one of a < b, a == b or a > b is true, and
// - transitive, a < b and b < c implies a < c. The same must hold for both == and >.
//
//	deque := gost.VecDequeNew[gost.I32]()
//	deque.Push(3)
//	deque.Push(2)
//	deque.Push(1)
//	deque.SortUnstableBy(func(lhs, rhs gost.I32) gost.Ordering {
//		if lhs < rhs {
//			return gost.OrderingLess
//		} else if lhs > rhs {
//			return gost.OrderingGreater
//		} else {
//			return gost.OrderingEqual
//		}
//	})
//	gost.AssertEq(deque.GetUnchecked(0), gost.I32(1))
func (self *VecDeque[T]) SortUnstableBy(f func(T, T) Ordering) {
	sort.Slice(self.buffer[:self.len], func(i, j int) bool {
		return f(self.buffer[i], self.buffer[j]) == OrderingLess
	})
}

// Returns `true` if the buffer is at full capacity.
func (self VecDeque[T]) _IsFull() bool {
	return self.len == USize(len(self.buffer))
}

// Double the buffer size.
// This method is inline(never), so we expect it to only be called in cold paths.
// This may panic or abort
func (self *VecDeque[T]) _Grow() {
	if !self._IsFull() {
		panic("VecDeque._Grow: VecDeque is not full")
	}

	oldCapacity := USize(len(self.buffer))
	newCapacity := oldCapacity * 2

	newBuffer := make([]T, newCapacity)
	copy(newBuffer, self.buffer)

	self.buffer = newBuffer

	self._HandleCapacityIncrease(oldCapacity)

	if self._IsFull() {
		panic("VecDeque._Grow: VecDeque is full")
	}
}

// Frobs the head and tail sections around to handle the fact that we just reallocated.
// Unsafe because it trusts old_capacity.
func (self *VecDeque[T]) _HandleCapacityIncrease(oldCapacity USize) {
	newCapacity := USize(len(self.buffer))
	if newCapacity < oldCapacity {
		panic("VecDeque._HandleCapacityIncrease: newCapacity < oldCapacity")
	}

	// Move the shortest contiguous section of the ring buffer
	//
	// H := head
	// L := last element (`self.to_physical_idx(self.len - 1)`)
	//
	//    H           L
	//   [o o o o o o o . ]
	//    H           L
	// A [o o o o o o o . . . . . . . . . ]
	//        L H
	//   [o o o o o o o o ]
	//          H           L
	// B [. . . o o o o o o o . . . . . . ]
	//              L H
	//   [o o o o o o o o ]
	//            L                   H
	// C [o o o o o . . . . . . . . . o o ]

	// can't use is_contiguous() because the capacity is already updated.
	if self.head <= oldCapacity-self.len {
		// A
		// Nop
	} else {
		headLen := oldCapacity - self.head
		tailLen := self.len - headLen

		if headLen > tailLen && newCapacity-oldCapacity >= tailLen {
			// B
			self._CopyNonoverlapping(USize(0), oldCapacity, tailLen)
		} else {
			// C
			newHead := newCapacity - headLen

			// can't use copy_nonoverlapping here, because if e.g. head_len = 2
			// and new_capacity = old_capacity + 1, then the heads overlap.
			self._Copy(self.head, newHead, headLen)
			self.head = newHead
		}
	}
}

// Copies a contiguous block of memory len long from src to dst
func (self *VecDeque[T]) _Copy(src USize, dst USize, len USize) {
	copy(self.buffer[dst:dst+len], self.buffer[src:])
}

// Copies a contiguous block of memory len long from src to dst
func (self *VecDeque[T]) _CopyNonoverlapping(src USize, dst USize, len USize) {
	copy(self.buffer[dst:dst+len], self.buffer[src:])
}

// / Returns the index in the underlying buffer for a given logical element index.
func _WrapIndex(logicalIndex USize, capacity USize) USize {
	if logicalIndex >= capacity {
		return logicalIndex - capacity
	} else {
		return logicalIndex
	}
}

// Returns the index in the underlying buffer for a given logical element
// index + addend.
func (self VecDeque[T]) _WrapAdd(index USize, addend USize) USize {
	capacity := USize(len(self.buffer))

	return _WrapIndex(index.WrappingAdd(addend), capacity)
}

// Returns the index in the underlying buffer for a given logical element
// index - subtrahend.
func (self VecDeque[T]) _WrapSub(index USize, subtrahend USize) USize {
	capacity := USize(len(self.buffer))

	return _WrapIndex(index.WrappingSub(subtrahend).WrappingAdd(capacity), capacity)
}

func (self VecDeque[T]) _ToPhysicalIndex(index USize) USize {
	return self._WrapAdd(self.head, index)
}

// iterator for VecDeque
type VecDequeIter[T any] struct {
	deque    *VecDeque[T]
	position USize
}

// into_iter
func (self VecDeque[T]) IntoIter() Iterator[T] {
	return &VecDequeIter[T]{deque: &self, position: 0}
}

// next
func (self *VecDequeIter[T]) Next() Option[T] {
	if self.position >= self.deque.Len() {
		return None[T]()
	}

	value := self.deque.Get(self.position)
	self.position++

	return value
}

// map
func (self *VecDequeIter[T]) Map(f func(T) T) Iterator[T] {
	newDeque := VecDequeWithCapacity[T](self.deque.Capacity())

	for i := USize(0); i < self.deque.Len(); i++ {
		newDeque.PushBack(f(self.deque.Get(i).Unwrap()))
	}

	return newDeque.IntoIter()
}

// filter
func (self *VecDequeIter[T]) Filter(f func(T) Bool) Iterator[T] {
	newDeque := VecDequeWithCapacity[T](self.deque.Capacity())

	for i := USize(0); i < self.deque.Len(); i++ {
		value := self.deque.Get(i).Unwrap()
		if f(value) {
			newDeque.PushBack(value)
		}
	}

	return newDeque.IntoIter()
}

// fold
func (self *VecDequeIter[T]) Fold(initial T, f func(T, T) T) T {
	accumulator := initial

	for i := USize(0); i < self.deque.Len(); i++ {
		accumulator = f(accumulator, self.deque.Get(i).Unwrap())
	}

	return accumulator
}

// rev
func (self *VecDequeIter[T]) Rev() Iterator[T] {
	newDeque := VecDequeWithCapacity[T](self.deque.Capacity())

	for i := self.deque.Len().WrappingSub(1); i >= 0; i-- {
		newDeque.PushBack(self.deque.Get(i).Unwrap())
	}

	return newDeque.IntoIter()
}

// collect to Vec
func (self *VecDequeIter[T]) CollectToVec() Vec[T] {
	newVec := VecWithCapacity[T](self.deque.Capacity())

	for i := USize(0); i < self.deque.Len(); i++ {
		newVec.Push(self.deque.Get(i).Unwrap())
	}

	return newVec
}

// collect to LinkedList
func (self *VecDequeIter[T]) CollectToLinkedList() LinkedList[T] {
	newLinkedList := LinkedListNew[T]()

	for i := USize(0); i < self.deque.Len(); i++ {
		newLinkedList.PushBack(self.deque.Get(i).Unwrap())
	}

	return newLinkedList
}

// impl Display for VecDeque
func (self VecDeque[T]) Display() String {
	buffer := String("")
	buffer += "VecDeque["

	for i := USize(0); i < self.Len(); i++ {
		e := self.Get(i).Unwrap()

		display := castToDisplay(e)
		if display.IsSome() {
			buffer += display.Unwrap().Display()
		} else {
			typeName := getTypeName(e)

			panic(fmt.Sprintf("'%s' does not implement Display[%s]", typeName, typeName))
		}

		if i != self.Len()-1 {
			buffer += ", "
		}
	}

	buffer += "]"

	return String(buffer)
}

// impl Debug for VecDeque
func (self VecDeque[T]) Debug() String {
	return self.Display()
}

// impl AsRef for VecDeque
func (self VecDeque[T]) AsRef() *VecDeque[T] {
	return &self
}

// impl Clone for VecDeque
func (self VecDeque[T]) Clone() *VecDeque[T] {
	cloned := VecDequeWithCapacity[T](self.Capacity())

	for i := USize(0); i < self.Len(); i++ {
		cloned.PushBack(self.Get(i).Unwrap())
	}

	return &cloned
}

// impl Eq for Vec
func (self VecDeque[T]) Eq(other VecDeque[T]) Bool {
	if self.Len() != other.Len() {
		return false
	}

	for i := USize(0); i < self.Len(); i++ {
		if !castToEq(self.Get(i).Unwrap()).Unwrap().Eq(other.Get(i).Unwrap()) {
			return false
		}
	}

	return true
}
