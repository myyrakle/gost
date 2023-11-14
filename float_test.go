package gost

import "testing"

func Test_F32_Floor(t *testing.T) {
	t.Parallel()

	f := F32(3.7)
	g := F32(3.0)
	h := F32(-3.7)

	AssertEq(f.Floor(), F32(3.0))
	AssertEq(g.Floor(), F32(3.0))
	AssertEq(h.Floor(), F32(-4.0))
}

func Test_F64_Floor(t *testing.T) {
	t.Parallel()

	f := F64(3.7)
	g := F64(3.0)
	h := F64(-3.7)

	AssertEq(f.Floor(), F64(3.0))
	AssertEq(g.Floor(), F64(3.0))
	AssertEq(h.Floor(), F64(-4.0))
}

func Test_F32_Ceil(t *testing.T) {
	t.Parallel()

	f := F32(3.01)
	g := F32(4.0)

	AssertEq(f.Ceil(), F32(4.0))
	AssertEq(g.Ceil(), F32(4.0))
}

func Test_F64_Ceil(t *testing.T) {
	t.Parallel()

	f := F64(3.01)
	g := F64(4.0)

	AssertEq(f.Ceil(), F64(4.0))
	AssertEq(g.Ceil(), F64(4.0))
}

func Test_F32_Round(t *testing.T) {
	t.Parallel()

	f := F32(3.3)
	g := F32(-3.3)
	h := F32(-3.7)
	i := F32(3.5)
	j := F32(4.5)

	AssertEq(f.Round(), F32(3.0))
	AssertEq(g.Round(), F32(-3.0))
	AssertEq(h.Round(), F32(-4.0))
	AssertEq(i.Round(), F32(4.0))
	AssertEq(j.Round(), F32(5.0))
}

func Test_F64_Round(t *testing.T) {
	t.Parallel()

	f := F64(3.3)
	g := F64(-3.3)
	h := F64(-3.7)
	i := F64(3.5)
	j := F64(4.5)

	AssertEq(f.Round(), F64(3.0))
	AssertEq(g.Round(), F64(-3.0))
	AssertEq(h.Round(), F64(-4.0))
	AssertEq(i.Round(), F64(4.0))
	AssertEq(j.Round(), F64(5.0))
}

func Test_F32_Abs(t *testing.T) {
	t.Parallel()

	f := F32(-3.0)
	g := F32(3.0)

	AssertEq(f.Abs(), F32(3.0))
	AssertEq(g.Abs(), F32(3.0))
}

func Test_F64_Abs(t *testing.T) {
	t.Parallel()

	f := F64(-3.0)
	g := F64(3.0)

	AssertEq(f.Abs(), F64(3.0))
	AssertEq(g.Abs(), F64(3.0))
}
