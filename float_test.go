package gost

import "testing"

func Test_F64_Floor(t *testing.T) {
	t.Parallel()

	f := F64(3.7)
	g := F64(3.0)
	h := F64(-3.7)

	AssertEq(f.Floor(), F64(3.0))
	AssertEq(g.Floor(), F64(3.0))
	AssertEq(h.Floor(), F64(-4.0))
}
