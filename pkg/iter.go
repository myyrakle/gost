package gost

type Iterator[T any] interface {
	Next() Option[T]
	Map(f func(T) T) Iterator[T]
	Filter(f func(T) Bool) Iterator[T]
	Fold(init T, f func(T, T) T) T
	Rev() Iterator[T]
	CollectToVec() Vec[T]
}

type IntoIterator[T any] interface {
	IntoIter() Iterator[T]
}
