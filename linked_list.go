package gost

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
}

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
	prev  *LinkedListNode[T]
}

// Creates an empty LinkedList.
func LinkedListNew[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

// Appends an element to the back of a list.
// This operation should compute in O(1) time.
func (list *LinkedList[T]) PushBack(value T) {
	newNode := LinkedListNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}

	newNode.prev = list.tail

	if list.tail != nil {
		list.tail.next = &newNode
	}

	list.tail = &newNode

	if list.head == nil {
		list.head = &newNode
	}
}

// Removes the last element from a list and returns it, or None if it is empty.
// This operation should compute in O(1) time.
func (list *LinkedList[T]) PopBack() Option[T] {
	if list.tail == nil {
		return None[T]()
	}

	value := list.tail.value
	list.tail = list.tail.prev

	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}

	return Some[T](value)
}

// into_iter
func (list *LinkedList[T]) IntoIter() LinkedListIter[T] {
	return LinkedListIter[T]{
		pointer: list.head,
	}
}

type LinkedListIter[T any] struct {
	pointer *LinkedListNode[T]
}

// next
func (self *LinkedListIter[T]) Next() Option[T] {
	if self.pointer == nil {
		return None[T]()
	}

	value := self.pointer.value
	self.pointer = self.pointer.next

	return Some[T](value)
}

// map
func (self LinkedListIter[T]) Map(f func(T) T) Iterator[T] {
	currentIter := self

	newList := LinkedListNew[T]()

	for {
		value := currentIter.Next()

		if value.IsNone() {
			break
		}

		newList.PushBack(f(value.Unwrap()))
	}

	return newList.IntoIter()
}

// filter
func (self LinkedListIter[T]) Filter(f func(T) bool) Iterator[T] {
	currentIter := self

	newList := LinkedListNew[T]()

	for {
		value := currentIter.Next()

		if value.IsNone() {
			break
		}

		if f(value.Unwrap()) {
			newList.PushBack(value.Unwrap())
		}
	}

	return newList.IntoIter()
}

// fold
func (self LinkedListIter[T]) Fold(init T, f func(T, T) T) T {
	currentIter := self

	result := init

	for {
		value := currentIter.Next()

		if value.IsNone() {
			break
		}

		result = f(result, value.Unwrap())
	}

	return result
}
