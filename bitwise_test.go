package gost

import "testing"

func Test_I128_Shl(t *testing.T) {
	t.Parallel()

	AssertEq(I128_FromU64(U64_MAX).Shl(I128_FromI64(1)).ToString(), "36893488147419103230", "I128.Shl out of U64 range")
	AssertEq(I128_FromU64(30).Shl(I128_FromI64(1)).ToString(), "60", "I128.Shl")
}
