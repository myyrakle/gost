package gost

// A priority queue implemented with a binary heap.
// This will be a max-heap.
// It is a logic error for an item to be modified in such a way that the item’s ordering relative to any other item, as determined by the Ord trait, changes while it is in the heap. This is normally only possible through interior mutability, global state, I/O, or unsafe code. The behavior resulting from such a logic error is not specified, but will be encapsulated to the BinaryHeap that observed the logic error and not result in undefined behavior. This could include panics, incorrect results, aborts, memory leaks, and non-termination.
// As long as no elements change their relative order while being in the heap as described above, the API of BinaryHeap guarantees that the heap invariant remains intact i.e. its methods all behave as documented. For example if a method is documented as iterating in sorted order, that’s guaranteed to work as long as elements in the heap have not changed order, even in the presence of closures getting unwinded out of, iterators getting leaked, and similar foolishness.
type BinaryHeap[T Ord[T]] struct {
	vec Vec[T]
}

// Constructs a new, empty BinaryHeapNew<T>.
func BinaryHeapNew[T Ord[T]]() BinaryHeap[T] {
	return BinaryHeap[T]{
		vec: VecNew[T](),
	}
}

// Constructs a new, empty BinaryHeap<T> with at least the specified capacity.
func BinaryHeapWithCapacity[T Ord[T]](capacity USize) BinaryHeap[T] {
	return BinaryHeap[T]{
		vec: VecWithCapacity[T](capacity),
	}
}

// Constructs a new, empty BinaryHeap<T> with at least the specified len.
func BinaryHeapWithLen[T Ord[T]](len USize) BinaryHeap[T] {
	return BinaryHeap[T]{
		vec: VecWithLen[T](len),
	}
}

// Returns the total number of elements the heap can hold without reallocating.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	gost.AssertEq(heap.Capacity(), gost.USize(0))
//
//	heap.Reserve(10)
//	gost.AssertEq(heap.Capacity(), gost.USize(10))
func (self BinaryHeap[T]) Capacity() USize {
	return self.Capacity()
}

// Returns the number of elements in the heap, also referred to as its ‘length’.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	gost.AssertEq(heap.Len(), gost.USize(0))
//
//	heap.Push(1)
//	gost.AssertEq(heap.Len(), gost.USize(1))
func (self BinaryHeap[T]) Len() USize {
	return self.Len()
}

// Returns true if the heap contains no elements.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	gost.Assert(heap.IsEmpty())
//
//	heap.Push(1)
//	gost.Assert(!heap.IsEmpty())
func (self BinaryHeap[T]) IsEmpty() Bool {
	return self.IsEmpty()
}

// Reserves capacity for at least additional more elements to be inserted in the given Vec<T>. The collection may reserve more space to speculatively avoid frequent reallocations. After calling reserve, capacity will be greater than or equal to self.len() + additional. Does nothing if capacity is already sufficient.
//
//	deque := gost.BinaryHeapNew[gost.I32]()
//	gost.AssertEq(deque.Capacity(), gost.USize(0))
//
//	deque.Reserve(10)
//	gost.AssertEq(deque.Capacity(), gost.USize(10))
func (self *BinaryHeap[T]) Reserve(capacity USize) {
	self.Reserve(capacity)
}

// Pushes an item onto the binary heap.
//
func (self *BinaryHeap[T]) Push(item T) {
	Todo()
}

// Removes the greatest item from the binary heap and returns it, or None if it is empty.
//
func (self *BinaryHeap[T]) Pop() Option[T] {
	panic("")
}

// Moves all the elements of other into self, leaving other empty.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	heap.Push(1)
//	heap.Push(2)
//	heap.Push(3)
//	heap2 := gost.BinaryHeapNew[gost.I32]()
//	heap2.Push(4)
//	heap2.Push(5)
//  heap.Append(&heap2)
//	gost.AssertEq(heap.Len(), gost.USize(5))
//	gost.AssertEq(heap2.Len(), gost.USize(0))
func (self *BinaryHeap[T]) Append(other *BinaryHeap[T]) {
	self.Reserve(self.Len() + other.Len())
	for _, item := range other.vec.data {
		self.Push(item)
	}
}

// Clears the heap, removing all values.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	heap.Push(1)
//	heap.Push(2)
//	heap.Push(3)
//	heap.Clear()
//	gost.AssertEq(heap.Len(), gost.USize(0))
func (self *BinaryHeap[T]) Clear() {
	self.Clear()
}
