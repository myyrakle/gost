package result

type Result[T any] struct {
	ok  *T
	err error
}

// Returns true if the result is Ok.
func (self Result[T]) IsOk() bool {
	return self.ok != nil
}

// Returns true if the result is Ok and the value inside of it matches a predicate.
func (self Result[T]) IsOkAnd(predicate func(T) bool) bool {
	if self.IsOk() {
		return predicate(*self.ok)
	} else {
		return false
	}
}

// Returns true if the result is Err.
func (self Result[T]) IsErr() bool {
	return self.err != nil
}

// Returns true if the result is Err and the value inside of it matches a predicate.
func (self Result[T]) IsErrAnd(predicate func(error) bool) bool {
	if self.IsErr() {
		return predicate(self.err)
	} else {
		return false
	}
}
