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

// impl Eq for Pair
func (p Pair[K, V]) Eq(other Pair[K, V]) Bool {
	keyEq := castToEq[K](p.Key).Unwrap().Eq(other.Key)
	valueEq := castToEq[V](p.Value).Unwrap().Eq(other.Value)

	return keyEq && valueEq
}

// impl Ord for Pair
func (p Pair[K, V]) Cmp(other Pair[K, V]) Ordering {
	keyOrd := castToOrd[K](p.Key).Unwrap().Cmp(other.Key)

	if keyOrd == OrderingEqual {
		valueOrd := castToOrd[V](p.Value).Unwrap().Cmp(other.Value)
		return valueOrd
	} else {
		return keyOrd
	}
}
