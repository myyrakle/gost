package gost

// Add is a trait for types that support addition.
type Add[T any] interface {
	Add(rhs T) T
}

// Sub is a trait for types that support subtraction.
type Sub[T any] interface {
	Sub(rhs T) T
}

// Mul is a trait for types that support multiplication.
type Mul[T any] interface {
	Mul(rhs T) T
}

// Div is a trait for types that support division.
type Div[T any] interface {
	Div(rhs T) T
}

// The addition assignment operator +=.
type AddAssign[T any] interface {
	AddAssign(rhs T)
}

// The subtraction assignment operator -=.
type SubAssign[T any] interface {
	SubAssign(rhs T)
}

// The multiplication assignment operator *=.
type MulAssign[T any] interface {
	MulAssign(rhs T)
}

// The division assignment operator /=.
type DivAssign[T any] interface {
	DivAssign(rhs T)
}
