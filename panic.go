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
func AssertNe[T Eq[T]](lhs T, rhs T, message String, args ...any) {
	if lhs.Eq(rhs) {
		panic(Format(message, args...))
	}
}

// Indicates unimplemented code by panicking with a message of “not implemented”.
func Unimplemented(messages ...String) {
	if len(messages) == 0 {
		panic("not implemented")
	} else {
		panic(messages[0])
	}
}

// Indicates unreachable code.
func Unreachable(messages ...String) {
	if len(messages) == 0 {
		panic("unreachable")
	} else {
		panic(messages[0])
	}
}

// Indicates unfinished code.
func Todo(messages ...String) {
	if len(messages) == 0 {
		panic("todo")
	} else {
		panic(messages[0])
	}
}
