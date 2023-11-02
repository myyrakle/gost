package gost

import "math"

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

// Div implements
func (self ISize) Div(rhs ISize) ISize {
	return self / rhs
}

func (self I8) Div(rhs I8) I8 {
	return self / rhs
}

func (self I16) Div(rhs I16) I16 {
	return self / rhs
}

func (self I32) Div(rhs I32) I32 {
	return self / rhs
}

func (self I64) Div(rhs I64) I64 {
	return self / rhs
}

func (self USize) Div(rhs USize) USize {
	return self / rhs
}

func (self U8) Div(rhs U8) U8 {
	return self / rhs
}

func (self U16) Div(rhs U16) U16 {
	return self / rhs
}

func (self U32) Div(rhs U32) U32 {
	return self / rhs
}

func (self U64) Div(rhs U64) U64 {
	return self / rhs
}

func (self F32) Div(rhs F32) F32 {
	return self / rhs
}

func (self F64) Div(rhs F64) F64 {
	return self / rhs
}

func (self Byte) Div(rhs Byte) Byte {
	return self / rhs
}

func (self Char) Div(rhs Char) Char {
	return self / rhs
}

func (self Complex64) Div(rhs Complex64) Complex64 {
	return self / rhs
}

func (self Complex128) Div(rhs Complex128) Complex128 {
	return self / rhs
}

// AddAssign implements
func (self *ISize) AddAssign(rhs ISize) {
	*self += rhs
}

func (self *I8) AddAssign(rhs I8) {
	*self += rhs
}

func (self *I16) AddAssign(rhs I16) {
	*self += rhs
}

func (self *I32) AddAssign(rhs I32) {
	*self += rhs
}

func (self *I64) AddAssign(rhs I64) {
	*self += rhs
}

func (self *USize) AddAssign(rhs USize) {
	*self += rhs
}

func (self *U8) AddAssign(rhs U8) {
	*self += rhs
}

func (self *U16) AddAssign(rhs U16) {
	*self += rhs
}

func (self *U32) AddAssign(rhs U32) {
	*self += rhs
}

func (self *U64) AddAssign(rhs U64) {
	*self += rhs
}

func (self *F32) AddAssign(rhs F32) {
	*self += rhs
}

func (self *F64) AddAssign(rhs F64) {
	*self += rhs
}

func (self *Byte) AddAssign(rhs Byte) {
	*self += rhs
}

func (self *Char) AddAssign(rhs Char) {
	*self += rhs
}

func (self *String) AddAssign(rhs String) {
	*self += rhs
}

func (self *Complex64) AddAssign(rhs Complex64) {
	*self += rhs
}

func (self *Complex128) AddAssign(rhs Complex128) {
	*self += rhs
}

// SubAssign implements
func (self *ISize) SubAssign(rhs ISize) {
	*self -= rhs
}

func (self *I8) SubAssign(rhs I8) {
	*self -= rhs
}

func (self *I16) SubAssign(rhs I16) {
	*self -= rhs
}

func (self *I32) SubAssign(rhs I32) {
	*self -= rhs
}

func (self *I64) SubAssign(rhs I64) {
	*self -= rhs
}

func (self *USize) SubAssign(rhs USize) {
	*self -= rhs
}

func (self *U8) SubAssign(rhs U8) {
	*self -= rhs
}

func (self *U16) SubAssign(rhs U16) {
	*self -= rhs
}

func (self *U32) SubAssign(rhs U32) {
	*self -= rhs
}

func (self *U64) SubAssign(rhs U64) {
	*self -= rhs
}

func (self *F32) SubAssign(rhs F32) {
	*self -= rhs
}

func (self *F64) SubAssign(rhs F64) {
	*self -= rhs
}

func (self *Byte) SubAssign(rhs Byte) {
	*self -= rhs
}

func (self *Char) SubAssign(rhs Char) {
	*self -= rhs
}

func (self *Complex64) SubAssign(rhs Complex64) {
	*self -= rhs
}

func (self *Complex128) SubAssign(rhs Complex128) {
	*self -= rhs
}

// MulAssign implements
func (self *ISize) MulAssign(rhs ISize) {
	*self *= rhs
}

func (self *I8) MulAssign(rhs I8) {
	*self *= rhs
}

func (self *I16) MulAssign(rhs I16) {
	*self *= rhs
}

func (self *I32) MulAssign(rhs I32) {
	*self *= rhs
}

func (self *I64) MulAssign(rhs I64) {
	*self *= rhs
}

func (self *USize) MulAssign(rhs USize) {
	*self *= rhs
}

func (self *U8) MulAssign(rhs U8) {
	*self *= rhs
}

func (self *U16) MulAssign(rhs U16) {
	*self *= rhs
}

func (self *U32) MulAssign(rhs U32) {
	*self *= rhs
}

func (self *U64) MulAssign(rhs U64) {
	*self *= rhs
}

func (self *F32) MulAssign(rhs F32) {
	*self *= rhs
}

func (self *F64) MulAssign(rhs F64) {
	*self *= rhs
}

func (self *Byte) MulAssign(rhs Byte) {
	*self *= rhs
}

func (self *Char) MulAssign(rhs Char) {
	*self *= rhs
}

func (self *Complex64) MulAssign(rhs Complex64) {
	*self *= rhs
}

func (self *Complex128) MulAssign(rhs Complex128) {
	*self *= rhs
}

// DivAssign implements
func (self *ISize) DivAssign(rhs ISize) {
	*self /= rhs
}

func (self *I8) DivAssign(rhs I8) {
	*self /= rhs
}

func (self *I16) DivAssign(rhs I16) {
	*self /= rhs
}

func (self *I32) DivAssign(rhs I32) {
	*self /= rhs
}

func (self *I64) DivAssign(rhs I64) {
	*self /= rhs
}

func (self *USize) DivAssign(rhs USize) {
	*self /= rhs
}

func (self *U8) DivAssign(rhs U8) {
	*self /= rhs
}

func (self *U16) DivAssign(rhs U16) {
	*self /= rhs
}

func (self *U32) DivAssign(rhs U32) {
	*self /= rhs
}

func (self *U64) DivAssign(rhs U64) {
	*self /= rhs
}

func (self *F32) DivAssign(rhs F32) {
	*self /= rhs
}

func (self *F64) DivAssign(rhs F64) {
	*self /= rhs
}

func (self *Byte) DivAssign(rhs Byte) {
	*self /= rhs
}

func (self *Char) DivAssign(rhs Char) {
	*self /= rhs
}

func (self *Complex64) DivAssign(rhs Complex64) {
	*self /= rhs
}

func (self *Complex128) DivAssign(rhs Complex128) {
	*self /= rhs
}

// Wrapping (modular) addition. Computes self + rhs, wrapping around at the boundary of the type.
func (self ISize) WrappingAdd(rhs ISize) ISize {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - ISize(math.MaxInt) - 1
	}
	return result
}

func (self I8) WrappingAdd(rhs I8) I8 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I8(math.MaxInt8) - 1
	}
	return result
}

func (self I16) WrappingAdd(rhs I16) I16 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I16(math.MaxInt16) - 1
	}
	return result
}

func (self I32) WrappingAdd(rhs I32) I32 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I32(math.MaxInt32) - 1
	}
	return result
}

func (self I64) WrappingAdd(rhs I64) I64 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I64(math.MaxInt64) - 1
	}
	return result
}

func (self USize) WrappingAdd(rhs USize) USize {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - USize(math.MaxUint) - 1
	}
	return result
}

func (self U8) WrappingAdd(rhs U8) U8 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U8(math.MaxUint8) - 1
	}
	return result
}

func (self U16) WrappingAdd(rhs U16) U16 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U16(math.MaxUint16) - 1
	}
	return result
}

func (self U32) WrappingAdd(rhs U32) U32 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U32(math.MaxUint32) - 1
	}
	return result
}

func (self U64) WrappingAdd(rhs U64) U64 {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U64(math.MaxUint64) - 1
	}
	return result
}

// Wrapping (modular) subtraction. Computes self - rhs, wrapping around at the boundary of the type.
func (self ISize) WrappingSub(rhs ISize) ISize {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + ISize(math.MaxInt) + 1
	}
	return result
}

func (self I8) WrappingSub(rhs I8) I8 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I8(math.MaxInt8) + 1
	}
	return result
}

func (self I16) WrappingSub(rhs I16) I16 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I16(math.MaxInt16) + 1
	}
	return result
}

func (self I32) WrappingSub(rhs I32) I32 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I32(math.MaxInt32) + 1
	}
	return result
}

func (self I64) WrappingSub(rhs I64) I64 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I64(math.MaxInt64) + 1
	}
	return result
}

func (self USize) WrappingSub(rhs USize) USize {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + USize(math.MaxUint) + 1
	}
	return result
}

func (self U8) WrappingSub(rhs U8) U8 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U8(math.MaxUint8) + 1
	}
	return result
}

func (self U16) WrappingSub(rhs U16) U16 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U16(math.MaxUint16) + 1
	}
	return result
}

func (self U32) WrappingSub(rhs U32) U32 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U32(math.MaxUint32) + 1
	}
	return result
}

func (self U64) WrappingSub(rhs U64) U64 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U64(math.MaxUint64) + 1
	}
	return result
}
