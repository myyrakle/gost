package gost

type BTreeSet[K Ord[K]] struct {
	_treemap *BTreeMap[K, struct{}]
}

// Creates an empty BTreeSet.
func BTreeSetNew[K Ord[K]]() BTreeSet[K] {
	return BTreeSet[K]{}
}

// Clears the set, removing all elements.
func (set *BTreeSet[K]) Clear() {
	set._treemap.Clear()
}

// Returns true if the set contains an element equal to the value.
// The value may be any borrowed form of the set’s element type, but the ordering on the borrowed form must match the ordering on the element type.
func (set *BTreeSet[K]) Contains(key K) Bool {
	return set._treemap.ContainsKey(key)
}
