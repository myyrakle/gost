package gost

import "testing"

func Test_I128_Shl(t *testing.T) {
	t.Parallel()

	AssertEq(I128_FromU64(U64_MAX).Shl(I128_FromI64(1)).ToString(), "36893488147419103230", "I128.Shl out of U64 range")
	AssertEq(I128_FromU64(30).Shl(I128_FromI64(1)).ToString(), "60", "I128.Shl")
}

func Test_U128_Shl(t *testing.T) {
	t.Parallel()

	AssertEq(U128_FromU64(U64_MAX).Shl(U128_FromU64(1)).ToString(), "36893488147419103230", "U128.Shl out of U64 range")
	AssertEq(U128_FromU64(30).Shl(U128_FromU64(1)).ToString(), "60", "U128.Shl")
}

func Test_I128_Shr(t *testing.T) {
	t.Parallel()

	AssertEq(I128_FromU64(U64_MAX).Shl(I128_FromI64(1)).Shr(I128_FromI64(1)), I128_FromU64(U64_MAX), "I128.Shr out of U64 range")
	AssertEq(I128_FromU64(60).Shr(I128_FromI64(1)).ToString(), "30", "I128.Shr")
}
