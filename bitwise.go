package gost

type Shl[T any] interface {
	Shl(rhs T) T
}

type Shr[T any] interface {
	Shr(rhs T) T
}

func (lhs ISize) Shl(rhs ISize) ISize {
	return ISize(lhs << lhs)
}

func (lhs I8) Shl(rhs I8) I8 {
	return I8(lhs << lhs)
}

func (lhs I16) Shl(rhs I16) I16 {
	return I16(lhs << lhs)
}

func (lhs I32) Shl(rhs I32) I32 {
	return I32(lhs << lhs)
}

func (lhs I64) Shl(rhs I64) I64 {
	return I64(lhs << lhs)
}

func (lhs I128) Shl(rhs I128) I128 {
	hasCarry := false

	low := lhs.low << rhs.low
	high := lhs.high << rhs.low

	// check if there is a carry
	if lhs.low>>63 == 1 {
		hasCarry = true
	}

	if hasCarry {
		high += 1
	}

	return I128{
		high: high,
		low:  low,
	}
}

func (lhs USize) Shl(rhs USize) USize {
	return USize(lhs << lhs)
}

func (lhs U8) Shl(rhs U8) U8 {
	return U8(lhs << lhs)
}

func (lhs U16) Shl(rhs U16) U16 {
	return U16(lhs << lhs)
}

func (lhs U32) Shl(rhs U32) U32 {
	return U32(lhs << lhs)
}

func (lhs U64) Shl(rhs U64) U64 {
	return U64(lhs << lhs)
}

func (lhs U128) Shl(rhs U128) U128 {
	hasCarry := false

	low := lhs.low << rhs.low
	high := lhs.high << rhs.low

	// check if there is a carry
	if lhs.low>>63 == 1 {
		hasCarry = true
	}

	if hasCarry {
		high += 1
	}

	return U128{
		high: high,
		low:  low,
	}
}

func (lhs ISize) Shr(rhs ISize) ISize {
	return ISize(lhs >> lhs)
}

func (lhs I8) Shr(rhs I8) I8 {
	return I8(lhs >> lhs)
}

func (lhs I16) Shr(rhs I16) I16 {
	return I16(lhs >> lhs)
}

func (lhs I32) Shr(rhs I32) I32 {
	return I32(lhs >> lhs)
}

func (lhs I64) Shr(rhs I64) I64 {
	return I64(lhs >> lhs)
}

func (lhs I128) Shr(rhs I128) I128 {
	hasCarry := false

	low := lhs.low >> rhs.low
	high := lhs.high >> rhs.low

	// check if there is a carry
	if lhs.high&1 == 1 {
		hasCarry = true
	}

	if hasCarry {
		low |= U64(uint64(1) << 63)
	}

	return I128{
		high: high,
		low:  low,
	}
}

func (lhs USize) Shr(rhs USize) USize {
	return USize(lhs >> lhs)
}

func (lhs U8) Shr(rhs U8) U8 {
	return U8(lhs >> lhs)
}

func (lhs U16) Shr(rhs U16) U16 {
	return U16(lhs >> lhs)
}

func (lhs U32) Shr(rhs U32) U32 {
	return U32(lhs >> lhs)
}

func (lhs U64) Shr(rhs U64) U64 {
	return U64(lhs >> lhs)
}
