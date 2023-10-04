package gost

type BTreeMap[K Ord[K], V any] struct {
	root *BTreeNode[K, V]
}

type BTreeNode[K Ord[K], V any] struct {
	key   K
	value V
	left  *BTreeNode[K, V]
	right *BTreeNode[K, V]
}

// Creates an empty BTreeMap.
func BTreeMapNew[K Ord[K], V any]() BTreeMap[K, V] {
	return BTreeMap[K, V]{}
}
