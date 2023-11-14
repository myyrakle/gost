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

// Returns the nearest integer to self. If a value is half-way between two integers, round away from 0.0.
//
//  f := gost.F32(3.3)
//  g := gost.F32(-3.3)
//  h := gost.F32(-3.7)
//  i := gost.F32(3.5)
//  j := gost.F32(4.5)
//
//  gost.AssertEq(f.Round(), F32(3.0))
//  gost.AssertEq(g.Round(), F32(-3.0))
//  gost.AssertEq(h.Round(), F32(-4.0))
//  gost.AssertEq(i.Round(), F32(4.0))
//  gost.AssertEq(j.Round(), F32(5.0))
func (self F32) Round() F32 {
	return F32(math.Round(float64(self)))
}

// Returns the nearest integer to self. If a value is half-way between two integers, round away from 0.0.
//
//  f := gost.F64(3.3)
//  g := gost.F64(-3.3)
//  h := gost.F64(-3.7)
//  i := gost.F64(3.5)
//  j := gost.F64(4.5)
//
//  gost.AssertEq(f.Round(), F64(3.0))
//  gost.AssertEq(g.Round(), F64(-3.0))
//  gost.AssertEq(h.Round(), F64(-4.0))
//  gost.AssertEq(i.Round(), F64(4.0))
//  gost.AssertEq(j.Round(), F64(5.0))
func (self F64) Round() F64 {
	return F64(math.Round(float64(self)))
}

// Computes the absolute value of self.
//
//  f := gost.F32(-3.0)
//  g := gost.F32(3.0)
//
//  gost.AssertEq(f.Abs(), F32(3.0))
//  gost.AssertEq(g.Abs(), F32(3.0))
func (self F32) Abs() F32 {
	return F32(math.Abs(float64(self)))
}

// Computes the absolute value of self.
//
//  f := gost.F64(-3.0)
//  g := gost.F64(3.0)
//
//  gost.AssertEq(f.Abs(), F64(3.0))
//  gost.AssertEq(g.Abs(), F64(3.0))
func (self F64) Abs() F64 {
	return F64(math.Abs(float64(self)))
}

// Returns the integer part of self. This means that non-integer numbers are always truncated towards zero.
//
//  f := gost.F32(3.7)
//  g := gost.F32(-3.7)
//
//  gost.AssertEq(f.Trunc(), F32(3.0))
//  gost.AssertEq(g.Trunc(), F32(-3.0))
func (self F32) Trunc() F32 {
	return F32(math.Trunc(float64(self)))
}

// Returns the integer part of self. This means that non-integer numbers are always truncated towards zero.
//
//  f := gost.F64(3.7)
//  g := gost.F64(-3.7)
//
//  gost.AssertEq(f.Trunc(), F64(3.0))
//  gost.AssertEq(g.Trunc(), F64(-3.0))
func (self F64) Trunc() F64 {
	return F64(math.Trunc(float64(self)))
}

// Returns the fractional part of self.
//
//  f := gost.F32(3.7)
//  g := gost.F32(-3.7)
//
//  gost.AssertEq(f.Frac(), F32(0.7))
//  gost.AssertEq(g.Frac(), F32(-0.7))
func (self F32) Frac() F32 {
	return F32(math.Mod(float64(self), 1.0))
}

// Returns the fractional part of self.
//
//  f := gost.F64(3.7)
//  g := gost.F64(-3.7)
//
//  gost.AssertEq(f.Frac(), F64(0.7))
//  gost.AssertEq(g.Frac(), F64(-0.7))
func (self F64) Frac() F64 {
	return F64(math.Mod(float64(self), 1.0))
}
