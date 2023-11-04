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
	return self.vec.Capacity()
}

// Returns the number of elements in the heap, also referred to as its ‘length’.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	gost.AssertEq(heap.Len(), gost.USize(0))
//
//	heap.Push(1)
//	gost.AssertEq(heap.Len(), gost.USize(1))
func (self BinaryHeap[T]) Len() USize {
	return self.vec.Len()
}

// Returns true if the heap contains no elements.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	gost.Assert(heap.IsEmpty())
//
//	heap.Push(1)
//	gost.Assert(!heap.IsEmpty())
func (self BinaryHeap[T]) IsEmpty() Bool {
	return self.vec.IsEmpty()
}

// Reserves capacity for at least additional more elements to be inserted in the given Vec<T>. The collection may reserve more space to speculatively avoid frequent reallocations. After calling reserve, capacity will be greater than or equal to self.len() + additional. Does nothing if capacity is already sufficient.
//
//	deque := gost.BinaryHeapNew[gost.I32]()
//	gost.AssertEq(deque.Capacity(), gost.USize(0))
//
//	deque.Reserve(10)
//	gost.AssertEq(deque.Capacity(), gost.USize(10))
func (self *BinaryHeap[T]) Reserve(capacity USize) {
	self.vec.Reserve(capacity)
}

// Pushes an item onto the binary heap.
//
// heap := gost.BinaryHeapNew[gost.I32]()
// heap.Push(1)
// gost.AssertEq(heap.Len(), gost.USize(1))
func (self *BinaryHeap[T]) Push(item T) {
	oldLen := self.Len()
	self.vec.Push(item)

	// SAFETY: Since we pushed a new item it means that
	//  old_len = self.len() - 1 < self.len()
	self._SiftUp(0, oldLen)
}

func (self *BinaryHeap[T]) _SiftUp(start USize, pos USize) USize {
	// Take out the value at `pos` and create a hole.
	// SAFETY: The caller guarantees that pos < self.len()
	hole := _HoleNew[T](&self.vec, pos)

	for hole.Pos() > start {
		parent := (hole.Pos() - 1) / 2

		// SAFETY: hole.pos() > start >= 0, which means hole.pos() > 0
		//  and so hole.pos() - 1 can't underflow.
		//  This guarantees that parent < hole.pos() so
		//  it's a valid index and also != hole.pos().
		order := hole.Element().Cmp(hole.Get(parent))
		if order == OrderingLess || order == OrderingEqual {
			break
		}

		// SAFETY: Same as above
		hole.MoveTo(parent)
	}

	return hole.Pos()
}

// Removes the greatest item from the binary heap and returns it, or None if it is empty.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	heap.Push(1)
//	heap.Push(2)
//	heap.Push(3)
//  gost.AssertEq(heap.Len(), gost.I32(3))
//  gost.AssertEq(heap.Pop(), gost.Some[gost.I32](3))
//  gost.AssertEq(heap.Len(), gost.I32(2))
func (self *BinaryHeap[T]) Pop() Option[T] {
	return self.vec.Pop().Map(func(item T) T {
		if !self.IsEmpty() {
			_Swap(&item, &self.vec.data[0])
			// SAFETY: !self.is_empty() means that self.len() > 0
			self._SiftDownToBottom(0)
		}

		return item
	})
}

func (self *BinaryHeap[T]) _SiftDownToBottom(pos USize) {
	end := self.Len()
	start := pos

	// SAFETY: pos < self.len() so pos + 1 <= self.len()
	hole := _HoleNew[T](&self.vec, pos)
	child := 2*hole.Pos() + 1

	// Loop invariant: child == 2 * hole.pos() + 1.
	for child <= end.SaturatingSub(2) {
		// SAFETY: child < end - 1 < self.len() and
		//  child + 1 < end <= self.len(), so they're valid indexes.
		//  child == 2 * hole.pos() + 1 != hole.pos() and
		//  child + 1 == 2 * hole.pos() + 2 != hole.pos().
		// FIXME: 2 * hole.pos() + 1 or 2 * hole.pos() + 2 could overflow
		//  if T is a ZST
		ordering := hole.Get(child).Cmp(hole.Get(child + 1))
		if ordering == OrderingLess || ordering == OrderingEqual {
			child += 1
		}

		// SAFETY: Same as above
		hole.MoveTo(child)
		child = 2*hole.Pos() + 1
	}

	if child == end-1 {
		// SAFETY: child == end - 1 < self.len(), so it's a valid index
		//  and child == 2 * hole.pos() + 1 != hole.pos().
		hole.MoveTo(child)
	}
	pos = hole.Pos()

	// SAFETY: pos is the position in the hole and was already proven
	//  to be a valid index.
	self._SiftUp(start, pos)
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
	self.vec.Clear()
}

// Extracts a slice containing the entire heap.
func (self *BinaryHeap[T]) AsSlice() []T {
	return self.vec.AsSlice()
}

// Consumes the BinaryHeap and returns the underlying vector in arbitrary order.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	heap.Push(1)
//	heap.Push(2)
//	heap.Push(3)
//  vec := gost.VecNew[gost.I32]()
//  vec.Push(3)
//  vec.Push(1)
//  vec.Push(2)
//	gost.AssertEq(heap.IntoVec(), vec)
func (self *BinaryHeap[T]) IntoVec() Vec[T] {
	return self.vec
}

// Consumes the BinaryHeap and returns a vector in sorted (ascending) order.
//
//	heap := gost.BinaryHeapNew[gost.I32]()
//	heap.Push(1)
//	heap.Push(3)
//	heap.Push(2)
//  vec := gost.VecNew[gost.I32]()
//  vec.Push(1)
//  vec.Push(2)
//  vec.Push(3)
//	gost.AssertEq(heap.IntoSortedVec(), vec)
func (self *BinaryHeap[T]) IntoSortedVec() Vec[T] {
	sortedVec := VecWithLen[T](self.Len())

	index := self.Len() - 1

	for !self.IsEmpty() {
		sortedVec.SetUnchecked(index, self.Pop().Unwrap())

		index -= 1
	}

	return sortedVec
}

/// Hole represents a hole in a slice i.e., an index without valid value
/// (because it was moved from or duplicated).
/// In drop, `Hole` will restore the slice by filling the hole
/// position with the value that was originally removed.
type _Hole[T any] struct {
	data *Vec[T]
	elt  T
	pos  USize
}

func _HoleNew[T any](data *Vec[T], pos USize) _Hole[T] {
	elt := data.GetUnchecked(pos)
	return _Hole[T]{
		data: data,
		elt:  elt,
		pos:  pos,
	}
}

func (self _Hole[T]) Pos() USize {
	return self.pos
}

// Returns a reference to the element removed.
func (self _Hole[T]) Element() T {
	return self.elt
}

/// Returns a reference to the element at `index`.
/// Unsafe because index must be within the data slice and not equal to pos.
func (self _Hole[T]) Get(pos USize) T {
	return self.data.data[pos]
}

// Move hole to new location
// Unsafe because index must be within the data slice and not equal to pos.
func (self *_Hole[T]) MoveTo(index USize) {
	self.data.Swap(self.pos, index)
	self.pos = index
}
