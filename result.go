package gost

import (
	"fmt"
	"reflect"
)

// Error handling with the Result type.
// Result<T> is the type used for returning and propagating errors.
// It is an enum with the variants, Ok(T), representing success and containing a value, and Err, representing error and containing an error value.
type Result[T any] struct {
	ok  *T
	err error
}

// Creates a new Result<T> containing an Ok value.
func Ok[T any](value T) Result[T] {
	return Result[T]{ok: &value}
}

// Creates a new Result<T> containing an Err value.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// Returns true if the result is Ok.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.Assert(x.IsOk())
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.Assert(!y.IsOk())
func (self Result[T]) IsOk() Bool {
	return self.ok != nil
}

// Returns true if the result is Ok and the value inside of it matches a predicate.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.Assert(x.IsOkAnd(func(value Int) Bool { return value == 2 }))
//
//	y := gost.Ok[gost.I32](gost.I32(3))
//	gost.Assert(!y.IsOkAnd(func(value Int) Bool { return value == 2 }))
func (self Result[T]) IsOkAnd(predicate func(T) Bool) Bool {
	if self.IsOk() {
		return predicate(*self.ok)
	} else {
		return false
	}
}

// Returns true if the result is Err.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.Assert(!x.IsErr())
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.Assert(y.IsErr())
func (self Result[T]) IsErr() Bool {
	return self.err != nil
}

// Returns true if the result is Err and the value inside of it matches a predicate.
//
//	x := gost.Err[gost.I32](errors.New("error"))
//	gost.Assert(x.IsErrAnd(func(err error) Bool { return err.Error() == "error" }))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.Assert(!y.IsErrAnd(func(err error) Bool { return err.Error() == "error2" }))
func (self Result[T]) IsErrAnd(predicate func(error) Bool) Bool {
	if self.IsErr() {
		return predicate(self.err)
	} else {
		return false
	}
}

// Converts from Result<T, E> to Option<T>.
// Converts self into an Option<T>, consuming self, and discarding the error, if any.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.Ok(), gost.Some[gost.I32](gost.I32(2)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.Ok(), gost.None[gost.I32]())
func (self Result[T]) Ok() Option[T] {
	if self.IsOk() {
		return Some[T](*self.ok)
	} else {
		return None[T]()
	}
}

// Converts from Result<T, E> to Option<E>.
// Converts self into an Option<E>, consuming self, and discarding the success value, if any.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.Err(), gost.None[error]())
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.Err(), gost.Some[error](errors.New("error")))
func (self Result[T]) Err() Option[error] {
	if self.IsErr() {
		return Some[error](self.err)
	} else {
		return None[error]()
	}
}

// Maps a Result<T, E> to Result<U, E> by applying a function to a contained Ok value, leaving an Err value untouched.
// This function can be used to compose the results of two functions.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.Map(func(value Int) Int { return value + 1 }), gost.Ok[gost.I32](gost.I32(3)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.Map(func(value Int) Int { return value + 1 }), gost.Err[gost.I32](errors.New("error")))
func (self Result[T]) Map(f func(T) T) Result[T] {
	if self.IsOk() {
		return Ok[T](f(*self.ok))
	} else {
		return self
	}
}

// Returns the provided default (if Err), or applies a function to the contained value (if Ok),
// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use map_or_else, which is lazily evaluated.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.MapOr(gost.I32(3), func(value Int) Int { return value + 1 }), gost.I32(3))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.MapOr(gost.I32(3), func(value Int) Int { return value + 1 }), gost.I32(3))
func (self Result[T]) MapOr(defaultValue T, f func(T) T) T {
	if self.IsOk() {
		return f(*self.ok)
	} else {
		return defaultValue
	}
}

// Maps a Result<T, E> to U by applying fallback function default to a contained Err value, or function f to a contained Ok value.
// This function can be used to unpack a successful result while handling an error.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.MapOrElse(func() Int { return gost.I32(3) }, func(value Int) Int { return value + 1 }), gost.I32(3))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.MapOrElse(func() Int { return gost.I32(3) }, func(value Int) Int { return value + 1 }), gost.I32(3))
func (self Result[T]) MapOrElse(defaultValue func() T, f func(T) T) T {
	if self.IsOk() {
		return f(*self.ok)
	} else {
		return defaultValue()
	}
}

// Maps a Result<T, E> to Result<T, F> by applying a function to a contained Err value, leaving an Ok value untouched.
// This function can be used to pass through a successful result while handling an error.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.MapErr(func(err error) error { return errors.New("error2") }), gost.Ok[gost.I32](gost.I32(2)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.MapErr(func(err error) error { return errors.New("error2") }), gost.Err[gost.I32](errors.New("error2")))
func (self Result[T]) MapErr(f func(error) error) Result[T] {
	if self.IsErr() {
		return Err[T](f(self.err))
	} else {
		return self
	}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.Expect("error"), gost.I32(2))
func (self Result[T]) Expect(message string) T {
	if self.IsOk() {
		return *self.ok
	} else {
		panic(message)
	}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.Unwrap(), gost.I32(2))
func (self Result[T]) Unwrap() T {
	if self.IsOk() {
		return *self.ok
	} else {
		panic(self.err)
	}
}

// Returns the contained Err value, consuming the self value.
//
//	x := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(x.ExpectErr("error"), errors.New("error"))
func (self Result[T]) ExpectErr(message string) error {
	if self.IsErr() {
		return self.err
	} else {
		panic(message)
	}
}

// Returns the contained Err value, consuming the self value.
//
//	x := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(x.UnwrapErr(), errors.New("error"))
func (self Result[T]) UnwrapErr() error {
	if self.IsErr() {
		return self.err
	} else {
		panic(self.ok)
	}
}

// Returns res if the result is Ok, otherwise returns the Err value of self.
// Arguments passed to and are eagerly evaluated; if you are passing the result of a function call, it is recommended to use and_then, which is lazily evaluated.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.And(gost.Ok[gost.I32](gost.I32(3))), gost.Ok[gost.I32](gost.I32(3)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.And(gost.Ok[gost.I32](gost.I32(3))), gost.Err[gost.I32](errors.New("error")))
func (self Result[T]) And(res Result[T]) Result[T] {
	if self.IsOk() {
		return res
	} else {
		return self
	}
}

// Calls op if the result is Ok, otherwise returns the Err value of self.
// This function can be used for control flow based on Result values.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.AndThen(func(value Int) Result[Int] { return gost.Ok[Int](value + 1) }), gost.Ok[gost.I32](gost.I32(3)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.AndThen(func(value Int) Result[Int] { return gost.Ok[Int](value + 1) }), gost.Err[gost.I32](errors.New("error")))
func (self Result[T]) AndThen(op func(T) Result[T]) Result[T] {
	if self.IsOk() {
		return op(*self.ok)
	} else {
		return self
	}
}

// Returns res if the result is Err, otherwise returns the Ok value of self.
// Arguments passed to or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use or_else, which is lazily evaluated.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.Or(gost.Ok[gost.I32](gost.I32(3))), gost.Ok[gost.I32](gost.I32(2)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.Or(gost.Ok[gost.I32](gost.I32(3))), gost.Ok[gost.I32](gost.I32(3)))
func (self Result[T]) Or(res Result[T]) Result[T] {
	if self.IsErr() {
		return res
	} else {
		return self
	}
}

// Calls op if the result is Err, otherwise returns the Ok value of self.
// This function can be used for control flow based on result values.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.OrElse(func(err error) Result[Int] { return gost.Ok[Int](gost.I32(3)) }), gost.Ok[gost.I32](gost.I32(2)))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.OrElse(func(err error) Result[Int] { return gost.Ok[Int](gost.I32(3)) }), gost.Ok[gost.I32](gost.I32(3)))
func (self Result[T]) OrElse(op func(error) Result[T]) Result[T] {
	if self.IsErr() {
		return op(self.err)
	} else {
		return self
	}
}

// Returns the contained Ok value or a provided default.
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use unwrap_or_else, which is lazily evaluated.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.UnwrapOr(gost.I32(3)), gost.I32(2))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.UnwrapOr(gost.I32(3)), gost.I32(3))
func (self Result[T]) UnwrapOr(defaultValue T) T {
	if self.IsOk() {
		return *self.ok
	} else {
		return defaultValue
	}
}

// Returns the contained Ok value or computes it from a closure.
//
//	x := gost.Ok[gost.I32](gost.I32(2))
//	gost.AssertEq(x.UnwrapOrElse(func() Int { return gost.I32(3) }), gost.I32(2))
//
//	y := gost.Err[gost.I32](errors.New("error"))
//	gost.AssertEq(y.UnwrapOrElse(func() Int { return gost.I32(3) }), gost.I32(3))
func (self Result[T]) UnwrapOrElse(f func() T) T {
	if self.IsOk() {
		return *self.ok
	} else {
		return f()
	}
}

// impl Display for Result
func (self Result[T]) Display() String {
	if self.IsOk() {
		value := reflect.ValueOf(self.ok)

		if display, ok := value.Interface().(Display[T]); ok {
			return "Ok(" + display.Display() + ")"
		}

		typeName := reflect.TypeOf(self.ok).Elem().Name()

		panic(fmt.Sprintf("'%s' does not implement Display[%s]", typeName, typeName))
	} else {
		return String(self.err.Error())
	}
}

// impl Debug for Result
func (self Result[T]) Debug() String {
	return self.Display()
}

// impl AsRef for Result
func (self Result[T]) AsRef() *Result[T] {
	return &self
}

// impl Clone for Result
func (self Result[T]) Clone() Result[T] {
	if self.IsOk() {
		return Ok[T](castToClone[T](*self.ok).Unwrap().Clone())
	} else {
		return Err[T](self.err)
	}
}

// impl Eq for Result
func (self Result[T]) Eq(rhs Result[T]) Bool {
	if self.IsOk() && rhs.IsOk() {
		return castToEq[T](*self.ok).Unwrap().Eq(*rhs.ok)
	} else if self.IsErr() && rhs.IsErr() {
		return self.err == rhs.err
	} else {
		return false
	}
}
