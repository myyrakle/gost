package gost

type AsRef[T any] interface {
	AsRef() *T
}
