package option

var None = Option[interface{}]{value: nil}

type Option[T any] struct {
	value *T
}

func New[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// Returns true if the option is a Some value.
func (self *Option[T]) IsSome() bool {
	return self.value != nil
}

// Returns true if the option is a Some and the value inside of it matches a predicate.
func (self *Option[T]) IsSomeAnd(f func(T) bool) bool {
	if self.IsNone() {
		return false
	} else {
		return f(*self.value)
	}
}

// Returns true if the option is a None value.
func (self *Option[T]) IsNone() bool {
	return self.value == nil
}

// Returns the contained Some value, consuming the self value.
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
func (self *Option[T]) Unwrap() T {
	return *self.value
}

// Returns the contained Some value or a provided default.
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use unwrap_or_else, which is lazily evaluated.
func (self *Option[T]) UnwrapOr(value T) T {
	if self.IsNone() {
		return value
	} else {
		return *self.value
	}
}

// Returns the contained Some value or computes it from a closure.
func (self *Option[T]) UnwrapOrElse(f func() T) T {
	if self.IsNone() {
		return f()
	} else {
		return *self.value
	}
}

// TODO: UnwrapOrDefault

// Returns None if the option is None, otherwise returns optb.
// Arguments passed to and are eagerly evaluated; if you are passing the result of a function call, it is recommended to use and_then, which is lazily evaluated.
func (self Option[T]) And(optb Option[any]) Option[any] {
	if self.IsNone() {
		return New[any](nil)
	} else {
		return optb
	}
}

// Returns None if the option is None, otherwise calls f with the wrapped value and returns the result.
// Some languages call this operation flatmap.
func (self Option[T]) AndThen(f func(T) Option[any]) Option[any] {
	if self.IsNone() {
		return New[any](nil)
	} else {
		return f(*self.value)
	}
}
