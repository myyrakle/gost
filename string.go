package gost

import (
	"strconv"
)

// A trait for converting a value to a String.
type ToString[T any] interface {
	ToString() String
}

func (self ISize) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self ISize8) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self ISize16) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self ISize32) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self ISize64) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self USize) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self USize8) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self USize16) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self USize32) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self USize64) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self Float32) ToString() String {
	return String(strconv.FormatFloat(float64(self), 'f', -1, 32))
}

func (self Float64) ToString() String {
	return String(strconv.FormatFloat(float64(self), 'f', -1, 64))
}

func (self Byte) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self Rune) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self String) ToString() String {
	return String(self)
}

func (self Bool) ToString() String {
	return String(strconv.FormatBool(bool(self)))
}

func (self Complex64) ToString() String {
	return String(strconv.FormatComplex(complex128(self), 'f', -1, 64))
}

func (self Complex128) ToString() String {
	return String(strconv.FormatComplex(complex128(self), 'f', -1, 128))
}
