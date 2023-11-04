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
