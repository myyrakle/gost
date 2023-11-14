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
	_, frac := math.Modf(float64(self))

	return F32(frac)
}

// Returns the fractional part of self.
//
//  f := gost.F64(3.7)
//  g := gost.F64(-3.7)
//
//  gost.AssertEq(f.Frac(), F64(0.7))
//  gost.AssertEq(g.Frac(), F64(-0.7))
func (self F64) Frac() F64 {
	_, frac := math.Modf(float64(self))

	return F64(frac)
}

// Raises a number to an integer power.
// Using this function is generally faster than using powf. It might have a different sequence of rounding operations than powf, so the results are not guaranteed to agree.
//
//  f := gost.F32(3.0)
//  g := gost.F32(3.0)
//
//  gost.AssertEq(f.Powi(2), F32(9.0))
//  gost.AssertEq(g.Powi(3), F32(27.0))
func (self F32) Powi(n int) F32 {
	return F32(math.Pow(float64(self), float64(n)))
}

// Raises a number to an integer power.
// Using this function is generally faster than using powf. It might have a different sequence of rounding operations than powf, so the results are not guaranteed to agree.
//
//  f := gost.F64(3.0)
//  g := gost.F64(3.0)
//
//  gost.AssertEq(f.Powi(2), F64(9.0))
//  gost.AssertEq(g.Powi(3), F64(27.0))
func (self F64) Powi(n int) F64 {
	return F64(math.Pow(float64(self), float64(n)))
}

// Raises a number to a floating point power.
//
//  f := gost.F32(3.0)
//  g := gost.F32(3.0)
//
//  gost.AssertEq(f.Powf(2.0), F32(9.0))
//  gost.AssertEq(g.Powf(3.0), F32(27.0))
func (self F32) Powf(n F32) F32 {
	return F32(math.Pow(float64(self), float64(n)))
}

// Raises a number to a floating point power.
//
//  f := gost.F64(3.0)
//  g := gost.F64(3.0)
//
//  gost.AssertEq(f.Powf(2.0), F64(9.0))
//  gost.AssertEq(g.Powf(3.0), F64(27.0))
func (self F64) Powf(n F64) F64 {
	return F64(math.Pow(float64(self), float64(n)))
}

// Returns the square root of a number.
// Returns NaN if self is a negative number other than -0.0.
//
//  f := gost.F32(9.0)
//  g := gost.F32(16.0)
//
//  gost.AssertEq(f.Sqrt(), F32(3.0))
//  gost.AssertEq(g.Sqrt(), F32(4.0))
func (self F32) Sqrt() F32 {
	return F32(math.Sqrt(float64(self)))
}

// Returns the square root of a number.
// Returns NaN if self is a negative number other than -0.0.
//
//  f := gost.F64(9.0)
//  g := gost.F64(16.0)
//
//  gost.AssertEq(f.Sqrt(), F64(3.0))
//  gost.AssertEq(g.Sqrt(), F64(4.0))
func (self F64) Sqrt() F64 {
	return F64(math.Sqrt(float64(self)))
}

// Returns e^(self), (the exponential function).
//
//  f := gost.F32(1.0)
//  g := gost.F32(2.0)
//
//  gost.AssertEq(f.Exp(), F32(2.7182817))
//  gost.AssertEq(g.Exp(), F32(7.389056))
func (self F32) Exp() F32 {
	return F32(math.Exp(float64(self)))
}

// Returns e^(self), (the exponential function).
//
//  f := gost.F64(1.0)
//  g := gost.F64(2.0)
//
//  gost.AssertEq(f.Exp(), F64(2.718281828459045))
//  gost.AssertEq(g.Exp(), F64(7.38905609893065))
func (self F64) Exp() F64 {
	return F64(math.Exp(float64(self)))
}

// Returns 2^(self).
//
//  f := gost.F32(1.0)
//  g := gost.F32(2.0)
//
//  gost.AssertEq(f.Exp2(), F32(2.0))
//  gost.AssertEq(g.Exp2(), F32(4.0))
func (self F32) Exp2() F32 {
	return F32(math.Exp2(float64(self)))
}

// Returns 2^(self).
//
//  f := gost.F64(1.0)
//  g := gost.F64(2.0)
//
//  gost.AssertEq(f.Exp2(), F64(2.0))
//  gost.AssertEq(g.Exp2(), F64(4.0))
func (self F64) Exp2() F64 {
	return F64(math.Exp2(float64(self)))
}

// Returns the natural logarithm of the number.
//
//  f := gost.F32(2.7182817)
//  g := gost.F32(7.389056)
//
//  gost.AssertEq(f.Ln(), F32(1.0))
//  gost.AssertEq(g.Ln(), F32(2.0))
func (self F32) Ln() F32 {
	return F32(math.Log(float64(self)))
}

// Returns the natural logarithm of the number.
//
//  f := gost.F64(2.718281828459045)
//  g := gost.F64(7.38905609893065)
//
//  gost.AssertEq(f.Ln(), F64(1.0))
//  gost.AssertEq(g.Ln(), F64(2.0))
func (self F64) Ln() F64 {
	return F64(math.Log(float64(self)))
}

// Returns the logarithm of the number with respect to an arbitrary base.
// The result might not be correctly rounded owing to implementation details; self.log2() can produce more accurate results for base 2, and self.log10() can produce more accurate results for base 10.
//
//  f := gost.F32(2.0)
//  g := gost.F32(4.0)
//
//  gost.AssertEq(f.Log(2.0), F32(1.0))
//  gost.AssertEq(g.Log(2.0), F32(2.0))
func (self F32) Log(base F32) F32 {
	return F32(math.Log(float64(self)) / math.Log(float64(base)))
}

// Returns the logarithm of the number with respect to an arbitrary base.
// The result might not be correctly rounded owing to implementation details; self.log2() can produce more accurate results for base 2, and self.log10() can produce more accurate results for base 10.
//
//  f := gost.F64(2.0)
//  g := gost.F64(4.0)
//
//  gost.AssertEq(f.Log(2.0), F64(1.0))
//  gost.AssertEq(g.Log(2.0), F64(2.0))
func (self F64) Log(base F64) F64 {
	return F64(math.Log(float64(self)) / math.Log(float64(base)))
}

// Returns the base 2 logarithm of the number.
//
//  f := gost.F32(2.0)
//  g := gost.F32(4.0)
//
//  gost.AssertEq(f.Log2(), F32(1.0))
//  gost.AssertEq(g.Log2(), F32(2.0))
func (self F32) Log2() F32 {
	return F32(math.Log2(float64(self)))
}

// Returns the base 2 logarithm of the number.
//
//  f := gost.F64(2.0)
//  g := gost.F64(4.0)
//
//  gost.AssertEq(f.Log2(), F64(1.0))
//  gost.AssertEq(g.Log2(), F64(2.0))
func (self F64) Log2() F64 {
	return F64(math.Log2(float64(self)))
}
