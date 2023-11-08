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

func Test_CheckedAdd_UnsignedInteger(t *testing.T) {
	AssertEq(USize(100).CheckedAdd(27), Some(USize(127)), "ISize.CheckedAdd")
	AssertEq(USize(math.MaxUint).CheckedAdd(2), None[USize](), "ISize.CheckedAdd overflow")

	AssertEq(U8(100).CheckedAdd(27), Some(U8(127)), "U8.CheckedAdd")
	AssertEq(U8(math.MaxUint8).CheckedAdd(2), None[U8](), "U8.CheckedAdd overflow")

	AssertEq(U16(100).CheckedAdd(27), Some(U16(127)), "U16.CheckedAdd")
	AssertEq(U16(math.MaxUint16).CheckedAdd(2), None[U16](), "U16.CheckedAdd overflow")

	AssertEq(U32(100).CheckedAdd(27), Some(U32(127)), "U32.CheckedAdd")
	AssertEq(U32(math.MaxUint32).CheckedAdd(2), None[U32](), "U32.CheckedAdd overflow")

	AssertEq(U64(100).CheckedAdd(27), Some(U64(127)), "U64.CheckedAdd")
	AssertEq(U64(math.MaxUint64).CheckedAdd(2), None[U64](), "U64.CheckedAdd overflow")
}

func Test_CheckedSub_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).CheckedSub(27), Some(ISize(73)), "ISize.CheckedSub")
	AssertEq(ISize(math.MinInt).CheckedSub(2), None[ISize](), "ISize.CheckedSub underflow")
	AssertEq(ISize(math.MaxInt).CheckedSub(-2), None[ISize](), "ISize.CheckedSub overflow")

	AssertEq(I8(100).CheckedSub(27), Some(I8(73)), "I8.CheckedSub")
	AssertEq(I8(math.MinInt8).CheckedSub(2), None[I8](), "I8.CheckedSub underflow")
	AssertEq(I8(math.MaxInt8).CheckedSub(-2), None[I8](), "I8.CheckedSub overflow")

	AssertEq(I16(100).CheckedSub(27), Some(I16(73)), "I16.CheckedSub")
	AssertEq(I16(math.MinInt16).CheckedSub(2), None[I16](), "I16.CheckedSub underflow")
	AssertEq(I16(math.MaxInt16).CheckedSub(-2), None[I16](), "I16.CheckedSub overflow")

	AssertEq(I32(100).CheckedSub(27), Some(I32(73)), "I32.CheckedSub")
	AssertEq(I32(math.MinInt32).CheckedSub(2), None[I32](), "I32.CheckedSub underflow")
	AssertEq(I32(math.MaxInt32).CheckedSub(-2), None[I32](), "I32.CheckedSub overflow")

	AssertEq(I64(100).CheckedSub(27), Some(I64(73)), "I64.CheckedSub")
	AssertEq(I64(math.MinInt64).CheckedSub(2), None[I64](), "I64.CheckedSub underflow")
	AssertEq(I64(math.MaxInt64).CheckedSub(-2), None[I64](), "I64.CheckedSub overflow")
}

func Test_CheckedSub_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).CheckedSub(27), Some(USize(73)), "USize.CheckedSub")
	AssertEq(USize(0).CheckedSub(2), None[USize](), "USize.CheckedSub underflow")

	AssertEq(U8(100).CheckedSub(27), Some(U8(73)), "U8.CheckedSub")
	AssertEq(U8(0).CheckedSub(2), None[U8](), "U8.CheckedSub underflow")

	AssertEq(U16(100).CheckedSub(27), Some(U16(73)), "U16.CheckedSub")
	AssertEq(U16(0).CheckedSub(2), None[U16](), "U16.CheckedSub underflow")

	AssertEq(U32(100).CheckedSub(27), Some(U32(73)), "U32.CheckedSub")
	AssertEq(U32(0).CheckedSub(2), None[U32](), "U32.CheckedSub underflow")

	AssertEq(U64(100).CheckedSub(27), Some(U64(73)), "U64.CheckedSub")
	AssertEq(U64(0).CheckedSub(2), None[U64](), "U64.CheckedSub underflow")
}

func Test_CheckedMul_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(10).CheckedMul(12), Some(ISize(120)), "ISize.CheckedMul")
	AssertEq(ISize(math.MaxInt).CheckedMul(2), None[ISize](), "ISize.CheckedMul overflow")
	AssertEq(ISize(math.MinInt).CheckedMul(2), None[ISize](), "ISize.CheckedMul underflow")

	AssertEq(I8(11).CheckedMul(5), Some(I8(55)), "I8.CheckedMul")
	AssertEq(I8(math.MaxInt8).CheckedMul(2), None[I8](), "I8.CheckedMul overflow")
	AssertEq(I8(math.MinInt8).CheckedMul(2), None[I8](), "I8.CheckedMul underflow")

	AssertEq(I16(11).CheckedMul(12), Some(I16(132)), "I16.CheckedMul")
	AssertEq(I16(math.MaxInt16).CheckedMul(2), None[I16](), "I16.CheckedMul overflow")
	AssertEq(I16(math.MinInt16).CheckedMul(2), None[I16](), "I16.CheckedMul underflow")

	AssertEq(I32(11).CheckedMul(12), Some(I32(132)), "I32.CheckedMul")
	AssertEq(I32(math.MaxInt32).CheckedMul(2), None[I32](), "I32.CheckedMul overflow")
	AssertEq(I32(math.MinInt32).CheckedMul(2), None[I32](), "I32.CheckedMul underflow")

	AssertEq(I64(11).CheckedMul(12), Some(I64(132)), "I64.CheckedMul")
	AssertEq(I64(math.MaxInt64).CheckedMul(2), None[I64](), "I64.CheckedMul overflow")
	AssertEq(I64(math.MinInt64).CheckedMul(2), None[I64](), "I64.CheckedMul underflow")
}

func Test_CheckedMul_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(10).CheckedMul(12), Some(USize(120)), "USize.CheckedMul")
	AssertEq(USize(math.MaxUint).CheckedMul(2), None[USize](), "USize.CheckedMul overflow")

	AssertEq(U8(11).CheckedMul(5), Some(U8(55)), "U8.CheckedMul")
	AssertEq(U8(math.MaxUint8).CheckedMul(2), None[U8](), "U8.CheckedMul overflow")

	AssertEq(U16(11).CheckedMul(12), Some(U16(132)), "U16.CheckedMul")
	AssertEq(U16(math.MaxUint16).CheckedMul(2), None[U16](), "U16.CheckedMul overflow")

	AssertEq(U32(11).CheckedMul(12), Some(U32(132)), "U32.CheckedMul")
	AssertEq(U32(math.MaxUint32).CheckedMul(2), None[U32](), "U32.CheckedMul overflow")

	AssertEq(U64(11).CheckedMul(12), Some(U64(132)), "U64.CheckedMul")
	AssertEq(U64(math.MaxUint64).CheckedMul(2), None[U64](), "U64.CheckedMul overflow")
}
