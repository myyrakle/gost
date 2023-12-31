package gost

import "fmt"

// LinkedList is a doubly-linked list.
type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	len  USize
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

// Returns the length of the LinkedList.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushBack(gost.I32(1))
//	list.PushBack(gost.I32(2))
//	gost.AssertEq(list.Len(), gost.USize(2))
func (list LinkedList[T]) Len() USize {
	return list.len
}

// Returns true if the LinkedList is empty.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	gost.Assert(list.IsEmpty())
//
//	list.PushBack(gost.I32(1))
//	gost.Assert(!list.IsEmpty())
func (list LinkedList[T]) IsEmpty() Bool {
	return list.len == 0
}

// Removes all elements from the LinkedList.
// This operation should compute in O(1) time. +GC
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushBack(gost.I32(1))
//	list.PushBack(gost.I32(2))
//	list.Clear()
//	gost.Assert(list.IsEmpty())
func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.len = 0
}

// Appends an element to the back of a list.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushBack(gost.I32(1))
//	list.PushBack(gost.I32(2))
//	gost.AssertEq(list.Len(), gost.USize(2))
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

	list.len++
}

// Removes the last element from a list and returns it, or None if it is empty.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushBack(gost.I32(1))
//	list.PushBack(gost.I32(2))
//	gost.AssertEq(list.PopBack().Unwrap(), gost.I32(2))
//	gost.AssertEq(list.PopBack().Unwrap(), gost.I32(1))
//	gost.Assert(list.PopBack().IsNone())
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

	list.len--

	return Some[T](value)
}

// Provides a reference to the back element, or None if the list is empty.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushBack(gost.I32(1))
//	list.PushBack(gost.I32(2))
//	gost.AssertEq(list.Back().Unwrap(), gost.I32(2))
func (list *LinkedList[T]) Back() Option[*T] {
	if list.tail == nil {
		return None[*T]()
	}

	return Some[*T](&list.tail.value)
}

// Adds an element first in the list.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushFront(gost.I32(1))
//	list.PushFront(gost.I32(2))
//	gost.AssertEq(list.Len(), gost.USize(2))
func (list *LinkedList[T]) PushFront(value T) {
	newNode := LinkedListNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}

	newNode.next = list.head

	if list.head != nil {
		list.head.prev = &newNode
	}

	list.head = &newNode

	if list.tail == nil {
		list.tail = &newNode
	}

	list.len++
}

// Removes the first element and returns it, or None if the list is empty.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushFront(gost.I32(1))
//	list.PushFront(gost.I32(2))
//	gost.AssertEq(list.PopFront().Unwrap(), gost.I32(2))
//	gost.AssertEq(list.PopFront().Unwrap(), gost.I32(1))
func (list *LinkedList[T]) PopFront() Option[T] {
	if list.head == nil {
		return None[T]()
	}

	value := list.head.value
	list.head = list.head.next

	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}

	list.len--

	return Some[T](value)
}

// Provides a reference to the front element, or None if the list is empty.
// This operation should compute in O(1) time.
//
//	list := gost.LinkedListNew[gost.I32]()
//	list.PushFront(gost.I32(1))
//	list.PushFront(gost.I32(2))
//	gost.AssertEq(list.Front().Unwrap(), gost.I32(2))
func (list *LinkedList[T]) Front() Option[*T] {
	if list.head == nil {
		return None[*T]()
	}

	return Some[*T](&list.head.value)
}

// Moves all elements from other to the end of the list.
// This reuses all the nodes from other and moves them into self. After this operation, other becomes empty.
// This operation should compute in O(1) time and O(1) memory.
//
//	list1 := gost.LinkedListNew[gost.I32]()
//	list1.PushBack(gost.I32(1))
//	list1.PushBack(gost.I32(2))
//	list1.PushBack(gost.I32(3))
//	list2 := gost.LinkedListNew[gost.I32]()
//	list2.PushBack(gost.I32(4))
//	list2.PushBack(gost.I32(5))
//  list1.Append(&list2)
//	gost.AssertEq(list1.Len(), gost.USize(5))
//	gost.AssertEq(list2.Len(), gost.USize(0))
func (list *LinkedList[T]) Append(other *LinkedList[T]) {
	if other.head == nil {
		return
	}

	if list.head == nil {
		list.head = other.head
		list.tail = other.tail
		list.len = other.len
	} else {
		list.tail.next = other.head
		other.head.prev = list.tail
		list.tail = other.tail
		list.len += other.len
	}

	other.Clear()
}

// into_iter
func (list *LinkedList[T]) IntoIter() Iterator[T] {
	return &LinkedListIter[T]{
		pointer: list.head,
	}
}

// linked list iterator
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
func (self LinkedListIter[T]) Filter(f func(T) Bool) Iterator[T] {
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

// rev
func (self LinkedListIter[T]) Rev() Iterator[T] {
	currentIter := self

	newList := LinkedListNew[T]()

	for {
		value := currentIter.Next()

		if value.IsNone() {
			break
		}

		newList.PushBack(value.Unwrap())
	}

	return newList.IntoIter()
}

// Collect To Vec
func (self LinkedListIter[T]) CollectToVec() Vec[T] {
	vec := Vec[T]{}
	for {
		value := self.Next()
		if value.IsNone() {
			return vec
		}
		vec.Push(value.Unwrap())
	}
}

// Collect To LinkedList
func (self LinkedListIter[T]) CollectToLinkedList() LinkedList[T] {
	list := LinkedListNew[T]()
	for {
		value := self.Next()
		if value.IsNone() {
			return list
		}
		list.PushBack(value.Unwrap())
	}
}

// impl Display for LinkedList
func (self LinkedList[T]) Display() String {
	buffer := String("")
	buffer += "LinkedList["

	iter := self.IntoIter()
	count := 0
	for {
		wrapped := iter.Next()

		if wrapped.IsNone() {
			break
		}
		e := wrapped.Unwrap()

		display := castToDisplay(e)
		if display.IsSome() {
			buffer += display.Unwrap().Display()
		} else {
			typeName := getTypeName(e)

			panic(fmt.Sprintf("'%s' does not implement Display[%s]", typeName, typeName))
		}

		if count < int(self.Len())-1 {
			buffer += ", "
		}

		count++
	}

	buffer += "]"

	return String(buffer)
}

// impl Debug for LinkedList
func (self LinkedList[T]) Debug() String {
	return self.Display()
}

// impl AsRef for LinkedList
func (self LinkedList[T]) AsRef() *LinkedList[T] {
	return &self
}

// impl Clone for LinkedList
func (self LinkedList[T]) Clone() LinkedList[T] {
	newList := LinkedListNew[T]()

	iter := self.IntoIter()
	for {
		wrapped := iter.Next()

		if wrapped.IsNone() {
			break
		}
		e := wrapped.Unwrap()

		clone := castToClone(e)
		if clone.IsSome() {
			newList.PushBack(clone.Unwrap().Clone())
		} else {
			typeName := getTypeName(e)

			panic(fmt.Sprintf("'%s' does not implement Clone[%s]", typeName, typeName))
		}
	}

	return newList
}

// impl Eq for LinkedList
func (self LinkedList[T]) Eq(rhs LinkedList[T]) Bool {
	if self.Len() != rhs.Len() {
		return false
	}

	iter1 := self.IntoIter()
	iter2 := rhs.IntoIter()

	for {
		wrapped1 := iter1.Next()
		wrapped2 := iter2.Next()

		if wrapped1.IsNone() {
			break
		}
		if wrapped2.IsNone() {
			break
		}

		e1 := wrapped1.Unwrap()
		e2 := wrapped2.Unwrap()

		eq := castToEq(e1)
		if eq.IsSome() {
			if !eq.Unwrap().Eq(e2) {
				return false
			}
		} else {
			typeName := getTypeName(e1)

			panic(fmt.Sprintf("'%s' does not implement Eq[%s]", typeName, typeName))
		}
	}

	return true
}
