package gost

const _B = 6
const _BTREE_CAPACITY = _B*2 - 1

type _NodeType int8

const _LEAF = _NodeType(0)
const _INTERNAL = _NodeType(1)
const _LEAF_OF_INTERNAL = _NodeType(2)

type BTreeMap[K Ord[K], V any] struct {
	root *BTreeNodeRef[K, V]
	len  uint
}

type BTreeNodeRef[K Ord[K], V any] struct {
	// The number of levels that the node and the level of leaves are apart, a
	// constant of the node that cannot be entirely described by `Type`, and that
	// the node itself does not store. We only need to store the height of the root
	// node, and derive every other node's height from it.
	// Must be zero if `Type` is `Leaf` and non-zero if `Type` is `Internal`.
	height uint

	nodeType _NodeType

	node BTreeNode[K, V]
}

type BTreeNode[K Ord[K], V any] struct {
	keys   []*K
	values []*V

	// The number of keys and values this node stores.
	len uint16
}

// Creates an empty BTreeMap.
func BTreeMapNew[K Ord[K], V any]() BTreeMap[K, V] {
	return BTreeMap[K, V]{}
}
