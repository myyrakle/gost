package gost

import "fmt"

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// impl Display for Pair
func (p Pair[K, V]) ToString() string {
	keyToString := castToToString[K](p.Key)
	valueToString := castToToString[V](p.Value)

	var key String
	var value String

	if keyToString.IsSome() {
		key = keyToString.Unwrap().ToString()
	} else {
		key = String(fmt.Sprintf("%v", p.Key))
	}

	if valueToString.IsSome() {
		value = valueToString.Unwrap().ToString()
	} else {
		value = String(fmt.Sprintf("%v", p.Value))
	}

	return fmt.Sprintf("Pair[%s, %s]", key, value)
}

// impl Display for Pair
func (p Pair[K, V]) String() string {
	return p.ToString()
}

// impl Debug for Pair
func (p Pair[K, V]) Debug() string {
	return p.ToString()
}

// impl Clone for Pair
func (p Pair[K, V]) Clone() Pair[K, V] {
	keyClone := castToClone[K](p.Key).Unwrap().Clone()
	valueClone := castToClone[V](p.Value).Unwrap().Clone()

	return Pair[K, V]{
		Key:   keyClone,
		Value: valueClone,
	}
}
