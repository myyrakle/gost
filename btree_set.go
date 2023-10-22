package gost

type BTreeSet[K Ord[K]] struct {
	_treemap *BTreeMap[K, struct{}]
}

// Creates an empty BTreeSet.
func BTreeSetNew[K Ord[K]]() BTreeSet[K] {
	return BTreeSet[K]{}
}

// Clears the set, removing all elements.
func (self *BTreeSet[K]) Clear() {
	self._treemap.Clear()
}

// Returns true if the set contains an element equal to the value.
// The value may be any borrowed form of the set’s element type, but the ordering on the borrowed form must match the ordering on the element type.
func (self *BTreeSet[K]) Contains(key K) Bool {
	return self._treemap.ContainsKey(key)
}

// Adds a value to the set.
// Returns whether the value was newly inserted. That is:
// If the set did not previously contain an equal value, true is returned.
// If the set already contained an equal value, false is returned, and the entry is not updated.
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
func (self *BTreeSet[K]) Remove(key K) Bool {
	result := self._treemap.Remove(key)

	if result.IsSome() {
		return true
	} else {
		return false
	}
}

// Returns true if the set contains no elements.
func (self *BTreeSet[K]) IsEmpty() Bool {
	return self._treemap.IsEmpty()
}

// Returns the number of elements in the set.
func (self *BTreeSet[K]) Len() USize {
	return self._treemap.Len()
}

type BTreeSetIter[K Ord[K]] struct {
	vec      Vec[K]
	position USize
}

// into_iter
func (self *BTreeSet[K]) IntoIter() BTreeSetIter[K] {
	keys := self._treemap.root._ToKeyVec()

	return &HashSetIter[K]{vec: keys, position: 0}
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
