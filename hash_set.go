package gost

import (
	"fmt"
	"strings"
)

type HashSet[K comparable] struct {
	hashMap HashMap[K, struct{}]
}

// Creates an empty HashSet.
func HashSetNew[K comparable]() HashSet[K] {
	return HashSet[K]{hashMap: HashMapNew[K, struct{}]()}
}

// Creates an empty HashSet with at least the specified capacity.
func HashSetWithCapacity[K comparable](capacity USize) HashSet[K] {
	return HashSet[K]{hashMap: HashMapWithCapacity[K, struct{}](capacity)}
}

// As Slice
func (self HashSet[K]) AsSlice() []K {
	result := make([]K, 0, self.hashMap.Len())
	for key := range self.hashMap.data {
		result = append(result, key)
	}

	return result
}

// From Slice
func HashSetFromSlice[K comparable](slice []K) HashSet[K] {
	result := HashSetNew[K]()
	for _, value := range slice {
		result.Insert(value)
	}

	return result
}

// Returns the number of elements in the set.
func (self HashSet[K]) Len() USize {
	return self.hashMap.Len()
}

// Returns true if the set contains no elements.
func (self HashSet[K]) IsEmpty() Bool {
	return self.hashMap.IsEmpty()
}

// Adds a value to the set.
// Returns whether the value was newly inserted. That is:
// If the set did not previously contain this value, true is returned.
// If the set already contained this value, false is returned, and the set is not modified: original value is not replaced, and the value passed as argument is dropped.
func (self *HashSet[K]) Insert(value K) Bool {
	result := self.hashMap.Insert(value, struct{}{})
	return result.IsNone()
}

// Removes a value from the set. Returns whether the value was present in the set.
// The value may be any borrowed form of the set’s value type, but Hash and Eq on the borrowed form must match those for the value type.
func (self *HashSet[K]) Remove(value K) Bool {
	result := self.hashMap.Remove(value)
	return result.IsSome()
}

// Clears the set, removing all values.
func (self *HashSet[K]) Clear() {
	self.hashMap.Clear()
}

// Returns a reference to the value in the set, if any, that is equal to the given value.
// The value may be any borrowed form of the set’s value type, but Hash and Eq on the borrowed form must match those for the value type.
func (self HashSet[K]) Get(value K) Option[K] {
	result := self.hashMap.Get(value)
	if result.IsSome() {
		return Some[K](value)
	} else {
		return None[K]()
	}
}

// Returns true if the set contains a value.
// The value may be any borrowed form of the set’s value type, but Hash and Eq on the borrowed form must match those for the value type.
func (self HashSet[K]) Contains(value K) Bool {
	return self.hashMap.ContainsKey(value)
}

type HashSetIter[K comparable] struct {
	vec      Vec[K]
	position USize
}

// into_iter
func (self HashSet[K]) IntoIter() Iterator[K] {
	vec := Vec[K]{}
	for key := range self.hashMap.data {
		vec.Push(key)
	}

	return &HashSetIter[K]{vec: vec, position: 0}
}

// next
func (self *HashSetIter[K]) Next() Option[K] {
	if self.position >= self.vec.Len() {
		return None[K]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[K](value)
}

// map
func (self HashSetIter[K]) Map(f func(K) K) Iterator[K] {
	newVec := VecNew[K]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.Push(f(value.Unwrap()))
	}
}

// filter
func (self HashSetIter[K]) Filter(f func(K) Bool) Iterator[K] {
	newVec := VecNew[K]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}

		unwraped := value.Unwrap()
		if f(unwraped) {
			newVec.Push(unwraped)
		}
	}
}

// fold
func (self HashSetIter[K]) Fold(init K, f func(K, K) K) K {
	for {
		value := self.Next()

		if value.IsNone() {
			return init
		}

		init = f(init, value.Unwrap())
	}
}

// rev
func (self HashSetIter[K]) Rev() Iterator[K] {
	newVec := VecWithLen[K](self.vec.Len())
	i := self.vec.Len() - 1

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.AsSlice()[i] = value.Unwrap()
		i--
	}
}

// Collect to Vec
func (self HashSetIter[K]) CollectToVec() Vec[K] {
	return self.vec
}

// Collect to LinkedList
func (self HashSetIter[K]) CollectToLinkedList() LinkedList[K] {
	list := LinkedListNew[K]()

	for {
		value := self.Next()

		if value.IsNone() {
			return list
		}
		list.PushBack(value.Unwrap())
	}
}

// impl Display for HashSet
func (self HashSet[K]) Display() String {
	buffer := String("")
	buffer += "HashSet{"

	fields := []string{}
	for key := range self.hashMap.data {
		fields = append(fields, string(Format("{}", key)))
	}

	buffer += String(strings.Join(fields, ", "))

	buffer += "}"

	return buffer
}

// impl Debug for HashSet
func (self HashSet[K]) Debug() String {
	return self.Display()
}

// impl AsRef for HashSet
func (self HashSet[K]) AsRef() *HashSet[K] {
	return &self
}

// impl Clone for HashSet
func (self HashSet[K]) Clone() HashSet[K] {
	result := HashSetNew[K]()
	for key := range self.hashMap.data {
		cloneKey := castToClone[K](key)

		if cloneKey.IsSome() {
			result.Insert(cloneKey.Unwrap().Clone())
		} else {
			typeName := getTypeName(key)
			panic(fmt.Sprintf("'%s' does not implement Clone[%s]", typeName, typeName))
		}

	}

	return result
}
