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
