package gost

type Iterator[T any] interface {
	Next() Option[T]
	Map(f func(T) T) Map[T]
}

type IntoIterator[T any] interface {
	IntoIter() Iterator[T]
}

type Map[T any] struct {
	iter Iterator[T]
}

func (self Map[T]) CollectToVec() Vec[T] {
	vec := Vec[T]{}
	for {
		value := self.iter.Next()
		if value.IsNone() {
			return vec
		}
		vec.Push(value.Unwrap())
	}
}

type Filter[T any] struct {
	iter Iterator[T]
}

func (self Filter[T]) CollectToVec() Vec[T] {
	vec := Vec[T]{}
	for {
		value := self.iter.Next()
		if value.IsNone() {
			return vec
		}
		vec.Push(value.Unwrap())
	}
}
