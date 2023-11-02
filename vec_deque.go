package gost

import "fmt"

const (
	_VECDEQUE_INITIAL_CAPACITY uint = 7 // 2^3 - 1
	_VECDEQUE_MINIMUM_CAPACITY uint = 1 // 2 - 1
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
func VecDequeWithCapacity[T any](capacity uint) VecDeque[T] {
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

// Provides a reference to the element at the given index.
// Element at index 0 is the front of the queue.
func (self VecDeque[T]) Get(index USize) Option[T] {
	if index >= self.Len() {
		return None[T]()
	}

	return Some[T](self.buffer[uint(self._ToPhysicalIndex(index))])
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
