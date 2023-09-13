package result

type Result[T any, Err error] struct {
	ok  *T
	err Err
}
