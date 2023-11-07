package gost

import (
	"math"
	"testing"
)

func Test_WrappingAdd_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).WrappingAdd(27), ISize(127), "ISize.WrappingAdd")
	AssertEq(ISize(math.MaxInt).WrappingAdd(2), ISize(math.MinInt+1), "ISize.WrappingAdd overflow")

	AssertEq(I8(100).WrappingAdd(27), I8(127), "I8.WrappingAdd")
	AssertEq(I8(math.MaxInt8).WrappingAdd(2), I8(math.MinInt8+1), "I8.WrappingAdd: overflow")

	AssertEq(I16(100).WrappingAdd(27), I16(127), "I16.WrappingAdd")
	AssertEq(I16(math.MaxInt16).WrappingAdd(2), I16(math.MinInt16+1), "I16.WrappingAdd: overflow")

	AssertEq(I32(100).WrappingAdd(27), I32(127), "I32.WrappingAdd")
	AssertEq(I32(math.MaxInt32).WrappingAdd(2), I32(math.MinInt32+1), "I32.WrappingAdd: overflow")

	AssertEq(I64(100).WrappingAdd(27), I64(127), "I64.WrappingAdd")
	AssertEq(I64(math.MaxInt64).WrappingAdd(2), I64(math.MinInt64+1), "I64.WrappingAdd: overflow")

}

func Test_WrappingAdd_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(200).WrappingAdd(55), USize(255), "USize.WrappingAdd")
	AssertEq(USize(math.MaxUint).WrappingAdd(200), USize(199), "USize.WrappingAdd: overflow")

	AssertEq(U8(200).WrappingAdd(55), U8(255), "U8.WrappingAdd")
	AssertEq(U8(math.MaxUint8).WrappingAdd(200), U8(199), "U8.WrappingAdd: overflow")

	AssertEq(U16(200).WrappingAdd(55), U16(255), "U16.WrappingAdd")
	AssertEq(U16(math.MaxUint16).WrappingAdd(200), U16(199), "U16.WrappingAdd: overflow")

	AssertEq(U32(200).WrappingAdd(55), U32(255), "U32.WrappingAdd")
	AssertEq(U32(math.MaxUint32).WrappingAdd(200), U32(199), "U32.WrappingAdd: overflow")

	AssertEq(U64(200).WrappingAdd(55), U64(255), "U64.WrappingAdd")
	AssertEq(U64(math.MaxUint64).WrappingAdd(200), U64(199), "U64.WrappingAdd: overflow")
}

func Test_WrappingSub_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(0).WrappingSub(127), ISize(-127), "ISize.WrappingSub")
	AssertEq(ISize(-2).WrappingSub(math.MaxInt), ISize(math.MaxInt), "ISize.WrappingSub overflow")

	AssertEq(I8(0).WrappingSub(127), I8(-127), "I8.WrappingSub")
	AssertEq(I8(-2).WrappingSub(math.MaxInt8), I8(math.MaxInt8), "I8.WrappingSub overflow")

	AssertEq(I16(0).WrappingSub(127), I16(-127), "I16.WrappingSub")
	AssertEq(I16(-2).WrappingSub(math.MaxInt16), I16(math.MaxInt16), "I16.WrappingSub overflow")

	AssertEq(I32(0).WrappingSub(127), I32(-127), "I32.WrappingSub")
	AssertEq(I32(-2).WrappingSub(math.MaxInt32), I32(math.MaxInt32), "I32.WrappingSub overflow")

	AssertEq(I64(0).WrappingSub(127), I64(-127), "I64.WrappingSub")
	AssertEq(I64(-2).WrappingSub(math.MaxInt64), I64(math.MaxInt64), "I64.WrappingSub overflow")
}

func Test_WrappingSub_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).WrappingSub(100), USize(0), "USize.WrappingSub")
	AssertEq(USize(100).WrappingSub(math.MaxUint), USize(101), "USize.WrappingSub overflow")

	AssertEq(U8(100).WrappingSub(100), U8(0), "U8.WrappingSub")
	AssertEq(U8(100).WrappingSub(math.MaxUint8), U8(101), "U8.WrappingSub overflow")

	AssertEq(U16(100).WrappingSub(100), U16(0), "U16.WrappingSub")
	AssertEq(U16(100).WrappingSub(math.MaxUint16), U16(101), "U16.WrappingSub overflow")

	AssertEq(U32(100).WrappingSub(100), U32(0), "U32.WrappingSub")
	AssertEq(U32(100).WrappingSub(math.MaxUint32), U32(101), "U32.WrappingSub overflow")

	AssertEq(U64(100).WrappingSub(100), U64(0), "U64.WrappingSub")
	AssertEq(U64(100).WrappingSub(math.MaxUint64), U64(101), "U64.WrappingSub overflow")
}

func Test_WrappingMul_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(10).WrappingMul(12), ISize(120), "ISize.WrappingMul")
	AssertEq(I8(11).WrappingMul(12), I8(-124), "I8.WrappingMul overflow")
}

func Test_WrappingMul_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(10).WrappingMul(12), USize(120), "USize.WrappingMul")
	AssertEq(U8(25).WrappingMul(12), U8(44), "U8.WrappingMul overflow")
}

func Test_WrappingDiv_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).WrappingDiv(10), ISize(10), "ISize.WrappingDiv")
	AssertEq(I8(-128).WrappingDiv(-1), I8(-128), "I8.WrappingDiv overflow")
}

func Test_WrappingDiv_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).WrappingDiv(10), USize(10), "USize.WrappingDiv")
}

func Test_CheckedAdd_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).CheckedAdd(27), Some(ISize(127)), "ISize.CheckedAdd")
	AssertEq(ISize(math.MaxInt).CheckedAdd(2), None[ISize](), "ISize.CheckedAdd overflow")
	AssertEq(ISize(math.MinInt).CheckedAdd(-2), None[ISize](), "ISize.CheckedAdd underflow")

	AssertEq(I8(100).CheckedAdd(27), Some(I8(127)), "I8.CheckedAdd")
	AssertEq(I8(math.MaxInt8).CheckedAdd(2), None[I8](), "I8.CheckedAdd overflow")
	AssertEq(I8(math.MinInt8).CheckedAdd(-2), None[I8](), "I8.CheckedAdd underflow")

	AssertEq(I16(100).CheckedAdd(27), Some(I16(127)), "I16.CheckedAdd")
	AssertEq(I16(math.MaxInt16).CheckedAdd(2), None[I16](), "I16.CheckedAdd overflow")
	AssertEq(I16(math.MinInt16).CheckedAdd(-2), None[I16](), "I16.CheckedAdd underflow")

	AssertEq(I32(100).CheckedAdd(27), Some(I32(127)), "I32.CheckedAdd")
	AssertEq(I32(math.MaxInt32).CheckedAdd(2), None[I32](), "I32.CheckedAdd overflow")
	AssertEq(I32(math.MinInt32).CheckedAdd(-2), None[I32](), "I32.CheckedAdd underflow")

	AssertEq(I64(100).CheckedAdd(27), Some(I64(127)), "I64.CheckedAdd")
	AssertEq(I64(math.MaxInt64).CheckedAdd(2), None[I64](), "I64.CheckedAdd overflow")
	AssertEq(I64(math.MinInt64).CheckedAdd(-2), None[I64](), "I64.CheckedAdd underflow")

}
