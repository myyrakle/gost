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

// Returns the smallest integer greater than or equal to self.
//
//  f := gost.F32(3.01)
//  g := gost.F32(4.0)
//
//  gost.AssertEq(f.Ceil(), F32(4.0))
//  gost.AssertEq(g.Ceil(), F32(4.0))
func (self F32) Ceil() F32 {
	return F32(math.Ceil(float64(self)))
}

// Returns the smallest integer greater than or equal to self.
//
//  f := gost.F64(3.01)
//  g := gost.F64(4.0)
//
//  gost.AssertEq(f.Ceil(), F64(4.0))
//  gost.AssertEq(g.Ceil(), F64(4.0))
func (self F64) Ceil() F64 {
	return F64(math.Ceil(float64(self)))
}
