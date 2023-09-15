package gost

type Iterator[T any] interface {
	Next() Option[T]
}

type IntoIterator[T any] interface {
	IntoIter() Iterator[T]
}
