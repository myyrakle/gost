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
