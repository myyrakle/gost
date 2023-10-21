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
