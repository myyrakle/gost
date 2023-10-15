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

// Add implements
func (self ISize) Add(rhs ISize) ISize {
	return self + rhs
}

func (self I8) Add(rhs I8) I8 {
	return self + rhs
}

func (self I16) Add(rhs I16) I16 {
	return self + rhs
}

func (self I32) Add(rhs I32) I32 {
	return self + rhs
}

func (self I64) Add(rhs I64) I64 {
	return self + rhs
}

func (self USize) Add(rhs USize) USize {
	return self + rhs
}

func (self U8) Add(rhs U8) U8 {
	return self + rhs
}

func (self U16) Add(rhs U16) U16 {
	return self + rhs
}

func (self U32) Add(rhs U32) U32 {
	return self + rhs
}

func (self U64) Add(rhs U64) U64 {
	return self + rhs
}

func (self F32) Add(rhs F32) F32 {
	return self + rhs
}

func (self F64) Add(rhs F64) F64 {
	return self + rhs
}

func (self Byte) Add(rhs Byte) Byte {
	return self + rhs
}

func (self Char) Add(rhs Char) Char {
	return self + rhs
}

func (self String) Add(rhs String) String {
	return self + rhs
}

func (self Complex64) Add(rhs Complex64) Complex64 {
	return self + rhs
}

func (self Complex128) Add(rhs Complex128) Complex128 {
	return self + rhs
}
