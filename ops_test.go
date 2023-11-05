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
