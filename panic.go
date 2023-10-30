package gost

// Panics the current thread.
func Panic(message String, args ...any) {
	panic(Format(message, args...))
}
