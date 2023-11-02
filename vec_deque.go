package gost

import "fmt"

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

	fmt.Println("#$", self.buffer)
	self._HandleCapacityIncrease(oldCapacity)

	fmt.Println("%%%%", self.buffer)

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
	return &VecDequeIter[T]{deque: self, position: 0}
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
func (self *VecDequeIter[T]) Filter(f func(T) bool) Iterator[T] {
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
