package gost

type HashSet[K comparable] struct {
	hashMap HashMap[K, struct{}]
}

// Creates an empty HashSet.
func HashSetNew[K comparable]() HashSet[K] {
	return HashSet[K]{hashMap: HashMapNew[K, struct{}]()}
}
