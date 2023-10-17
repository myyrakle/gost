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
	// Must be zero if `_Type` is `Leaf` and non-zero if `Type` is `Internal`.
	height uint

	_Type _NodeType

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

// Inserts a key-value pair into the map.
// If the map did not have this key present, None is returned.
// If the map did have this key present, the value is updated, and the old value is returned. The key is not updated, though; this matters for types that can be == without being identical. See the module-level documentation for more.
func (m *BTreeMap[K, V]) Insert(key K, value V) Option[V] {
	if m.root == nil {
		m.root = &BTreeNodeRef[K, V]{
			_Type: _LEAF_OF_INTERNAL,
			node: BTreeNode[K, V]{
				keys:   make([]*K, _BTREE_CAPACITY),
				values: make([]*V, _BTREE_CAPACITY),
			},
		}
	}
	return m.root.insert(key, value)
}
