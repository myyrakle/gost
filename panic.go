package gost

// Panics the current thread.
func Panic(message String, args ...any) {
	panic(Format(message, args...))
}

// Asserts that a boolean expression is true at runtime.
func Assert(condition Bool, message String, args ...any) {
	if !condition {
		panic(Format(message, args...))
	}
}

// Asserts that two expressions are equal to each other
func AssertEq[T Eq[T]](lhs T, rhs T, message String, args ...any) {
	if !lhs.Eq(rhs) {
		panic(Format(message, args...))
	}
}

// Asserts that two expressions are not equal to each other
func AssertNotEq[T Eq[T]](lhs T, rhs T, message String, args ...any) {
	if lhs.Eq(rhs) {
		panic(Format(message, args...))
	}
}
