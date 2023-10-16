package gost

type HashSet[K comparable] struct {
	hashMap HashMap[K, struct{}]
}

// Creates an empty HashSet.
func HashSetNew[K comparable]() HashSet[K] {
	return HashSet[K]{hashMap: HashMapNew[K, struct{}]()}
}

// Creates an empty HashSet with at least the specified capacity.
func HashSetWithCapacity[K comparable](capacity ISize) HashSet[K] {
	return HashSet[K]{hashMap: HashMapWithCapacity[K, struct{}](capacity)}
}

// As Slice
func (self HashSet[K]) AsSlice() []K {
	result := make([]K, 0, self.hashMap.Len())
	for key, _ := range self.hashMap.data {
		result = append(result, key)
	}

	return result
}

// From Slice
func HashSetFromSlice[K comparable](slice []K) HashSet[K] {
	result := HashSetNew[K]()
	for _, value := range slice {
		result.Insert(value)
	}

	return result
}

// Returns the number of elements in the set.
func (self HashSet[K]) Len() ISize {
	return self.hashMap.Len()
}

// Returns true if the set contains no elements.
func (self HashSet[K]) IsEmpty() Bool {
	return self.hashMap.IsEmpty()
}

// Adds a value to the set.
// Returns whether the value was newly inserted. That is:
// If the set did not previously contain this value, true is returned.
// If the set already contained this value, false is returned, and the set is not modified: original value is not replaced, and the value passed as argument is dropped.
func (self *HashSet[K]) Insert(value K) Bool {
	result := self.hashMap.Insert(value, struct{}{})
	return result.IsNone()
}

// Removes a value from the set. Returns whether the value was present in the set.
// The value may be any borrowed form of the set’s value type, but Hash and Eq on the borrowed form must match those for the value type.
func (self *HashSet[K]) Remove(value K) Bool {
	result := self.hashMap.Remove(value)
	return result.IsSome()
}
