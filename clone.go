package gost

import "reflect"

type Clone[T any] interface {
	Clone() T
}

func castToClone[T any](value T) Option[Clone[T]] {
	reflectedValue := reflect.ValueOf(value)

	if casted, ok := reflectedValue.Interface().(Clone[T]); ok {
		return Some[Clone[T]](casted)
	} else {
		return None[Clone[T]]()
	}
}

func (self ISize) Clone() ISize {
	return self
}

func (self I8) Clone() I8 {
	return self
}

func (self I16) Clone() I16 {
	return self
}

func (self I32) Clone() I32 {
	return self
}

func (self I64) Clone() I64 {
	return self
}

func (self USize) Clone() USize {
	return self
}

func (self U8) Clone() U8 {
	return self
}

func (self U16) Clone() U16 {
	return self
}

func (self U32) Clone() U32 {
	return self
}

func (self U64) Clone() U64 {
	return self
}

func (self F32) Clone() F32 {
	return self
}

func (self F64) Clone() F64 {
	return self
}

func (self Byte) Clone() Byte {
	return self
}

func (self Char) Clone() Char {
	return self
}

func (self String) Clone() String {
	return self
}

func (self Bool) Clone() Bool {
	return self
}

func (self Complex64) Clone() Complex64 {
	return self
}

func (self Complex128) Clone() Complex128 {
	return self
}
