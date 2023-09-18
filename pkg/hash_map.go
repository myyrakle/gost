package gost

type HashMap[K comparable, V any] struct {
	data map[K]V
}

// Creates an empty HashMap.
func HashMapNew[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{data: map[K]V{}}
}

// Creates an empty HashMap with at least the specified capacity.
func HashMapWithCapacity[K comparable, V any](capacity Int) HashMap[K, V] {
	return HashMap[K, V]{data: make(map[K]V, capacity)}
}

// As Raw Map
func (m HashMap[K, V]) AsMap() map[K]V {
	return m.data
}

// Returns the number of elements in the map.
func (m HashMap[K, V]) Len() Int {
	return Int(len(m.data))
}

// Returns true if the map contains no elements.
func (m HashMap[K, V]) IsEmpty() Bool {
	return m.Len() == 0
}
