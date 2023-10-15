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

// Sub implements
func (self ISize) Sub(rhs ISize) ISize {
	return self - rhs
}

func (self I8) Sub(rhs I8) I8 {
	return self - rhs
}

func (self I16) Sub(rhs I16) I16 {
	return self - rhs
}

func (self I32) Sub(rhs I32) I32 {
	return self - rhs
}

func (self I64) Sub(rhs I64) I64 {
	return self - rhs
}

func (self USize) Sub(rhs USize) USize {
	return self - rhs
}

func (self U8) Sub(rhs U8) U8 {
	return self - rhs
}

func (self U16) Sub(rhs U16) U16 {
	return self - rhs
}

func (self U32) Sub(rhs U32) U32 {
	return self - rhs
}

func (self U64) Sub(rhs U64) U64 {
	return self - rhs
}

func (self F32) Sub(rhs F32) F32 {
	return self - rhs
}

func (self F64) Sub(rhs F64) F64 {
	return self - rhs
}

func (self Byte) Sub(rhs Byte) Byte {
	return self - rhs
}

func (self Char) Sub(rhs Char) Char {
	return self - rhs
}

func (self Complex64) Sub(rhs Complex64) Complex64 {
	return self - rhs
}

func (self Complex128) Sub(rhs Complex128) Complex128 {
	return self - rhs
}

// Mul implements
func (self ISize) Mul(rhs ISize) ISize {
	return self * rhs
}

func (self I8) Mul(rhs I8) I8 {

	return self * rhs
}

func (self I16) Mul(rhs I16) I16 {
	return self * rhs
}

func (self I32) Mul(rhs I32) I32 {
	return self * rhs
}

func (self I64) Mul(rhs I64) I64 {
	return self * rhs
}

func (self USize) Mul(rhs USize) USize {
	return self * rhs
}

func (self U8) Mul(rhs U8) U8 {
	return self * rhs
}

func (self U16) Mul(rhs U16) U16 {
	return self * rhs
}

func (self U32) Mul(rhs U32) U32 {
	return self * rhs
}

func (self U64) Mul(rhs U64) U64 {
	return self * rhs
}

func (self F32) Mul(rhs F32) F32 {
	return self * rhs
}

func (self F64) Mul(rhs F64) F64 {
	return self * rhs
}

func (self Byte) Mul(rhs Byte) Byte {
	return self * rhs
}

func (self Char) Mul(rhs Char) Char {
	return self * rhs
}

func (self Complex64) Mul(rhs Complex64) Complex64 {
	return self * rhs
}

func (self Complex128) Mul(rhs Complex128) Complex128 {
	return self * rhs
}
