package gost

import (
	"fmt"
	"strings"
)

// A ordered set based on a B-Tree.
type BTreeSet[K Ord[K]] struct {
	_treemap BTreeMap[K, struct{}]
}

// Creates an empty BTreeSet.
func BTreeSetNew[K Ord[K]]() BTreeSet[K] {
	return BTreeSet[K]{}
}

// Clears the set, removing all elements.
//
//	set := gost.BTreeSetNew[Int]()
//	set.Insert(gost.I32(1))
//	set.Insert(gost.I32(2))
//	set.Clear()
//	gost.AssertEq(set.Len(), gost.USize(0))
func (self *BTreeSet[K]) Clear() {
	self._treemap.Clear()
}

// Returns true if the set contains an element equal to the value.
// The value may be any borrowed form of the set’s element type, but the ordering on the borrowed form must match the ordering on the element type.
//
//	set := gost.BTreeSetNew[Int]()
//	set.Insert(gost.I32(1))
//	set.Insert(gost.I32(2))
//	gost.Assert(set.Contains(gost.I32(1)))
//	gost.Assert(!set.Contains(gost.I32(3)))
func (self *BTreeSet[K]) Contains(key K) Bool {
	return self._treemap.ContainsKey(key)
}

// Adds a value to the set.
// Returns whether the value was newly inserted. That is:
// If the set did not previously contain an equal value, true is returned.
// If the set already contained an equal value, false is returned, and the entry is not updated.
//
//	set := gost.BTreeSetNew[Int]()
//	gost.Assert(set.Insert(gost.I32(1)))
//	gost.Assert(!set.Insert(gost.I32(1)))
func (self *BTreeSet[K]) Insert(key K) Bool {
	result := self._treemap.Insert(key, struct{}{})

	if result.IsSome() {
		return false
	} else {
		return true
	}
}

// If the set contains an element equal to the value, removes it from the set and drops it. Returns whether such an element was present.
// The value may be any borrowed form of the set’s element type, but the ordering on the borrowed form must match the ordering on the element type.
//
//	set := gost.BTreeSetNew[Int]()
//	set.Insert(gost.I32(1))
//	set.Insert(gost.I32(2))
//	gost.Assert(set.Remove(gost.I32(1)))
//	gost.Assert(!set.Remove(gost.I32(3)))
func (self *BTreeSet[K]) Remove(key K) Bool {
	result := self._treemap.Remove(key)

	if result.IsSome() {
		return true
	} else {
		return false
	}
}

// Returns true if the set contains no elements.
//
//	set := gost.BTreeSetNew[Int]()
//	gost.Assert(set.IsEmpty())
//
//	set.Insert(gost.I32(1))
//	gost.Assert(!set.IsEmpty())
func (self BTreeSet[K]) IsEmpty() Bool {
	return self._treemap.IsEmpty()
}

// Returns the number of elements in the set.
//
//	set := gost.BTreeSetNew[Int]()
//	gost.AssertEq(set.Len(), gost.USize(0))
//
//	set.Insert(gost.I32(1))
//	gost.AssertEq(set.Len(), gost.USize(1))
func (self BTreeSet[K]) Len() USize {
	return self._treemap.Len()
}

// Returns true if the set is a subset of another, i.e., other contains at least all the elements in self.
//
//	set1 := gost.BTreeSetNew[I32]()
//	set1.Insert(gost.I32(1))
//	set1.Insert(gost.I32(2))
//
//	set2 := gost.BTreeSetNew[I32]()
//	set2.Insert(gost.I32(1))
//	set2.Insert(gost.I32(2))
//	set2.Insert(gost.I32(3))
//
//	gost.Assert(set1.IsSubset(set2))
//	gost.Assert(!set2.IsSubset(set1))
func (self BTreeSet[K]) IsSubset(other BTreeSet[K]) Bool {
	if self.Len() > other.Len() {
		return false
	}

	iter := self.IntoIter()
	for {
		value := iter.Next()

		if value.IsNone() {
			return true
		}

		if !other.Contains(value.Unwrap()) {
			return false
		}
	}
}

// Returns true if the set is a superset of another, i.e., self contains at least all the elements in other.
//
//	set1 := gost.BTreeSetNew[I32]()
//	set1.Insert(gost.I32(1))
//	set1.Insert(gost.I32(2))
//
//	set2 := gost.BTreeSetNew[I32]()
//	set2.Insert(gost.I32(1))
//	set2.Insert(gost.I32(2))
//	set2.Insert(gost.I32(3))
//
//	gost.Assert(!set1.IsSuperset(set2))
//	gost.Assert(set2.IsSuperset(set1))
func (self BTreeSet[K]) IsSuperset(other BTreeSet[K]) Bool {
	return other.IsSubset(self)
}

// Visits the elements representing the intersection, i.e., the elements that are both in self and other, in ascending order.
//
//	set1 := gost.BTreeSetNew[I32]()
//	set1.Insert(gost.I32(1))
//	set1.Insert(gost.I32(2))
//	set1.Insert(gost.I32(5))
//
//	set2 := gost.BTreeSetNew[I32]()
//	set2.Insert(gost.I32(1))
//	set2.Insert(gost.I32(2))
//	set2.Insert(gost.I32(3))
//
//	intersection := set1.Intersection(set2)
//	gost.Assert(intersection.Len() == gost.USize(2))
//	gost.Assert(intersection.Contains(gost.I32(1)))
//	gost.Assert(intersection.Contains(gost.I32(2)))
func (self BTreeSet[K]) Intersection(other BTreeSet[K]) BTreeSet[K] {
	newSet := BTreeSetNew[K]()

	iter := self.IntoIter()
	for {
		value := iter.Next()

		if value.IsNone() {
			return newSet
		}

		if other.Contains(value.Unwrap()) {
			newSet.Insert(value.Unwrap())
		}
	}
}

// Returns an iterator over the set.
type BTreeSetIter[K Ord[K]] struct {
	vec      Vec[K]
	position USize
}

// into_iter
func (self *BTreeSet[K]) IntoIter() Iterator[K] {
	keys := self._treemap.root._ToKeyVec()

	return &BTreeSetIter[K]{vec: keys, position: 0}
}

// next
func (self *BTreeSetIter[K]) Next() Option[K] {
	if self.position >= self.vec.Len() {
		return None[K]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[K](value)
}

// map
func (self BTreeSetIter[K]) Map(f func(K) K) Iterator[K] {
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
func (self BTreeSetIter[K]) Filter(f func(K) Bool) Iterator[K] {
	newVec := VecNew[K]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}

		if f(value.Unwrap()) {
			newVec.Push(value.Unwrap())
		}
	}
}

// fold
func (self BTreeSetIter[K]) Fold(init K, f func(K, K) K) K {
	for {
		value := self.Next()

		if value.IsNone() {
			return init
		}

		init = f(init, value.Unwrap())
	}
}

// rev
func (self BTreeSetIter[K]) Rev() Iterator[K] {
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
func (self BTreeSetIter[K]) CollectToVec() Vec[K] {
	return self.vec
}

// Collect to LinkedList
func (self BTreeSetIter[K]) CollectToLinkedList() LinkedList[K] {
	list := LinkedListNew[K]()

	for {
		value := self.Next()

		if value.IsNone() {
			return list
		}
		list.PushBack(value.Unwrap())
	}
}

// impl Display for BTreeSet
func (self BTreeSet[K]) Display() String {
	keys := self.IntoIter().CollectToVec()

	buffer := String("")
	buffer += "BTreeSet{"

	fields := []string{}

	for i := USize(0); i < keys.Len(); i++ {
		key := keys.GetUnchecked(i)

		fields = append(fields, string(Format("{}", key)))
	}

	buffer += String(strings.Join(fields, ", "))

	buffer += "}"

	return buffer
}

// impl Debug for BTreeSet
func (self BTreeSet[K]) Debug() String {
	return self.Display()
}

// impl AsRef for BTreeSet
func (self BTreeSet[K]) AsRef() *BTreeSet[K] {
	return &self
}

// impl Clone for BTreeSet
func (self BTreeSet[K]) Clone() BTreeSet[K] {
	newSet := BTreeSetNew[K]()

	for {
		value := self.IntoIter().Next()

		if value.IsNone() {
			return newSet
		}

		e := value.Unwrap()

		clone := castToClone[K](e)

		if clone.IsNone() {
			typeName := getTypeName(e)
			panic(fmt.Sprintf("'%s' does not implement Clone[%s]", typeName, typeName))
		} else {
			newSet.Insert(clone.Unwrap().Clone())
		}
	}
}

// impl Eq for BTreeSet
func (self BTreeSet[K]) Eq(other BTreeSet[K]) Bool {
	if self.Len() != other.Len() {
		return false
	}

	iter := self.IntoIter()
	for {
		value := iter.Next()

		if value.IsNone() {
			return true
		}

		if !other.Contains(value.Unwrap()) {
			return false
		}
	}
}
