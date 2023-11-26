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

func Test_CheckedDiv_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).CheckedDiv(10), Some(ISize(10)), "ISize.CheckedDiv")
	AssertEq(ISize(math.MinInt).CheckedDiv(0), None[ISize](), "ISize.CheckedDiv divide by zero")
	AssertEq(ISize(math.MinInt).CheckedDiv(-1), None[ISize](), "ISize.CheckedDiv overflow")

	AssertEq(I8(100).CheckedDiv(10), Some(I8(10)), "I8.CheckedDiv")
	AssertEq(I8(math.MinInt8).CheckedDiv(0), None[I8](), "I8.CheckedDiv divide by zero")
	AssertEq(I8(math.MinInt8).CheckedDiv(-1), None[I8](), "I8.CheckedDiv overflow")

	AssertEq(I16(100).CheckedDiv(10), Some(I16(10)), "I16.CheckedDiv")
	AssertEq(I16(math.MinInt16).CheckedDiv(0), None[I16](), "I16.CheckedDiv divide by zero")
	AssertEq(I16(math.MinInt16).CheckedDiv(-1), None[I16](), "I16.CheckedDiv overflow")

	AssertEq(I32(100).CheckedDiv(10), Some(I32(10)), "I32.CheckedDiv")
	AssertEq(I32(math.MinInt32).CheckedDiv(0), None[I32](), "I32.CheckedDiv divide by zero")
	AssertEq(I32(math.MinInt32).CheckedDiv(-1), None[I32](), "I32.CheckedDiv overflow")

	AssertEq(I64(100).CheckedDiv(10), Some(I64(10)), "I64.CheckedDiv")
	AssertEq(I64(math.MinInt64).CheckedDiv(0), None[I64](), "I64.CheckedDiv divide by zero")
	AssertEq(I64(math.MinInt64).CheckedDiv(-1), None[I64](), "I64.CheckedDiv overflow")
}

func Test_CheckedDiv_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).CheckedDiv(10), Some(USize(10)), "USize.CheckedDiv")
	AssertEq(USize(100).CheckedDiv(0), None[USize](), "USize.CheckedDiv divide by zero")

	AssertEq(U8(100).CheckedDiv(10), Some(U8(10)), "U8.CheckedDiv")
	AssertEq(U8(100).CheckedDiv(0), None[U8](), "U8.CheckedDiv divide by zero")

	AssertEq(U16(100).CheckedDiv(10), Some(U16(10)), "U16.CheckedDiv")
	AssertEq(U16(100).CheckedDiv(0), None[U16](), "U16.CheckedDiv divide by zero")

	AssertEq(U32(100).CheckedDiv(10), Some(U32(10)), "U32.CheckedDiv")
	AssertEq(U32(100).CheckedDiv(0), None[U32](), "U32.CheckedDiv divide by zero")

	AssertEq(U64(100).CheckedDiv(10), Some(U64(10)), "U64.CheckedDiv")
	AssertEq(U64(100).CheckedDiv(0), None[U64](), "U64.CheckedDiv divide by zero")
}

func Test_Pow_Integer(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(2).Pow(3), ISize(8), "ISize.Pow")
	AssertEq(I8(2).Pow(3), I8(8), "I8.Pow")
	AssertEq(I16(2).Pow(3), I16(8), "I16.Pow")
	AssertEq(I32(2).Pow(3), I32(8), "I32.Pow")
	AssertEq(I64(2).Pow(3), I64(8), "I64.Pow")

	AssertEq(USize(2).Pow(3), USize(8), "USize.Pow")
	AssertEq(U8(2).Pow(3), U8(8), "U8.Pow")
	AssertEq(U16(2).Pow(3), U16(8), "U16.Pow")
	AssertEq(U32(2).Pow(3), U32(8), "U32.Pow")
	AssertEq(U64(2).Pow(3), U64(8), "U64.Pow")
}

func Test_SaturatingAdd_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).SaturatingAdd(27), ISize(127), "ISize.SaturatingAdd")
	AssertEq(ISize(math.MaxInt).SaturatingAdd(2), ISize(math.MaxInt), "ISize.SaturatingAdd overflow")
	AssertEq(ISize(math.MinInt).SaturatingAdd(-2), ISize(math.MinInt), "ISize.SaturatingAdd underflow")

	AssertEq(I8(100).SaturatingAdd(27), I8(127), "I8.SaturatingAdd")
	AssertEq(I8(math.MaxInt8).SaturatingAdd(2), I8(math.MaxInt8), "I8.SaturatingAdd overflow")
	AssertEq(I8(math.MinInt8).SaturatingAdd(-2), I8(math.MinInt8), "I8.SaturatingAdd underflow")

	AssertEq(I16(100).SaturatingAdd(27), I16(127), "I16.SaturatingAdd")
	AssertEq(I16(math.MaxInt16).SaturatingAdd(2), I16(math.MaxInt16), "I16.SaturatingAdd overflow")
	AssertEq(I16(math.MinInt16).SaturatingAdd(-2), I16(math.MinInt16), "I16.SaturatingAdd underflow")

	AssertEq(I32(100).SaturatingAdd(27), I32(127), "I32.SaturatingAdd")
	AssertEq(I32(math.MaxInt32).SaturatingAdd(2), I32(math.MaxInt32), "I32.SaturatingAdd overflow")
	AssertEq(I32(math.MinInt32).SaturatingAdd(-2), I32(math.MinInt32), "I32.SaturatingAdd underflow")

	AssertEq(I64(100).SaturatingAdd(27), I64(127), "I64.SaturatingAdd")
	AssertEq(I64(math.MaxInt64).SaturatingAdd(2), I64(math.MaxInt64), "I64.SaturatingAdd overflow")
	AssertEq(I64(math.MinInt64).SaturatingAdd(-2), I64(math.MinInt64), "I64.SaturatingAdd underflow")
}

func Test_SaturatingAdd_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).SaturatingAdd(27), USize(127), "USize.SaturatingAdd")
	AssertEq(USize(math.MaxUint).SaturatingAdd(200), USize(math.MaxUint), "USize.SaturatingAdd overflow")

	AssertEq(U8(100).SaturatingAdd(27), U8(127), "U8.SaturatingAdd")
	AssertEq(U8(math.MaxUint8).SaturatingAdd(200), U8(math.MaxUint8), "U8.SaturatingAdd overflow")

	AssertEq(U16(100).SaturatingAdd(27), U16(127), "U16.SaturatingAdd")
	AssertEq(U16(math.MaxUint16).SaturatingAdd(200), U16(math.MaxUint16), "U16.SaturatingAdd overflow")

	AssertEq(U32(100).SaturatingAdd(27), U32(127), "U32.SaturatingAdd")
	AssertEq(U32(math.MaxUint32).SaturatingAdd(200), U32(math.MaxUint32), "U32.SaturatingAdd overflow")

	AssertEq(U64(100).SaturatingAdd(27), U64(127), "U64.SaturatingAdd")
	AssertEq(U64(math.MaxUint64).SaturatingAdd(200), U64(math.MaxUint64), "U64.SaturatingAdd overflow")
}

func Test_SaturatingSub_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).SaturatingSub(27), ISize(73), "ISize.SaturatingSub")
	AssertEq(ISize(math.MinInt).SaturatingSub(2), ISize(math.MinInt), "ISize.SaturatingSub underflow")
	AssertEq(ISize(math.MaxInt).SaturatingSub(-2), ISize(math.MaxInt), "ISize.SaturatingSub overflow")

	AssertEq(I8(100).SaturatingSub(27), I8(73), "I8.SaturatingSub")
	AssertEq(I8(math.MinInt8).SaturatingSub(2), I8(math.MinInt8), "I8.SaturatingSub underflow")
	AssertEq(I8(math.MaxInt8).SaturatingSub(-2), I8(math.MaxInt8), "I8.SaturatingSub overflow")

	AssertEq(I16(100).SaturatingSub(27), I16(73), "I16.SaturatingSub")
	AssertEq(I16(math.MinInt16).SaturatingSub(2), I16(math.MinInt16), "I16.SaturatingSub underflow")
	AssertEq(I16(math.MaxInt16).SaturatingSub(-2), I16(math.MaxInt16), "I16.SaturatingSub overflow")

	AssertEq(I32(100).SaturatingSub(27), I32(73), "I32.SaturatingSub")
	AssertEq(I32(math.MinInt32).SaturatingSub(2), I32(math.MinInt32), "I32.SaturatingSub underflow")
	AssertEq(I32(math.MaxInt32).SaturatingSub(-2), I32(math.MaxInt32), "I32.SaturatingSub overflow")

	AssertEq(I64(100).SaturatingSub(27), I64(73), "I64.SaturatingSub")
	AssertEq(I64(math.MinInt64).SaturatingSub(2), I64(math.MinInt64), "I64.SaturatingSub underflow")
	AssertEq(I64(math.MaxInt64).SaturatingSub(-2), I64(math.MaxInt64), "I64.SaturatingSub overflow")
}

func Test_SaturatingSub_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).SaturatingSub(27), USize(73), "USize.SaturatingSub")
	AssertEq(USize(0).SaturatingSub(2), USize(0), "USize.SaturatingSub underflow")

	AssertEq(U8(100).SaturatingSub(27), U8(73), "U8.SaturatingSub")
	AssertEq(U8(0).SaturatingSub(2), U8(0), "U8.SaturatingSub underflow")

	AssertEq(U16(100).SaturatingSub(27), U16(73), "U16.SaturatingSub")
	AssertEq(U16(0).SaturatingSub(2), U16(0), "U16.SaturatingSub underflow")

	AssertEq(U32(100).SaturatingSub(27), U32(73), "U32.SaturatingSub")
	AssertEq(U32(0).SaturatingSub(2), U32(0), "U32.SaturatingSub underflow")

	AssertEq(U64(100).SaturatingSub(27), U64(73), "U64.SaturatingSub")
	AssertEq(U64(0).SaturatingSub(2), U64(0), "U64.SaturatingSub underflow")
}

func Test_SaturatingMul_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(10).SaturatingMul(12), ISize(120), "ISize.SaturatingMul")
	AssertEq(ISize(math.MaxInt).SaturatingMul(2), ISize(math.MaxInt), "ISize.SaturatingMul overflow")
	AssertEq(ISize(math.MinInt).SaturatingMul(2), ISize(math.MinInt), "ISize.SaturatingMul underflow")

	AssertEq(I8(11).SaturatingMul(5), I8(55), "I8.SaturatingMul")
	AssertEq(I8(math.MaxInt8).SaturatingMul(2), I8(math.MaxInt8), "I8.SaturatingMul overflow")
	AssertEq(I8(math.MinInt8).SaturatingMul(2), I8(math.MinInt8), "I8.SaturatingMul underflow")

	AssertEq(I16(11).SaturatingMul(12), I16(132), "I16.SaturatingMul")
	AssertEq(I16(math.MaxInt16).SaturatingMul(2), I16(math.MaxInt16), "I16.SaturatingMul overflow")
	AssertEq(I16(math.MinInt16).SaturatingMul(2), I16(math.MinInt16), "I16.SaturatingMul underflow")

	AssertEq(I32(11).SaturatingMul(12), I32(132), "I32.SaturatingMul")
	AssertEq(I32(math.MaxInt32).SaturatingMul(2), I32(math.MaxInt32), "I32.SaturatingMul overflow")
	AssertEq(I32(math.MinInt32).SaturatingMul(2), I32(math.MinInt32), "I32.SaturatingMul underflow")

	AssertEq(I64(11).SaturatingMul(12), I64(132), "I64.SaturatingMul")
	AssertEq(I64(math.MaxInt64).SaturatingMul(2), I64(math.MaxInt64), "I64.SaturatingMul overflow")
	AssertEq(I64(math.MinInt64).SaturatingMul(2), I64(math.MinInt64), "I64.SaturatingMul underflow")
}

func Test_SaturatingMul_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(10).SaturatingMul(12), USize(120), "USize.SaturatingMul")
	AssertEq(USize(math.MaxUint).SaturatingMul(2), USize(math.MaxUint), "USize.SaturatingMul overflow")

	AssertEq(U8(11).SaturatingMul(5), U8(55), "U8.SaturatingMul")
	AssertEq(U8(math.MaxUint8).SaturatingMul(2), U8(math.MaxUint8), "U8.SaturatingMul overflow")

	AssertEq(U16(11).SaturatingMul(12), U16(132), "U16.SaturatingMul")
	AssertEq(U16(math.MaxUint16).SaturatingMul(2), U16(math.MaxUint16), "U16.SaturatingMul overflow")

	AssertEq(U32(11).SaturatingMul(12), U32(132), "U32.SaturatingMul")
	AssertEq(U32(math.MaxUint32).SaturatingMul(2), U32(math.MaxUint32), "U32.SaturatingMul overflow")

	AssertEq(U64(11).SaturatingMul(12), U64(132), "U64.SaturatingMul")
	AssertEq(U64(math.MaxUint64).SaturatingMul(2), U64(math.MaxUint64), "U64.SaturatingMul overflow")
}

func Test_SaturatingDiv_SignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(ISize(100).SaturatingDiv(10), ISize(10), "ISize.SaturatingDiv")
	AssertEq(ISize(math.MaxInt).SaturatingDiv(-1), ISize(math.MinInt+1), "ISize.SaturatingDiv overflow")
	AssertEq(ISize(math.MinInt).SaturatingDiv(-1), ISize(math.MaxInt), "ISize.SaturatingDiv underflow")

	AssertEq(I8(100).SaturatingDiv(10), I8(10), "I8.SaturatingDiv")
	AssertEq(I8(math.MaxInt8).SaturatingDiv(-1), I8(math.MinInt8+1), "I8.SaturatingDiv overflow")
	AssertEq(I8(math.MinInt8).SaturatingDiv(-1), I8(math.MaxInt8), "I8.SaturatingDiv underflow")

	AssertEq(I16(100).SaturatingDiv(10), I16(10), "I16.SaturatingDiv")
	AssertEq(I16(math.MaxInt16).SaturatingDiv(-1), I16(math.MinInt16+1), "I16.SaturatingDiv overflow")
	AssertEq(I16(math.MinInt16).SaturatingDiv(-1), I16(math.MaxInt16), "I16.SaturatingDiv underflow")

	AssertEq(I32(100).SaturatingDiv(10), I32(10), "I32.SaturatingDiv")
	AssertEq(I32(math.MaxInt32).SaturatingDiv(-1), I32(math.MinInt32+1), "I32.SaturatingDiv overflow")
	AssertEq(I32(math.MinInt32).SaturatingDiv(-1), I32(math.MaxInt32), "I32.SaturatingDiv underflow")

	AssertEq(I64(100).SaturatingDiv(10), I64(10), "I64.SaturatingDiv")
	AssertEq(I64(math.MaxInt64).SaturatingDiv(-1), I64(math.MinInt64+1), "I64.SaturatingDiv overflow")
	AssertEq(I64(math.MinInt64).SaturatingDiv(-1), I64(math.MaxInt64), "I64.SaturatingDiv underflow")
}

func Test_SaturatingDiv_UnsignedInteger(t *testing.T) {
	t.Parallel()

	AssertEq(USize(100).SaturatingDiv(10), USize(10), "USize.SaturatingDiv")
	AssertEq(USize(0).SaturatingDiv(10), USize(0), "USize.SaturatingDiv underflow")

	AssertEq(U8(100).SaturatingDiv(10), U8(10), "U8.SaturatingDiv")
	AssertEq(U8(0).SaturatingDiv(10), U8(0), "U8.SaturatingDiv underflow")

	AssertEq(U16(100).SaturatingDiv(10), U16(10), "U16.SaturatingDiv")
	AssertEq(U16(0).SaturatingDiv(10), U16(0), "U16.SaturatingDiv underflow")

	AssertEq(U32(100).SaturatingDiv(10), U32(10), "U32.SaturatingDiv")
	AssertEq(U32(0).SaturatingDiv(10), U32(0), "U32.SaturatingDiv underflow")

	AssertEq(U64(100).SaturatingDiv(10), U64(10), "U64.SaturatingDiv")
	AssertEq(U64(0).SaturatingDiv(10), U64(0), "U64.SaturatingDiv underflow")
}

func Test_U128_Add(t *testing.T) {
	t.Parallel()

	AssertEq(U128_FromU64(10).Add(U128_FromU64(5)), U128_FromU64(15), "U128.Add")
	AssertEq(U128_FromU64(U64_MAX).Add(U128_FromU64(1)), U128{high: 1, low: 0}, "U128.Add over range of U64")
}

func Test_I128_Add(t *testing.T) {
	t.Parallel()

	AssertEq(I128_FromI64(10).Add(I128_FromI64(5)), I128_FromI64(15), "I128.Add")
	AssertEq(I128_FromI64(I64_MAX).Add(I128_FromI64(1)), I128{high: 0, low: U64(I64_MAX) + 1}, "I128.Add over range of I64")
	AssertEq(I128_FromI64(I64_MAX).Add(I128_FromI64(I64_MAX)).Add(I128_FromI64(2)), I128{high: 1, low: 0}, "I128.Add over range of U64")

	AssertEq(I128_FromI64(-10).Add(I128_FromI64(-5)), I128_FromI64(-15), "I128.Add negative")
	AssertEq(I128_FromI64(I64_MIN).Add(I128_FromI64(-1)), I128_FromU64(U64(I64_MIN.Abs())+1).Neg(), "I128.Add negative over range of I64")
}

func Test_U128_Sub(t *testing.T) {
	t.Parallel()

	AssertEq(U128_FromU64(10).Sub(U128_FromU64(5)), U128_FromU64(5), "U128.Sub")
	AssertEq(U128_FromU64(5).Sub(U128_FromU64(10)).ToString(), "340282366920938463463374607431768211451", "U128.Sub underflow")
}

func Test_I128_Sub(t *testing.T) {
	t.Parallel()

	AssertEq(I128_FromI64(10).Sub(I128_FromI64(5)), I128_FromI64(5), "I128.Sub 10-5")
	AssertEq(I128_FromI64(5).Sub(I128_FromI64(10)), I128_FromI64(-5), "I128.Sub 5-10")

	AssertEq(I128_FromI64(-10).Sub(I128_FromI64(-5)), I128_FromI64(-5), "I128.Sub -10-(-5)")
	AssertEq(I128_FromI64(-5).Sub(I128_FromI64(10)), I128_FromI64(-15), "I128.Sub -5-10")

	AssertEq(I128_MAX.Sub(I128_FromI64(-1)).ToString(), "-170141183460469231731687303715884105728", "I128.Sub overflow")
	AssertEq(I128_MIN.Sub(I128_FromI64(1)).ToString(), "170141183460469231731687303715884105727", "I128.Sub underflow")
}

func Test_I128_Mul(t *testing.T) {
	t.Parallel()

	AssertEq(I128_FromI64(10).Mul(I128_FromI64(5)), I128_FromI64(50), "I128.Mul 10*5")
	AssertEq(I128_FromI64(5).Mul(I128_FromI64(10)), I128_FromI64(50), "I128.Mul 5*10")

	AssertEq(I128_FromI64(-10).Mul(I128_FromI64(-5)), I128_FromI64(50), "I128.Mul -10*-5")
	AssertEq(I128_FromI64(-5).Mul(I128_FromI64(10)), I128_FromI64(-50), "I128.Mul -5*10")
}
