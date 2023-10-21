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
