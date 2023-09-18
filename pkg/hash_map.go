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

// Removes a key from the map, returning the value at the key if the key was previously in the map.
func (self *HashMap[K, V]) Remove(key K) Option[V] {
	old, ok := self.data[key]
	delete(self.data, key)
	if ok {
		return Some(old)
	} else {
		return None[V]()
	}
}

// Clears the map, removing all key-value pairs. Keeps the allocated memory for reuse.
func (self *HashMap[K, V]) Clear() {
	self.data = map[K]V{}
}

// Returns a reference to the value corresponding to the key.
func (self HashMap[K, V]) Get(key K) Option[V] {
	value, ok := self.data[key]
	if ok {
		return Some(value)
	} else {
		return None[V]()
	}
}

// Returns true if the map contains a value for the specified key.
func (self HashMap[K, V]) ContainsKey(key K) Bool {
	_, ok := self.data[key]
	return Bool(ok)
}

type MapIter[K comparable, V any] struct {
	vec      Vec[Pair[K, V]]
	position Int
}

// into_iter
func (self HashMap[K, V]) IntoIter() Iterator[Pair[K, V]] {
	vec := Vec[Pair[K, V]]{}
	for key, value := range self.data {
		vec.Push(Pair[K, V]{Key: key, Value: value})
	}

	return &MapIter[K, V]{vec: vec, position: 0}
}

// next
func (self *MapIter[K, V]) Next() Option[Pair[K, V]] {
	if self.position >= self.vec.Len() {
		return None[Pair[K, V]]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[Pair[K, V]](value)
}

// map
func (self MapIter[K, V]) Map(f func(Pair[K, V]) Pair[K, V]) Iterator[Pair[K, V]] {
	newVec := VecNew[Pair[K, V]]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.Push(f(value.Unwrap()))
	}
}

// filter
func (self MapIter[K, V]) Filter(f func(Pair[K, V]) Bool) Iterator[Pair[K, V]] {
	newVec := VecNew[Pair[K, V]]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}

		unwraped := value.Unwrap()
		if f(unwraped) {
			newVec.Push(unwraped)
		}
	}
}
