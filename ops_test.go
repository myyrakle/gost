package gost

import (
	"math"
	"testing"
)

func Test_WrappingAdd_SignedInteger(t *testing.T) {
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
