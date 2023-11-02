package gost

// Panics the current thread.
//
//  gost.Panic("This is a panic")
func Panic(message String, args ...any) {
	panic(Format(message, args...))
}

// Asserts that a boolean expression is true at runtime.
//
//  gost.Assert(true, "This is true")
func Assert(condition Bool, args ...any) {
	if !condition {
		if len(args) > 0 {
			if message, ok := args[0].(String); ok {
				panic(Format(message, args...))
			}
		}

		panic(Format("assertion failed"))
	}
}

// Asserts that two expressions are equal to each other
//
//  gost.AssertEq(1, 1, "These are equal")
func AssertEq[T Eq[T]](lhs T, rhs T, args ...any) {
	if !lhs.Eq(rhs) {
		if len(args) > 0 {
			if message, ok := args[0].(String); ok {
				panic(Format(message, args...))
			}
		}

		panic(Format("assertion failed: {} != {}", lhs, rhs))
	}
}

// Asserts that two expressions are not equal to each other
//
//  gost.AssertNe(1, 2, "These are not equal")
func AssertNe[T Eq[T]](lhs T, rhs T, args ...any) {
	if lhs.Eq(rhs) {
		if len(args) > 0 {
			if message, ok := args[0].(String); ok {
				panic(Format(message, args...))
			}
		}

		panic(Format("assertion failed: {} == {}", lhs, rhs))
	}
}

// Indicates unimplemented code by panicking with a message of “not implemented”.
//
//  gost.Unimplemented()
func Unimplemented(messages ...String) {
	if len(messages) == 0 {
		panic("not implemented")
	} else {
		panic(messages[0])
	}
}

// Indicates unreachable code.
//
//  gost.Unreachable()
func Unreachable(messages ...String) {
	if len(messages) == 0 {
		panic("unreachable")
	} else {
		panic(messages[0])
	}
}

// Indicates unfinished code.
//
//  gost.Todo()
func Todo(messages ...String) {
	if len(messages) == 0 {
		panic("todo")
	} else {
		panic(messages[0])
	}
}
