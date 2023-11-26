package gost

type Unit struct{}
type Void struct{}

type ISize int
type I8 int8
type I16 int16
type I32 int32
type I64 int64
type I128 struct {
	high I64
	low  U64
}

type USize uint
type U8 uint8
type U16 uint16
type U32 uint32
type U64 uint64
type U128 struct {
	high U64
	low  U64
}

type F32 float32
type F64 float64

type Byte byte

type Char rune

type String string

type Bool bool

type Complex64 complex64
type Complex128 complex128

type Error error

type Any interface{}

func U128_FromU64(low U64) U128 {
	return U128{
		high: 0,
		low:  low,
	}
}

func I128_FromI64(low I64) I128 {
	isNegative := low < 0

	if low == 0 {
		return I128{
			high: 0,
			low:  0,
		}
	}

	if isNegative {
		return I128{
			high: 0,
			low:  U64(low.Abs()),
		}.Neg()
	} else {
		return I128{
			high: 0,
			low:  U64(low),
		}
	}
}

func I128_FromU64(low U64) I128 {
	return I128{
		high: 0,
		low:  low,
	}
}
