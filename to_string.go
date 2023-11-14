package gost

import (
	"reflect"
	"strconv"
)

// A trait for converting a value to a String.
type ToString[T any] interface {
	ToString() String
}

func castToToString[T any](value T) Option[ToString[T]] {
	reflectedValue := reflect.ValueOf(value)

	if casted, ok := reflectedValue.Interface().(ToString[T]); ok {
		return Some[ToString[T]](casted)
	} else {
		return None[ToString[T]]()
	}
}

func (self ISize) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self I8) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self I16) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self I32) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self I64) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self USize) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self U8) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self U16) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self U32) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self U64) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self F32) ToString() String {
	return String(strconv.FormatFloat(float64(self), 'f', -1, 32))
}

func (self F64) ToString() String {
	return String(strconv.FormatFloat(float64(self), 'f', -1, 64))
}

func (self Byte) ToString() String {
	return String(strconv.Itoa(int(self)))
}

func (self Char) ToString() String {
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
