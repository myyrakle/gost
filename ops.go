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

// Rem is a trait for types that support remainder after division.
type Rem[T any] interface {
	Rem(rhs T) T
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

// The remainder assignment operator %=.
type RemAssign[T any] interface {
	RemAssign(rhs T)
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

// Rem implements
func (self ISize) Rem(rhs ISize) ISize {
	return self % rhs
}

func (self I8) Rem(rhs I8) I8 {
	return self % rhs
}

func (self I16) Rem(rhs I16) I16 {
	return self % rhs
}

func (self I32) Rem(rhs I32) I32 {
	return self % rhs
}

func (self I64) Rem(rhs I64) I64 {
	return self % rhs
}

func (self USize) Rem(rhs USize) USize {
	return self % rhs
}

func (self U8) Rem(rhs U8) U8 {
	return self % rhs
}

func (self U16) Rem(rhs U16) U16 {
	return self % rhs
}

func (self U32) Rem(rhs U32) U32 {
	return self % rhs
}

func (self U64) Rem(rhs U64) U64 {
	return self % rhs
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

// RemAssign implements
func (self *ISize) RemAssign(rhs ISize) {
	*self %= rhs
}

func (self *I8) RemAssign(rhs I8) {
	*self %= rhs
}

func (self *I16) RemAssign(rhs I16) {
	*self %= rhs
}

func (self *I32) RemAssign(rhs I32) {
	*self %= rhs
}

func (self *I64) RemAssign(rhs I64) {
	*self %= rhs
}

func (self *USize) RemAssign(rhs USize) {
	*self %= rhs
}

func (self *U8) RemAssign(rhs U8) {
	*self %= rhs
}

func (self *U16) RemAssign(rhs U16) {
	*self %= rhs
}

func (self *U32) RemAssign(rhs U32) {
	*self %= rhs
}

func (self *U64) RemAssign(rhs U64) {
	*self %= rhs
}

func (self ISize) _HasOverflow_Add(rhs ISize) bool {
	if self > 0 && rhs > math.MaxInt-self {
		return true
	}
	if self < 0 && rhs < math.MinInt-self {
		return true
	}
	return false
}

func (self I8) _HasOverflow_Add(rhs I8) bool {
	if self > 0 && rhs > math.MaxInt8-self {
		return true
	}
	if self < 0 && rhs < math.MinInt8-self {
		return true
	}
	return false
}

// Wrapping (modular) addition. Computes self + rhs, wrapping around at the boundary of the type.
func (self ISize) WrappingAdd(rhs ISize) ISize {
	result := self + rhs

	return result
}

func (self I8) WrappingAdd(rhs I8) I8 {
	result := self + rhs

	return result
}

func (self I16) WrappingAdd(rhs I16) I16 {
	result := self + rhs

	return result
}

func (self I32) WrappingAdd(rhs I32) I32 {
	result := self + rhs

	return result
}

func (self I64) WrappingAdd(rhs I64) I64 {
	result := self + rhs

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

// Wrapping (modular) multiplication. Computes self * rhs, wrapping around at the boundary of the type.
func (self ISize) WrappingMul(rhs ISize) ISize {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - ISize(math.MaxInt) - 1
	}
	return result
}

func (self I8) WrappingMul(rhs I8) I8 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I8(math.MaxInt8) - 1
	}
	return result
}

func (self I16) WrappingMul(rhs I16) I16 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I16(math.MaxInt16) - 1
	}
	return result
}

func (self I32) WrappingMul(rhs I32) I32 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I32(math.MaxInt32) - 1
	}
	return result
}

func (self I64) WrappingMul(rhs I64) I64 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - I64(math.MaxInt64) - 1
	}
	return result
}

func (self USize) WrappingMul(rhs USize) USize {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - USize(math.MaxUint) - 1
	}
	return result
}

func (self U8) WrappingMul(rhs U8) U8 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U8(math.MaxUint8) - 1
	}
	return result
}

func (self U16) WrappingMul(rhs U16) U16 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U16(math.MaxUint16) - 1
	}
	return result
}

func (self U32) WrappingMul(rhs U32) U32 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U32(math.MaxUint32) - 1
	}
	return result
}

func (self U64) WrappingMul(rhs U64) U64 {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred, wrap around
		result = result - U64(math.MaxUint64) - 1
	}
	return result
}

// Wrapping (modular) division. Computes self / rhs, wrapping around at the boundary of the type.
func (self ISize) WrappingDiv(rhs ISize) ISize {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + ISize(math.MaxInt) + 1
	}
	return result
}

func (self I8) WrappingDiv(rhs I8) I8 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I8(math.MaxInt8) + 1
	}
	return result
}

func (self I16) WrappingDiv(rhs I16) I16 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I16(math.MaxInt16) + 1
	}
	return result
}

func (self I32) WrappingDiv(rhs I32) I32 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I32(math.MaxInt32) + 1
	}
	return result
}

func (self I64) WrappingDiv(rhs I64) I64 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + I64(math.MaxInt64) + 1
	}
	return result
}

func (self USize) WrappingDiv(rhs USize) USize {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + USize(math.MaxUint) + 1
	}
	return result
}

func (self U8) WrappingDiv(rhs U8) U8 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U8(math.MaxUint8) + 1
	}
	return result
}

func (self U16) WrappingDiv(rhs U16) U16 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U16(math.MaxUint16) + 1
	}
	return result
}

func (self U32) WrappingDiv(rhs U32) U32 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U32(math.MaxUint32) + 1
	}
	return result
}

func (self U64) WrappingDiv(rhs U64) U64 {
	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred, wrap around
		result = result + U64(math.MaxUint64) + 1
	}
	return result
}

// Checked integer addition. Computes self + rhs, returning None if overflow occurred.
func (self ISize) CheckedAdd(rhs ISize) Option[ISize] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[ISize]()
	}
	return Some[ISize](result)
}

func (self I8) CheckedAdd(rhs I8) Option[I8] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I8]()
	}
	return Some[I8](result)
}

func (self I16) CheckedAdd(rhs I16) Option[I16] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I16]()
	}
	return Some[I16](result)
}

func (self I32) CheckedAdd(rhs I32) Option[I32] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I32]()
	}
	return Some[I32](result)
}

func (self I64) CheckedAdd(rhs I64) Option[I64] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I64]()
	}
	return Some[I64](result)
}

func (self USize) CheckedAdd(rhs USize) Option[USize] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[USize]()
	}
	return Some[USize](result)
}

func (self U8) CheckedAdd(rhs U8) Option[U8] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U8]()
	}
	return Some[U8](result)
}

func (self U16) CheckedAdd(rhs U16) Option[U16] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U16]()
	}
	return Some[U16](result)
}

func (self U32) CheckedAdd(rhs U32) Option[U32] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U32]()
	}
	return Some[U32](result)
}

func (self U64) CheckedAdd(rhs U64) Option[U64] {
	result := self + rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U64]()
	}
	return Some[U64](result)
}

// Checked integer subtraction. Computes self - rhs, returning None if overflow occurred.
func (self ISize) CheckedSub(rhs ISize) Option[ISize] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[ISize]()
	}
	return Some[ISize](result)
}

func (self I8) CheckedSub(rhs I8) Option[I8] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I8]()
	}
	return Some[I8](result)
}

func (self I16) CheckedSub(rhs I16) Option[I16] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I16]()
	}
	return Some[I16](result)
}

func (self I32) CheckedSub(rhs I32) Option[I32] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I32]()
	}
	return Some[I32](result)
}

func (self I64) CheckedSub(rhs I64) Option[I64] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I64]()
	}
	return Some[I64](result)
}

func (self USize) CheckedSub(rhs USize) Option[USize] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[USize]()
	}
	return Some[USize](result)
}

func (self U8) CheckedSub(rhs U8) Option[U8] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U8]()
	}
	return Some[U8](result)
}

func (self U16) CheckedSub(rhs U16) Option[U16] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U16]()
	}
	return Some[U16](result)
}

func (self U32) CheckedSub(rhs U32) Option[U32] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U32]()
	}
	return Some[U32](result)
}

func (self U64) CheckedSub(rhs U64) Option[U64] {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U64]()
	}
	return Some[U64](result)
}

// Checked integer multiplication. Computes self * rhs, returning None if overflow occurred.
func (self ISize) CheckedMul(rhs ISize) Option[ISize] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[ISize]()
	}
	return Some[ISize](result)
}

func (self I8) CheckedMul(rhs I8) Option[I8] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I8]()
	}
	return Some[I8](result)
}

func (self I16) CheckedMul(rhs I16) Option[I16] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I16]()
	}
	return Some[I16](result)
}

func (self I32) CheckedMul(rhs I32) Option[I32] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I32]()
	}
	return Some[I32](result)
}

func (self I64) CheckedMul(rhs I64) Option[I64] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[I64]()
	}
	return Some[I64](result)
}

func (self USize) CheckedMul(rhs USize) Option[USize] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[USize]()
	}
	return Some[USize](result)
}

func (self U8) CheckedMul(rhs U8) Option[U8] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U8]()
	}
	return Some[U8](result)
}

func (self U16) CheckedMul(rhs U16) Option[U16] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U16]()
	}
	return Some[U16](result)
}

func (self U32) CheckedMul(rhs U32) Option[U32] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U32]()
	}
	return Some[U32](result)
}

func (self U64) CheckedMul(rhs U64) Option[U64] {
	result := self * rhs

	if result < self || result < rhs {
		// Overflow occurred
		return None[U64]()
	}
	return Some[U64](result)
}

// Checked integer division. Computes self / rhs, returning None if rhs == 0 or the operation results in overflow.
func (self ISize) CheckedDiv(rhs ISize) Option[ISize] {
	if rhs == 0 {
		return None[ISize]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[ISize]()
	}
	return Some[ISize](result)
}

func (self I8) CheckedDiv(rhs I8) Option[I8] {
	if rhs == 0 {
		return None[I8]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I8]()
	}
	return Some[I8](result)
}

func (self I16) CheckedDiv(rhs I16) Option[I16] {
	if rhs == 0 {
		return None[I16]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I16]()
	}
	return Some[I16](result)
}

func (self I32) CheckedDiv(rhs I32) Option[I32] {
	if rhs == 0 {
		return None[I32]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I32]()
	}
	return Some[I32](result)
}

func (self I64) CheckedDiv(rhs I64) Option[I64] {
	if rhs == 0 {
		return None[I64]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I64]()
	}
	return Some[I64](result)
}

func (self USize) CheckedDiv(rhs USize) Option[USize] {
	if rhs == 0 {
		return None[USize]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[USize]()
	}
	return Some[USize](result)
}

func (self U8) CheckedDiv(rhs U8) Option[U8] {
	if rhs == 0 {
		return None[U8]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U8]()
	}
	return Some[U8](result)
}

func (self U16) CheckedDiv(rhs U16) Option[U16] {
	if rhs == 0 {
		return None[U16]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U16]()
	}
	return Some[U16](result)
}

func (self U32) CheckedDiv(rhs U32) Option[U32] {
	if rhs == 0 {
		return None[U32]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U32]()
	}
	return Some[U32](result)
}

func (self U64) CheckedDiv(rhs U64) Option[U64] {
	if rhs == 0 {
		return None[U64]()
	}

	result := self / rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U64]()
	}
	return Some[U64](result)
}

// Checked integer remainder. Computes self % rhs, returning None if rhs == 0 or the operation results in overflow.
func (self ISize) CheckedRem(rhs ISize) Option[ISize] {
	if rhs == 0 {
		return None[ISize]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[ISize]()
	}
	return Some[ISize](result)
}

func (self I8) CheckedRem(rhs I8) Option[I8] {
	if rhs == 0 {
		return None[I8]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I8]()
	}
	return Some[I8](result)
}

func (self I16) CheckedRem(rhs I16) Option[I16] {
	if rhs == 0 {
		return None[I16]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I16]()
	}
	return Some[I16](result)
}

func (self I32) CheckedRem(rhs I32) Option[I32] {
	if rhs == 0 {
		return None[I32]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I32]()
	}
	return Some[I32](result)
}

func (self I64) CheckedRem(rhs I64) Option[I64] {
	if rhs == 0 {
		return None[I64]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[I64]()
	}
	return Some[I64](result)
}

func (self USize) CheckedRem(rhs USize) Option[USize] {
	if rhs == 0 {
		return None[USize]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[USize]()
	}
	return Some[USize](result)
}

func (self U8) CheckedRem(rhs U8) Option[U8] {
	if rhs == 0 {
		return None[U8]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U8]()
	}
	return Some[U8](result)
}

func (self U16) CheckedRem(rhs U16) Option[U16] {
	if rhs == 0 {
		return None[U16]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U16]()
	}
	return Some[U16](result)
}

func (self U32) CheckedRem(rhs U32) Option[U32] {
	if rhs == 0 {
		return None[U32]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U32]()
	}
	return Some[U32](result)
}

func (self U64) CheckedRem(rhs U64) Option[U64] {
	if rhs == 0 {
		return None[U64]()
	}

	result := self % rhs

	if result > self || result > rhs {
		// Overflow occurred
		return None[U64]()
	}
	return Some[U64](result)
}

// Saturating integer subtraction. Computes self - rhs, saturating at the numeric bounds instead of overflowing.
func (self ISize) SaturatingSub(rhs ISize) ISize {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		if self < 0 {
			return ISize(math.MinInt)
		}
		return ISize(math.MaxInt)
	}
	return result
}

func (self I8) SaturatingSub(rhs I8) I8 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		if self < 0 {
			return I8(math.MinInt8)
		}
		return I8(math.MaxInt8)
	}
	return result
}

func (self I16) SaturatingSub(rhs I16) I16 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		if self < 0 {
			return I16(math.MinInt16)
		}
		return I16(math.MaxInt16)
	}
	return result
}

func (self I32) SaturatingSub(rhs I32) I32 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		if self < 0 {
			return I32(math.MinInt32)
		}
		return I32(math.MaxInt32)
	}
	return result
}

func (self I64) SaturatingSub(rhs I64) I64 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		if self < 0 {
			return I64(math.MinInt64)
		}
		return I64(math.MaxInt64)
	}
	return result
}

func (self USize) SaturatingSub(rhs USize) USize {
	if self < rhs {
		return 0
	}
	return self - rhs
}

func (self U8) SaturatingSub(rhs U8) U8 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		return U8(math.MaxUint8)
	}
	return result
}

func (self U16) SaturatingSub(rhs U16) U16 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		return U16(math.MaxUint16)
	}
	return result
}

func (self U32) SaturatingSub(rhs U32) U32 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		return U32(math.MaxUint32)
	}
	return result
}

func (self U64) SaturatingSub(rhs U64) U64 {
	result := self - rhs

	if result > self || result > rhs {
		// Overflow occurred, saturate
		return U64(math.MaxUint64)
	}
	return result
}
