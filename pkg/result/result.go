package result

import "github.com/myyrakle/gost/pkg/option"

type Result[T any] struct {
	ok  *T
	err error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{ok: &value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
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
		return option.Some[T](*self.ok)
	} else {
		return option.None[T]()
	}
}

// Converts from Result<T, E> to Option<E>.
// Converts self into an Option<E>, consuming self, and discarding the success value, if any.
func (self Result[T]) Err() option.Option[error] {
	if self.IsErr() {
		return option.Some[error](self.err)
	} else {
		return option.None[error]()
	}
}

// Maps a Result<T, E> to Result<U, E> by applying a function to a contained Ok value, leaving an Err value untouched.
// This function can be used to compose the results of two functions.
func (self Result[T]) Map(f func(T) T) Result[T] {
	if self.IsOk() {
		return Ok[T](f(*self.ok))
	} else {
		return self
	}
}

// Returns the provided default (if Err), or applies a function to the contained value (if Ok),
// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use map_or_else, which is lazily evaluated.
func (self Result[T]) MapOr(defaultValue T, f func(T) T) T {
	if self.IsOk() {
		return f(*self.ok)
	} else {
		return defaultValue
	}
}

// Maps a Result<T, E> to U by applying fallback function default to a contained Err value, or function f to a contained Ok value.
// This function can be used to unpack a successful result while handling an error.
func (self Result[T]) MapOrElse(defaultValue func() T, f func(T) T) T {
	if self.IsOk() {
		return f(*self.ok)
	} else {
		return defaultValue()
	}
}

// Maps a Result<T, E> to Result<T, F> by applying a function to a contained Err value, leaving an Ok value untouched.
// This function can be used to pass through a successful result while handling an error.
func (self Result[T]) MapErr(f func(error) error) Result[T] {
	if self.IsErr() {
		return Err[T](f(self.err))
	} else {
		return self
	}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
func (self Result[T]) Expect(message string) T {
	if self.IsOk() {
		return *self.ok
	} else {
		panic(message)
	}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
func (self Result[T]) Unwrap() T {
	if self.IsOk() {
		return *self.ok
	} else {
		panic(self.err)
	}
}

// Returns the contained Err value, consuming the self value.
func (self Result[T]) ExpectErr(message string) error {
	if self.IsErr() {
		return self.err
	} else {
		panic(message)
	}
}

// Returns the contained Err value, consuming the self value.
func (self Result[T]) UnwrapErr() error {
	if self.IsErr() {
		return self.err
	} else {
		panic(self.ok)
	}
}
