package gost

import "strings"

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

type HashMapIter[K comparable, V any] struct {
	vec      Vec[Pair[K, V]]
	position Int
}

// into_iter
func (self HashMap[K, V]) IntoIter() Iterator[Pair[K, V]] {
	vec := Vec[Pair[K, V]]{}
	for key, value := range self.data {
		vec.Push(Pair[K, V]{Key: key, Value: value})
	}

	return &HashMapIter[K, V]{vec: vec, position: 0}
}

// next
func (self *HashMapIter[K, V]) Next() Option[Pair[K, V]] {
	if self.position >= self.vec.Len() {
		return None[Pair[K, V]]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[Pair[K, V]](value)
}

// map
func (self HashMapIter[K, V]) Map(f func(Pair[K, V]) Pair[K, V]) Iterator[Pair[K, V]] {
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
func (self HashMapIter[K, V]) Filter(f func(Pair[K, V]) Bool) Iterator[Pair[K, V]] {
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

// fold
func (self HashMapIter[K, V]) Fold(init Pair[K, V], f func(Pair[K, V], Pair[K, V]) Pair[K, V]) Pair[K, V] {
	for {
		value := self.Next()

		if value.IsNone() {
			return init
		}

		init = f(init, value.Unwrap())
	}
}

// rev
func (self HashMapIter[K, V]) Rev() Iterator[Pair[K, V]] {
	newVec := VecWithLen[Pair[K, V]](self.vec.Len())
	i := self.vec.Len() - 1

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.AsSlice()[i] = value.Unwrap()
		i--
	}
}

func (self HashMapIter[K, V]) CollectToVec() Vec[Pair[K, V]] {
	return self.vec
}

type HashMapKeys[K any] struct {
	vec      Vec[K]
	position Int
}

// An iterator visiting all keys in arbitrary order. The iterator element type is K.
func (self HashMap[K, V]) Keys() Iterator[K] {
	vec := Vec[K]{}
	for key := range self.data {
		vec.Push(key)
	}

	return &HashMapKeys[K]{vec: vec, position: 0}
}

// next
func (self *HashMapKeys[K]) Next() Option[K] {
	if self.position >= self.vec.Len() {
		return None[K]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[K](value)
}

// map
func (self HashMapKeys[K]) Map(f func(K) K) Iterator[K] {
	newVec := VecNew[K]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.Push(f(value.Unwrap()))
	}
}

// filter
func (self HashMapKeys[K]) Filter(f func(K) Bool) Iterator[K] {
	newVec := VecNew[K]()

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

// fold
func (self HashMapKeys[K]) Fold(init K, f func(K, K) K) K {
	for {
		value := self.Next()

		if value.IsNone() {
			return init
		}

		init = f(init, value.Unwrap())
	}
}

// rev
func (self HashMapKeys[K]) Rev() Iterator[K] {
	newVec := VecWithLen[K](self.vec.Len())
	i := self.vec.Len() - 1

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.AsSlice()[i] = value.Unwrap()
		i--
	}
}

func (self HashMapKeys[K]) CollectToVec() Vec[K] {
	return self.vec
}

type HashMapValues[V any] struct {
	vec      Vec[V]
	position Int
}

// An iterator visiting all values in arbitrary order. The iterator element type is V.
func (self HashMap[K, V]) Values() Iterator[V] {
	vec := Vec[V]{}
	for _, value := range self.data {
		vec.Push(value)
	}

	return &HashMapValues[V]{vec: vec, position: 0}
}

// next
func (self *HashMapValues[V]) Next() Option[V] {
	if self.position >= self.vec.Len() {
		return None[V]()
	}

	value := self.vec.GetUnchecked(self.position)
	self.position++

	return Some[V](value)
}

// map
func (self HashMapValues[V]) Map(f func(V) V) Iterator[V] {
	newVec := VecNew[V]()

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.Push(f(value.Unwrap()))
	}
}

// filter
func (self HashMapValues[V]) Filter(f func(V) Bool) Iterator[V] {
	newVec := VecNew[V]()

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

// fold
func (self HashMapValues[V]) Fold(init V, f func(V, V) V) V {
	for {
		value := self.Next()

		if value.IsNone() {
			return init
		}

		init = f(init, value.Unwrap())
	}
}

// rev
func (self HashMapValues[V]) Rev() Iterator[V] {
	newVec := VecWithLen[V](self.vec.Len())
	i := self.vec.Len() - 1

	for {
		value := self.Next()

		if value.IsNone() {
			return newVec.IntoIter()
		}
		newVec.AsSlice()[i] = value.Unwrap()
		i--
	}
}

func (self HashMapValues[V]) CollectToVec() Vec[V] {
	return self.vec
}

// impl Display for HashMap
func (self HashMap[K, V]) Display() String {
	buffer := String("")
	buffer += "HashMap{"

	fields := []string{}
	for key, value := range self.data {
		fields = append(fields, string(Format("{}: {}", key, value)))
	}

	buffer += String(strings.Join(fields, ", "))

	buffer += "}"

	return buffer
}

// impl Debug for HashMap
func (self HashMap[K, V]) Debug() String {
	return self.Display()
}
