package trait

// Add is a trait for types that support addition.
type Add[T any] interface {
	Add(T, T) T
}

// Sub is a trait for types that support subtraction.
type Sub[T any] interface {
	Sub(T, T) T
}

// Mul is a trait for types that support multiplication.
type Mul[T any] interface {
	Mul(T, T) T
}

// Div is a trait for types that support division.
type Div[T any] interface {
	Div(T, T) T
}
