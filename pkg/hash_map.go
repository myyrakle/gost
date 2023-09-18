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
func (self HashMap[K, V]) AsMap() map[K]V {
	return self.data
}

// Returns the number of elements in the map.
func (self HashMap[K, V]) Len() Int {
	return Int(len(self.data))
}

// Returns true if the map contains no elements.
func (self HashMap[K, V]) IsEmpty() Bool {
	return self.Len() == 0
}

// Inserts a key-value pair into the map.
// If the map did not have this key present, None is returned.
func (self *HashMap[K, V]) Insert(key K, value V) Option[V] {
	old, ok := self.data[key]
	self.data[key] = value
	if ok {
		return Some(old)
	} else {
		return None[V]()
	}
}
