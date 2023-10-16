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

// Returns the number of elements in the set.
func (self HashSet[K]) Len() ISize {
	return self.hashMap.Len()
}
