package gost

const _B = 6
const _BTREE_CAPACITY = _B*2 - 1

type _NodeType int8

const _LEAF = _NodeType(0)
const _INTERNAL = _NodeType(1)

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

// Returns value corresponding to the key.
func (self *BTreeMap[K, V]) Get(key K) Option[V] {
	if self.root == nil {
		return None[V]()
	}

	result, index := self.root._Search(key)
	if result.IsNone() {
		return None[V]()
	}

	return Some(result.Unwrap().values.data[index])
}

// Removes a key from the map, returning the value at the key if the key was previously in the map.
// The key may be any borrowed form of the map’s key type, but the ordering on the borrowed form must match the ordering on the key type.
func (self *BTreeMap[K, V]) Remove(key K) Option[V] {
	if self.root == nil {
		return None[V]()
	}

	result, index := self.root._Search(key)

	// Call the remove function for root
	self.root._Remove(key)

	// If the root node has 0 keys, make its first child as the new root
	//  if it has a child, otherwise set root as NULL
	if self.root.n == 0 {
		if self.root._Type == _LEAF {
			self.root = nil
		} else {
			self.root = self.root.childs.GetUnchecked(0)
		}
	}

	if result.IsNone() {
		return None[V]()
	} else {
		self.len--
		return Some(result.Unwrap().values.data[index])
	}
}

// Debug only
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

// To Vec[Pair[K, V]]
func (self BTreeNode[K, V]) _ToVec() Vec[Pair[K, V]] {
	vec := Vec[Pair[K, V]]{}

	// There are n keys and n+1 children, traverse through n keys
	// and first n children

	i := USize(0)
	for i < self.keys.Len() {
		// If this is not leaf, then before printing key[i],
		// traverse the subtree rooted with child C[i].
		if self._Type != _LEAF {
			if self.childs.GetUnchecked(i) != nil {
				childsVec := self.childs.GetUnchecked(i)._ToVec()
				vec.Append(&childsVec)
			}
		}
		vec.Push(Pair[K, V]{Key: self.keys.GetUnchecked(i), Value: self.values.GetUnchecked(i)})
		i++
	}

	return vec
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

// A utility function that returns the index of the first key that is
// greater than or equal to k
func (self BTreeNode[K, V]) _FindKey(key K) int {
	i := 0
	for i < self.n && self.keys.GetUnchecked(USize(i)).Cmp(key) == OrderingLess {
		i++
	}
	return i
}

// A function to remove the key k from the sub-tree rooted with this node
func (self *BTreeNode[K, V]) _Remove(key K) {
	i := self._FindKey(key)

	// The key to be removed is present in this node
	if i < self.n && self.keys.GetUnchecked(USize(i)).Cmp(key) == OrderingEqual {
		if self._Type == _LEAF {
			// If the key is in a leaf node - removeFromLeaf is called
			self._RemoveFromLeaf(i)
		} else {
			// If the key is in a non-leaf node - removeFromNonLeaf is called
			self._RemoveFromNonLeaf(i)
		}
	} else {
		// If this node is a leaf node, then the key is not present in tree
		if self._Type == _LEAF {
			return
		}

		// The key to be removed is present in the sub-tree rooted with this node
		// The flag indicates whether the key is present in the sub-tree rooted
		// with the last child of this node
		flag := Bool(false)
		if i == self.n {
			flag = true
		}

		// If the child where the key is supposed to exist has less that t keys,
		// we fill that child
		if self.childs.GetUnchecked(USize(i)).n < self._MinimumDegree {
			self._Fill(i)
		}

		// If the last child has been merged, it must have merged with the previous
		// child and so we recurse on the (i-1)th child. Else, we recurse on the
		// (i)th child which now has atleast t keys
		if flag && i > self.n {
			self.childs.GetUnchecked(USize(i - 1))._Remove(key)
		} else {
			self.childs.GetUnchecked(USize(i))._Remove(key)
		}
	}
}

// A function to remove the idx-th key from this node - which is a leaf node
func (self *BTreeNode[K, V]) _RemoveFromLeaf(idx int) {
	// Move all the keys after the idx-th pos one place backward
	for i := idx + 1; i < self.n; i++ {
		self.keys.SetUnchecked(USize(i-1), self.keys.GetUnchecked(USize(i)))
		self.values.SetUnchecked(USize(i-1), self.values.GetUnchecked(USize(i)))
	}

	// Reduce the count of keys
	self.n--
}

// A function to remove the idx-th key from this node - which is a non-leaf node
func (self *BTreeNode[K, V]) _RemoveFromNonLeaf(idx int) {
	key := self.keys.GetUnchecked(USize(idx))

	// If the child that precedes k (C[idx]) has atleast t keys,
	// find the predecessor 'pred' of k in the subtree rooted at
	// C[idx]. Replace k by pred. Recursively delete pred
	// in C[idx]
	if self.childs.GetUnchecked(USize(idx)).n >= self._MinimumDegree {
		pred := self._GetPred(idx)
		self.keys.SetUnchecked(USize(idx), pred.keys.GetUnchecked(USize(pred.n-1)))
		self.values.SetUnchecked(USize(idx), pred.values.GetUnchecked(USize(pred.n-1)))
		child := self.childs.GetUnchecked(USize(idx))
		child._Remove(pred.keys.GetUnchecked(USize(pred.n - 1)))
	} else if self.childs.GetUnchecked(USize(idx+1)).n >= self._MinimumDegree {
		// If the child C[idx] has less that t keys, examine C[idx+1].
		// If C[idx+1] has atleast t keys, find the successor 'succ' of k in
		// the subtree rooted at C[idx+1]
		// Replace k by succ
		// Recursively delete succ in C[idx+1]
		succ := self._GetSucc(idx)
		self.keys.SetUnchecked(USize(idx), succ.keys.GetUnchecked(0))
		self.values.SetUnchecked(USize(idx), succ.values.GetUnchecked(0))
		self.childs.GetUnchecked(USize(idx + 1))._Remove(succ.keys.GetUnchecked(0))
	} else {
		// If both C[idx] and C[idx+1] has less that t keys,merge k and all of C[idx+1]
		// into C[idx]
		// Now C[idx] contains 2t-1 keys
		// Free C[idx+1] and recursively delete k from C[idx]
		self._Merge(idx)
		self.childs.GetUnchecked(USize(idx))._Remove(key)
	}
}

// A function to get predecessor of keys[idx]
func (self BTreeNode[K, V]) _GetPred(idx int) *BTreeNode[K, V] {
	// Keep moving to the right most node until we reach a leaf
	curr := self.childs.GetUnchecked(USize(idx))
	for curr._Type != _LEAF {
		curr = curr.childs.GetUnchecked(USize(curr.n))
	}

	// Return the last key of the leaf
	return curr
}

// A function to get successor of keys[idx]
func (self BTreeNode[K, V]) _GetSucc(idx int) *BTreeNode[K, V] {
	// Keep moving the left most node starting from C[idx+1] until we reach a leaf
	curr := self.childs.GetUnchecked(USize(idx + 1))
	for curr._Type != _LEAF {
		curr = curr.childs.GetUnchecked(0)
	}

	// Return the first key of the leaf
	return curr
}

// A function to fill child C[idx] which has less than t-1 keys
func (self *BTreeNode[K, V]) _Fill(idx int) {
	// If the previous child(C[idx-1]) has more than t-1 keys, borrow a key
	// from that child
	if idx != 0 && self.childs.GetUnchecked(USize(idx-1)).n >= self._MinimumDegree {
		self._BorrowFromPrev(idx)
	} else if idx != self.n && self.childs.GetUnchecked(USize(idx+1)).n >= self._MinimumDegree {
		/* If the next child(C[idx+1]) has more than t-1 keys, borrow a key from that child */
		self._BorrowFromNext(idx)
	} else {
		// Merge C[idx] with its sibling
		// If C[idx] is the last child, merge it with with its previous sibling
		// Otherwise merge it with its next sibling
		if idx != self.n {
			self._Merge(idx)
		} else {
			self._Merge(idx - 1)
		}
	}
}

// A function to borrow a key from C[idx-1] and insert it
// into C[idx]
func (self *BTreeNode[K, V]) _BorrowFromPrev(idx int) {
	child := self.childs.GetUnchecked(USize(idx))
	sibling := self.childs.GetUnchecked(USize(idx - 1))

	// The last key from C[idx-1] goes up to the parent and key[idx-1]
	// from parent is inserted as the first key in C[idx]. Thus, the  loses
	// sibling one key and child gains one key

	// Moving all key in C[idx] one step ahead
	for i := child.n - 1; i >= 0; i-- {
		child.keys.SetUnchecked(USize(i+1), child.keys.GetUnchecked(USize(i)))
		child.values.SetUnchecked(USize(i+1), child.values.GetUnchecked(USize(i)))
	}

	// If C[idx] is not a leaf, move all its child pointers one step ahead
	if child._Type != _LEAF {
		for i := child.n; i >= 0; i-- {
			child.childs.SetUnchecked(USize(i+1), child.childs.GetUnchecked(USize(i)))
		}
	}

	// Setting child's first key equal to keys[idx-1] from the current node
	child.keys.SetUnchecked(0, self.keys.GetUnchecked(USize(idx-1)))
	child.values.SetUnchecked(0, self.values.GetUnchecked(USize(idx-1)))

	// Moving sibling's last child as C[idx]'s first child
	if child._Type != _LEAF {
		child.childs.SetUnchecked(0, sibling.childs.GetUnchecked(USize(sibling.n)))
	}

	// Moving the key from the sibling to the parent
	// This reduces the number of keys in the sibling
	self.keys.SetUnchecked(USize(idx-1), sibling.keys.GetUnchecked(USize(sibling.n-1)))
	self.values.SetUnchecked(USize(idx-1), sibling.values.GetUnchecked(USize(sibling.n-1)))

	child.n++
	sibling.n--
}

// A function to borrow a key from the C[idx+1] and place
// it in C[idx]
func (self *BTreeNode[K, V]) _BorrowFromNext(idx int) {
	child := self.childs.GetUnchecked(USize(idx))
	sibling := self.childs.GetUnchecked(USize(idx + 1))

	// keys[idx] is inserted as the last key in C[idx]
	child.keys.SetUnchecked(USize(child.n), self.keys.GetUnchecked(USize(idx)))
	child.values.SetUnchecked(USize(child.n), self.values.GetUnchecked(USize(idx)))

	// Sibling's first child is inserted as the last child
	// into C[idx]
	if child._Type != _LEAF {
		child.childs.SetUnchecked(USize(child.n+1), sibling.childs.GetUnchecked(0))
	}

	//The first key from sibling is inserted into keys[idx]
	self.keys.SetUnchecked(USize(idx), sibling.keys.GetUnchecked(0))
	self.values.SetUnchecked(USize(idx), sibling.values.GetUnchecked(0))

	// Moving all keys in sibling one step behind
	for i := 1; i < sibling.n; i++ {
		sibling.keys.SetUnchecked(USize(i-1), sibling.keys.GetUnchecked(USize(i)))
		sibling.values.SetUnchecked(USize(i-1), sibling.values.GetUnchecked(USize(i)))
	}

	// Moving the child pointers one step behind
	if sibling._Type != _LEAF {
		for i := 1; i <= sibling.n; i++ {
			sibling.childs.SetUnchecked(USize(i-1), sibling.childs.GetUnchecked(USize(i)))
		}
	}

	// Increasing and decreasing the key count of C[idx] and C[idx+1]
	// respectively
	child.n++
	sibling.n--
}

// A function to merge C[idx] with C[idx+1]
// C[idx+1] is freed after merging
func (self *BTreeNode[K, V]) _Merge(idx int) {
	child := self.childs.GetUnchecked(USize(idx))
	sibling := self.childs.GetUnchecked(USize(idx + 1))

	// Pulling a key from the current node and inserting it into (t-1)th
	// position of C[idx]
	child.keys.SetUnchecked(USize(self._MinimumDegree-1), self.keys.GetUnchecked(USize(idx)))
	child.values.SetUnchecked(USize(self._MinimumDegree-1), self.values.GetUnchecked(USize(idx)))

	// Copying the keys from C[idx+1] to C[idx] at the end
	for i := 0; i < sibling.n; i++ {
		child.keys.SetUnchecked(USize(i+self._MinimumDegree), sibling.keys.GetUnchecked(USize(i)))
		child.values.SetUnchecked(USize(i+self._MinimumDegree), sibling.values.GetUnchecked(USize(i)))
	}

	// Copying the child pointers from C[idx+1] to C[idx]
	if child._Type != _LEAF {
		for i := 0; i <= sibling.n; i++ {
			child.childs.SetUnchecked(USize(i+self._MinimumDegree), sibling.childs.GetUnchecked(USize(i)))
		}
	}

	// Moving all keys after idx in the current node one step before -
	// to fill the gap created by moving keys[idx] to C[idx]
	for i := idx + 1; i < self.n; i++ {
		self.keys.SetUnchecked(USize(i-1), self.keys.GetUnchecked(USize(i)))
		self.values.SetUnchecked(USize(i-1), self.values.GetUnchecked(USize(i)))
	}

	// Moving the child pointers after (idx+1) in the current node one
	// step before
	for i := idx + 2; i <= self.n; i++ {
		self.childs.SetUnchecked(USize(i-1), self.childs.GetUnchecked(USize(i)))
	}

	// Updating the key count of child and the current node
	child.n += sibling.n + 1
	self.n--

	// Freeing the memory occupied by sibling
	sibling = nil
}

type BTreeMapIter[K Ord[K], V any] struct {
	vec      Vec[Pair[K, V]]
	position USize
}

// into_iter
func (self BTreeMap[K, V]) IntoIter() Iterator[K, V] {
	vec := Vec[Pair[K, V]]{}

	if self.root != nil {
		vec = self.root._ToVec()
	}

	return BTreeMapIter[K, V]{vec: vec, position: 0}
}

// next
func (self *BTreeMapIter[K, V]) Next() Option[Pair[K, V]] {
	if self.position >= self.vec.Len() {
		return None[Pair[K, V]]()
	}

	result := self.vec.GetUnchecked(self.position)
	self.position++
	return Some(result)
}
