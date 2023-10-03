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
