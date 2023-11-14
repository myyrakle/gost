package gost

import "math"

// Returns the largest integer less than or equal to self.
//
//  f := gost.F32(3.7)
//  g := gost.F32(3.0)
//  h := gost.F32(-3.7)
//
//  gost.AssertEq(f.Floor(), F32(3.0))
//  gost.AssertEq(g.Floor(), F32(3.0))
//  gost.AssertEq(h.Floor(), F32(-4.0))
func (self F32) Floor() F32 {
	return F32(math.Floor(float64(self)))
}

// Returns the largest integer less than or equal to self.
//
//  f := gost.F64(3.7)
//  g := gost.F64(3.0)
//  h := gost.F64(-3.7)
//
//  gost.AssertEq(f.Floor(), F64(3.0))
//  gost.AssertEq(g.Floor(), F64(3.0))
//  gost.AssertEq(h.Floor(), F64(-4.0))
func (self F64) Floor() F64 {
	return F64(math.Floor(float64(self)))
}
