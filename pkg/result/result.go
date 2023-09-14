package result

import "github.com/myyrakle/gost/pkg/option"

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

// Converts from Result<T, E> to Option<T>.
// Converts self into an Option<T>, consuming self, and discarding the error, if any.
func (self Result[T]) Ok() option.Option[T] {
	if self.IsOk() {
		return option.New[T](*self.ok)
	} else {
		return option.None[T]()
	}
}

// Converts from Result<T, E> to Option<E>.
// Converts self into an Option<E>, consuming self, and discarding the success value, if any.
func (self Result[T]) Err() option.Option[error] {
	if self.IsErr() {
		return option.New[error](self.err)
	} else {
		return option.None[error]()
	}
}
