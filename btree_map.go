package gost

const _B = 6
const _BTREE_CAPACITY = _B*2 - 1

type _NodeType int8

const _LEAF = _NodeType(0)
const _INTERNAL = _NodeType(1)
const _ROOT = _NodeType(2)

type BTreeMap[K Ord[K], V any] struct {
	root *BTreeNode[K, V]
	len  uint
}

type BTreeNode[K Ord[K], V any] struct {
	_Type          _NodeType
	_MinimumDegree int // Minimum degree (defines the range for number of keys)

	keys   Vec[K]
	values Vec[V]

	n int // Current number of keys

	childs Vec[*BTreeNode[K, V]]
}

// Creates an empty BTreeMap.
func BTreeMapNew[K Ord[K], V any]() BTreeMap[K, V] {
	return BTreeMap[K, V]{}
}

// Inserts a key-value pair into the map.
// If the map did not have this key present, None is returned.
// If the map did have this key present, the value is updated, and the old value is returned. The key is not updated, though; this matters for types that can be == without being identical. See the module-level documentation for more.
func (self *BTreeMap[K, V]) Insert(key K, value V) Option[V] {
	// If tree is empty
	if self.root == nil {
		// Allocate memory for root
		self.root = &BTreeNode[K, V]{
			_Type:  _LEAF,
			keys:   VecWithLen[K](_BTREE_CAPACITY),
			values: VecWithLen[V](_BTREE_CAPACITY),
			childs: VecWithLen[*BTreeNode[K, V]](_BTREE_CAPACITY + 1),
			n:      1,
		}

		self.root.keys.SetUnchecked(0, key)
		self.root.values.SetUnchecked(0, value)

		self.len = 1
		return None[V]()
	} else /* If tree is not empty */ {
		// If exists, update value
		// TODO: implement update

		// If root is full, then tree grows in height
		if self.root.n == _BTREE_CAPACITY {
			// Allocate memory for new root
			newRoot := &BTreeNode[K, V]{
				_Type:          _INTERNAL,
				_MinimumDegree: _BTREE_CAPACITY,
				keys:           VecWithLen[K](_BTREE_CAPACITY),
				values:         VecWithLen[V](_BTREE_CAPACITY),
				childs:         VecWithLen[*BTreeNode[K, V]](_BTREE_CAPACITY + 1),
			}

			// Make old root as child of new root
			newRoot.childs.SetUnchecked(0, self.root)

			// Split the old root and move 1 key to the new root
			newRoot.splitChild(0, self.root)

			// New root has two children now.  Decide which of the
			// two children is going to have new key
			i := USize(0)
			if newRoot.keys.GetUnchecked(0).Cmp(key) == OrderingLess {
				i++
			}
			newRoot.childs.GetUnchecked(i)._InsertNonFull(key, value)

			// Change root
			self.root = newRoot
			self.len++
			return None[V]()
		} else /* If root is not full, call insertNonFull for root */ {
			self.root._InsertNonFull(key, value)
			self.len++
			return None[V]()
		}
	}
}

// Returns true if the map contains a value for the specified key.
// The key may be any borrowed form of the map’s key type, but the ordering on the borrowed form must match the ordering on the key type.
func (self *BTreeMap[K, V]) ContainsKey(key K) Bool {
	if self.root == nil {
		return false
	}

	result, _ := self.root._Search(key)
	return Bool(result.IsSome())
}

// Returns the number of elements in the map.
func (self *BTreeMap[K, V]) Len() USize {
	return USize(self.len)
}

// Returns true if the map contains no elements.
func (self *BTreeMap[K, V]) IsEmpty() Bool {
	return self.len == 0
}

// Clears the map, removing all elements.
func (self *BTreeMap[K, V]) Clear() {
	self.root = nil
	self.len = 0
}

// The key may be any borrowed form of the map’s key type, but the ordering on the borrowed form must match the ordering on the key type.
func (self *BTreeMap[K, V]) Get(key K) Option[*V] {
	if self.root == nil {
		return None[*V]()
	}

	result, index := self.root._Search(key)
	if result.IsNone() {
		return None[*V]()
	}

	return Some(&result.Unwrap().values.data[index])
}

func (self *BTreeMap[K, V]) Test() {
	self.root._Traverse()
}

// Function to traverse all nodes in a subtree rooted with this node
func (self BTreeNode[K, V]) _Traverse() {
	// There are n keys and n+1 children, traverse through n keys
	// and first n children

	i := USize(0)
	for i < self.keys.Len() {
		// If this is not leaf, then before printing key[i],
		// traverse the subtree rooted with child C[i].
		if self._Type != _LEAF {
			if self.childs.GetUnchecked(i) != nil {
				self.childs.GetUnchecked(i)._Traverse()
			}
		}
		Println("{}, {}", self.keys.GetUnchecked(i), self.values.GetUnchecked(i))
		i++
	}
}

// Function to search key k in subtree rooted with this node
func (self BTreeNode[K, V]) _Search(key K) (Option[*BTreeNode[K, V]], uint) {
	// Find the first key greater than or equal to k
	i := USize(0)
	for i < self.keys.Len() && key.Cmp(self.keys.GetUnchecked(i)) == OrderingGreater {
		i++
	}

	// If the found key is equal to k, return this node
	if i < self.keys.Len() && key.Cmp(self.keys.GetUnchecked(i)) == OrderingEqual {
		return Some[*BTreeNode[K, V]](&self), uint(i)
	}

	// If key is not found here and this is a leaf node
	if self._Type == _LEAF {
		return None[*BTreeNode[K, V]](), 0
	}

	// Go to the appropriate child
	return self.childs.GetUnchecked(i)._Search(key)
}

// The main function that inserts a new key in this B-Tree
func (self *BTreeMap[K, V]) _Insert(key K, value V) {
	// If tree is empty
	if self.root == nil {
		// Allocate memory for root
		self.root = &BTreeNode[K, V]{
			_Type:  _LEAF,
			keys:   VecWithCapacity[K](_BTREE_CAPACITY),
			values: VecWithCapacity[V](_BTREE_CAPACITY),
		}
		self.root.keys.Push(key)
		self.root.values.Push(value)
		self.len = 1
	} else { // If tree is not empty
		// If root is full, then tree grows in height
		if self.root.keys.Len() == _BTREE_CAPACITY {
			// Allocate memory for new root
			newRoot := &BTreeNode[K, V]{
				_Type:  _INTERNAL,
				keys:   VecWithCapacity[K](_BTREE_CAPACITY),
				values: VecWithCapacity[V](_BTREE_CAPACITY),
			}

			// Make old root as child of new root
			newRoot.childs.Push(self.root)

			// Split the old root and move 1 key to the new root
			newRoot.splitChild(0, self.root)

			// New root has two children now.  Decide which of the
			// two children is going to have new key
			i := USize(0)
			if newRoot.keys.GetUnchecked(0).Cmp(key) == OrderingLess {
				i++
			}
			newRoot.childs.GetUnchecked(i)._InsertNonFull(key, value)

			// Change root
			self.root = newRoot
		} else { // If root is not full, call insertNonFull for root
			self.root._InsertNonFull(key, value)
		}
	}
}

// A utility function to insert a new key in this node
// The assumption is, the node must be non-full when this
// function is called
func (self *BTreeNode[K, V]) _InsertNonFull(key K, value V) {
	// Initialize index as index of rightmost element
	i := self.n - 1

	// If this is a leaf node
	if self._Type == _LEAF {
		// The following loop does two things
		// a) Finds the location of new key to be inserted
		// b) Moves all greater keys to one place ahead
		for i >= 0 && self.keys.GetUnchecked(USize(i)).Cmp(key) == OrderingGreater {
			self.keys.SetUnchecked(USize(i+1), self.keys.GetUnchecked(USize(i)))
			self.values.SetUnchecked(USize(i+1), self.values.GetUnchecked(USize(i)))
			i--
		}

		// Insert the new key at found location
		self.keys.SetUnchecked(USize(i+1), key)
		self.values.SetUnchecked(USize(i+1), value)

		self.n++
	} else { // If this node is not leaf
		// Find the child which is going to have the new key
		for i >= 0 && self.keys.GetUnchecked(USize(i)).Cmp(key) == OrderingGreater {
			i--
		}

		// See if the found child is full
		if self.childs.GetUnchecked(USize(i+1)).keys.Len() == _BTREE_CAPACITY {
			// If the child is full, then split it
			self.splitChild(i+1, self.childs.GetUnchecked(USize(i+1)))

			// After split, the middle key of C[i] goes up and
			// C[i] is splitted into two.  See which of the two
			// is going to have the new key
			if self.keys.GetUnchecked(USize(i+1)).Cmp(key) == OrderingLess {
				i++
			}
		}
		self.childs.GetUnchecked(USize(i+1))._InsertNonFull(key, value)
	}
}

// A utility function to split the child y of this node
// Note that y must be full when this function is called
func (self *BTreeNode[K, V]) splitChild(i int, y *BTreeNode[K, V]) {
	// Create a new node which is going to store (t-1) keys of y
	z := &BTreeNode[K, V]{
		_Type:          y._Type,
		_MinimumDegree: y._MinimumDegree,
		keys:           VecWithLen[K](_BTREE_CAPACITY),
		values:         VecWithLen[V](_BTREE_CAPACITY),
		childs:         VecWithLen[*BTreeNode[K, V]](_BTREE_CAPACITY + 1),
	}
	z.n = self._MinimumDegree - 1

	// Copy the last (t-1) keys of y to z
	for j := 0; j < z._MinimumDegree; j++ {
		z.keys.SetUnchecked(USize(j), y.keys.GetUnchecked(USize(j+self._MinimumDegree)))
		z.values.SetUnchecked(USize(j), y.values.GetUnchecked(USize(j+self._MinimumDegree)))
	}

	// Copy the last t children of y to z
	if y._Type != _LEAF {
		for j := 0; j < self._MinimumDegree; j++ {
			z.childs.SetUnchecked(USize(j), y.childs.GetUnchecked(USize(j+self._MinimumDegree)))
		}
	}

	// Reduce the number of keys in y
	y.n = self._MinimumDegree - 1

	// Since this node is going to have a new child,
	// create space of new child
	for j := self.n; j >= i+1; j-- {
		self.childs.SetUnchecked(USize(j+1), self.childs.GetUnchecked(USize(j)))
	}

	// Link the new child to this node
	self.childs.SetUnchecked(USize(i+1), z)

	// A key of y will move to this node. Find location of
	// new key and move all greater keys one space ahead
	for j := self.n; j >= i; j-- {
		self.keys.SetUnchecked(USize(j+1), self.keys.GetUnchecked(USize(j)))
		self.values.SetUnchecked(USize(j+1), self.values.GetUnchecked(USize(j)))
	}

	// Copy the middle key of y to this node
	self.keys.SetUnchecked(USize(i), y.keys.GetUnchecked(USize(self._MinimumDegree-1)))
	self.values.SetUnchecked(USize(i), y.values.GetUnchecked(USize(self._MinimumDegree-1)))

	// Increment count of keys in this node
	self.n++
}
