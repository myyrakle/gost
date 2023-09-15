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

// The addition assignment operator +=.
type AddAssign[Lhs any, Rhs any] interface {
	AddAssign(Lhs, Rhs)
}

// The subtraction assignment operator -=.
type SubAssign[Lhs any, Rhs any] interface {
	SubAssign(Lhs, Rhs)
}

// The multiplication assignment operator *=.
type MulAssign[Lhs any, Rhs any] interface {
	MulAssign(Lhs, Rhs)
}

// The division assignment operator /=.
type DivAssign[Lhs any, Rhs any] interface {
	DivAssign(Lhs, Rhs)
}
