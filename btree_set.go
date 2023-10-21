package gost

type BTreeSet[K Ord[K]] struct {
	_treemap *BTreeNode[K, struct{}]
}

// Creates an empty BTreeSet.
func BTreeSetNew[K Ord[K]]() BTreeSet[K] {
	return BTreeSet[K]{}
}
