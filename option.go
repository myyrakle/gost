package gost

import (
	"fmt"
	"reflect"
)

// Type Option represents an optional value: every Option is either Some and contains a value, or None, and does not.
// Option types are very common in Rust code
type Option[T any] struct {
	value *T
}

// Some returns an Option<T> with a Some value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// None returns an Option<T> with a None value.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// Returns true if the option is a Some value.
//
//	x := Some(gost.I32(2))
//	gost.Assert(x.IsSome())
func (self Option[T]) IsSome() Bool {
	return self.value != nil
}

// Returns true if the option is a Some and the value inside of it matches a predicate.
//
//	x := Some(gost.I32(2))
//	gost.Assert(x.IsSomeAnd(func(value gost.I32) gost.Bool { return value == 2 }))
func (self *Option[T]) IsSomeAnd(f func(T) Bool) Bool {
	if self.IsNone() {
		return false
	} else {
		return f(*self.value)
	}
}

// Returns true if the option is a None value.
//
//	x := None[gost.I32]()
//	gost.Assert(x.IsNone())
func (self Option[T]) IsNone() Bool {
	return self.value == nil
}

// Returns the contained Some value, consuming the self value.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.Expect(gost.String("failed")), gost.I32(2))
func (self Option[T]) Expect(message string) T {
	if self.IsNone() {
		panic(message)
	} else {
		return *self.value
	}
}

// Returns the contained Some value, consuming the self value.
// Because this function may panic, its use is generally discouraged.
// Instead, prefer to use pattern matching and handle the None case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.Unwrap(), gost.I32(2))
func (self Option[T]) Unwrap() T {
	return *self.value
}

// Returns the contained Some value or a provided default.
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use unwrap_or_else, which is lazily evaluated.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.UnwrapOr(gost.I32(3)), gost.I32(2))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.UnwrapOr(gost.I32(3)), gost.I32(3))
func (self Option[T]) UnwrapOr(value T) T {
	if self.IsNone() {
		return value
	} else {
		return *self.value
	}
}

// Returns the contained Some value or computes it from a closure.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.UnwrapOrElse(func() gost.I32 { return gost.I32(3) }), gost.I32(2))
func (self Option[T]) UnwrapOrElse(f func() T) T {
	if self.IsNone() {
		return f()
	} else {
		return *self.value
	}
}

// TODO: UnwrapOrDefault

// Maps an Option<T> to other Option<T> by applying a function to a contained value (if Some) or returns None (if None).
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.Map(func(value gost.I32) gost.I32 { return value + 1 }), Some(gost.I32(3)))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.Map(func(value gost.I32) gost.I32 { return value + 1 }), None[gost.I32]())
func (self Option[T]) Map(f func(T) T) Option[T] {
	if self.IsNone() {
		return self
	} else {
		return Some[T](f(*self.value))
	}
}

// Returns the provided default result (if none), or applies a function to the contained value (if any).
// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use map_or_else, which is lazily evaluated.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.MapOr(gost.I32(3), func(value gost.I32) gost.I32 { return value + 1 }), gost.I32(3))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.MapOr(gost.I32(3), func(value gost.I32) gost.I32 { return value + 1 }), gost.I32(3))
func (self Option[T]) MapOr(defaultValue T, f func(T) T) T {
	if self.IsNone() {
		return defaultValue
	} else {
		return f(*self.value)
	}
}

// Computes a default function result (if none), or applies a different function to the contained value (if any).
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.MapOrElse(func() gost.I32 { return gost.I32(3) }, func(value gost.I32) gost.I32 { return value + 1 }), gost.I32(3))
func (self Option[T]) MapOrElse(defaultValue func() T, f func(T) T) T {
	if self.IsNone() {
		return defaultValue()
	} else {
		return f(*self.value)
	}
}

// Transforms the Option<T> into a Result<T, E>, mapping Some(v) to Ok(v) and None to Err(err).
// Arguments passed to ok_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use ok_or_else, which is lazily evaluated.
// TODO ok_or

// Transforms the Option<T> into a Result<T, E>, mapping Some(v) to Ok(v) and None to Err(err()).
// TODO ok_or_else

// TODO iter

// Returns None if the option is None, otherwise returns optb.
// Arguments passed to and are eagerly evaluated; if you are passing the result of a function call, it is recommended to use and_then, which is lazily evaluated.
//
//	x := Some(gost.I32(2))
//	y := None[gost.I32]()
//	gost.AssertEq(x.and(y), None);
//
//	x := Some(gost.I32(2))
//	y := Some(gost.I32(3))
//	gost.AssertEq(x.and(y), Some(gost.I32(3)));
//
//	x := None[gost.I32]()
//	y := None[gost.I32]()
//	gost.AssertEq(x.and(y), None);
//
//	x := Some(gost.I32(2))
//	y := Some(gost.I32(3))
//	gost.AssertEq(x.and(y), Some(gost.I32(3)));
func (self Option[T]) And(optb Option[any]) Option[any] {
	if self.IsNone() {
		return Some[any](nil)
	} else {
		return optb
	}
}

// Returns None if the option is None, otherwise calls f with the wrapped value and returns the result.
// Some languages call this operation flatmap.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.and_then(func(value gost.I32) Option[gost.I32] { return Some[gost.I32](value + 1) }), Some(gost.I32(3)))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.and_then(func(value gost.I32) Option[gost.I32] { return Some[gost.I32](value + 1) }), None[gost.I32]())
func (self Option[T]) AndThen(f func(T) Option[any]) Option[any] {
	if self.IsNone() {
		return Some[any](nil)
	} else {
		return f(*self.value)
	}
}

// Returns None if the option is None, otherwise calls predicate with the wrapped value and returns:
// 1. Some(t) if predicate returns true (where t is the wrapped value), and
// 2. None if predicate returns false.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.filter(func(value gost.I32) gost.Bool { return value == 2 }), Some(gost.I32(2)))
func (self Option[T]) Filter(predicate func(T) Bool) Option[T] {
	if self.IsNone() {
		return self
	} else if predicate(*self.value) {
		return self
	} else {
		return None[T]()
	}
}

// Returns the option if it contains a value, otherwise returns optb.
// Arguments passed to or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use or_else, which is lazily evaluated.
//
//	x := Some(gost.I32(2))
//	y := None[gost.I32]()
//	gost.AssertEq(x.or(y), Some(gost.I32(2)));
//
//	x := Some(gost.I32(2))
//	y := Some(gost.I32(3))
//	gost.AssertEq(x.or(y), Some(gost.I32(2)));
func (self Option[T]) Or(optb Option[T]) Option[T] {
	if self.IsNone() {
		return optb
	} else {
		return self
	}
}

// Returns the option if it contains a value, otherwise calls f and returns the result.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.or_else(func() Option[gost.I32] { return Some[gost.I32](3) }), Some(gost.I32(2)))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.or_else(func() Option[gost.I32] { return Some[gost.I32](3) }), Some(gost.I32(3)))
func (self Option[T]) OrElse(f func() Option[T]) Option[T] {
	if self.IsNone() {
		return f()
	} else {
		return self
	}
}

// Returns Some if exactly one of self, optb is Some, otherwise returns None.
//
//	x := Some(gost.I32(2))
//	y := None[gost.I32]()
//	gost.AssertEq(x.Xor(y), Some(gost.I32(2)));
//
//	x := Some(gost.I32(2))
//	y := Some(gost.I32(3))
//	gost.AssertEq(x.Xor(y), None[gost.I32]());
func (self Option[T]) Xor(optb Option[T]) Option[T] {
	if self.IsNone() {
		return optb
	} else if optb.IsNone() {
		return self
	} else {
		return None[T]()
	}
}

// Inserts value into the option, then returns a mutable reference to it.
// If the option already contains a value, the old value is dropped.
// See also Option::get_or_insert, which doesn’t update the value if the option already contains Some.
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.Insert(gost.I32(2)), Some(gost.I32(2)))
//
//	x := Some(gost.I32(3))
//	gost.AssertEq(x.Insert(gost.I32(2)), Some(gost.I32(2)))
func (self *Option[T]) Insert(value T) *T {
	self.value = &value
	return self.value
}

// Inserts value into the option if it is None, then returns a mutable reference to the contained value.
// See also Option::insert, which updates the value even if the option already contains Some.
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.GetOrInsert(gost.I32(2)), Some(gost.I32(2)))
//
//	x := Some(gost.I32(3))
//	gost.AssertEq(x.GetOrInsert(gost.I32(2)), Some(gost.I32(3)))
func (self *Option[T]) GetOrInsert(value T) *T {
	if self.IsNone() {
		self.value = &value
	}
	return self.value
}

// Inserts a value computed from f into the option if it is None, then returns a mutable reference to the contained value.
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.GetOrInsertWith(func() gost.I32 { return gost.I32(2) }), Some(gost.I32(2)))
//
//	x := Some(gost.I32(3))
//	gost.AssertEq(x.GetOrInsertWith(func() gost.I32 { return gost.I32(2) }), Some(gost.I32(3)))
func (self *Option[T]) GetOrInsertWith(f func() T) *T {
	if self.IsNone() {
		value := f()
		self.value = &value
	}
	return self.value
}

// Takes the value out of the option, leaving a None in its place.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.Take(), Some(gost.I32(2)))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.Take(), None[gost.I32]())
func (self *Option[T]) Take() Option[T] {
	if self.IsNone() {
		return *self
	} else {
		value := *self.value
		self.value = nil
		return Some[T](value)
	}
}

// Replaces the actual value in the option by the value given in parameter, returning the old value if present, leaving a Some in its place without deinitializing either one.
//
//	x := Some(gost.I32(2))
//	gost.AssertEq(x.Replace(gost.I32(3)), Some(gost.I32(2)))
//
//	x := None[gost.I32]()
//	gost.AssertEq(x.Replace(gost.I32(3)), None[gost.I32]())
func (self *Option[T]) Replace(value T) Option[T] {
	if self.IsNone() {
		return *self
	} else {
		oldValue := *self.value
		self.value = &value
		return Some[T](oldValue)
	}
}

// impl Display for Option
func (self Option[T]) Display() String {
	if self.IsNone() {
		return "None"
	} else {
		value := reflect.ValueOf(self.value)

		if display, ok := value.Interface().(Display[T]); ok {
			return "Some(" + display.Display() + ")"
		}

		typeName := reflect.TypeOf(self.value).Elem().Name()

		panic(fmt.Sprintf("'%s' does not implement Display[%s]", typeName, typeName))
	}
}

// impl Debug for Option
func (self Option[T]) Debug() String {
	return self.Display()
}

// impl AsRef for Option
func (self Option[T]) AsRef() *Option[T] {
	return &self
}

// impl Clone for Option
func (self Option[T]) Clone() Option[T] {
	if self.IsNone() {
		return None[T]()
	} else {
		return Some[T](castToClone[T](*self.value).Unwrap().Clone())
	}
}

// impl Eq for Option
func (self Option[T]) Eq(rhs Option[T]) Bool {
	if self.IsNone() && rhs.IsNone() {
		return true
	} else if self.IsSome() && rhs.IsSome() {
		return castToEq[T](*self.value).Unwrap().Eq(*rhs.value)
	} else {
		return false
	}
}

type _InternalOption[T any] struct {
	value *T
}

func (self _InternalOption[T]) IsSome() Bool {
	return self.value != nil
}

func (self _InternalOption[T]) IsNone() Bool {
	return self.value == nil
}

func (self _InternalOption[T]) Unwrap() T {
	return *self.value
}

func _Some[T any](value T) _InternalOption[T] {
	return _InternalOption[T]{value: &value}
}

func _None[T any]() _InternalOption[T] {
	return _InternalOption[T]{value: nil}
}
