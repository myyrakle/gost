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
