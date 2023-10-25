package gost

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
func (self VecDeque[T]) Len() USize {
	return USize(self.len)
}

// Prepends an element to the deque.
func (self *VecDeque[T]) PushFront(value T) {
	if self._IsFull() {
		self._Grow()
	}

	self.head = self._WrapAdd(self.head, -1)
	self.buffer[self.head] = value
	self.len++
}

// Appends an element to the back of the deque.
func (self *VecDeque[T]) PushBack(value T) {
	if self._IsFull() {
		self._Grow()
	}

	self.head = self._WrapAdd(self.head, self.len)
	self.buffer[self._ToPhysicalIndex(self.len)] = value
	self.len++
}

func (self VecDeque[T]) _IsFull() bool {
	return self.len == USize(len(self.buffer))
}

func (self *VecDeque[T]) _Grow() {
	if !self._IsFull() {
		panic("VecDeque._Grow: VecDeque is not full")
	}

	oldCapacity := uint(len(self.buffer))
	newCapacity := oldCapacity * 2

	newBuffer := make([]T, newCapacity)
	copy(newBuffer, self.buffer)

	self.buffer = newBuffer

	if self._IsFull() {
		panic("VecDeque._Grow: VecDeque is full")
	}
}

func _WrapIndex(logicalIndex USize, capacity USize) USize {
	if logicalIndex >= capacity {
		return logicalIndex - capacity
	} else {
		return logicalIndex
	}
}

func (self VecDeque[T]) _WrapAdd(index USize, subtrahend USize) USize {
	return _WrapIndex(index.WrappingSub(subtrahend), USize(len(self.buffer)))
}

func (self VecDeque[T]) _WrapSub(index USize, subtrahend USize) USize {
	capacity := USize(len(self.buffer))

	return _WrapIndex(index.WrappingAdd(subtrahend).WrappingAdd(capacity), capacity)
}

func (self VecDeque[T]) _ToPhysicalIndex(index USize) USize {
	return self._WrapAdd(self.head, index)
}
