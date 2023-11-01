package gost

import (
	"fmt"
	"reflect"
)

// Format trait for an empty format, {}
type Display[T any] interface {
	Display() String
}

func (self ISize) Display() String {
	return self.ToString()
}

func (self I8) Display() String {
	return self.ToString()
}

func (self I16) Display() String {
	return self.ToString()
}

func (self I32) Display() String {
	return self.ToString()
}

func (self I64) Display() String {
	return self.ToString()
}

func (self USize) Display() String {
	return self.ToString()
}

func (self U8) Display() String {
	return self.ToString()
}

func (self U16) Display() String {
	return self.ToString()
}

func (self U32) Display() String {
	return self.ToString()
}

func (self U64) Display() String {
	return self.ToString()
}

func (self F32) Display() String {
	return self.ToString()
}

func (self F64) Display() String {
	return self.ToString()
}

func (self Byte) Display() String {
	return self.ToString()
}

func (self Char) Display() String {
	return self.ToString()
}

func (self String) Display() String {
	return self
}

func (self Bool) Display() String {
	return self.ToString()
}

func (self Complex64) Display() String {
	return self.ToString()
}

func (self Complex128) Display() String {
	return self.ToString()
}

func castToDisplay[T any](value T) Option[Display[T]] {
	reflectedValue := reflect.ValueOf(value)

	if display, ok := reflectedValue.Interface().(Display[T]); ok {
		return Some[Display[T]](display)
	} else {
		return None[Display[T]]()
	}
}

// ? formatting.
// Debug should format the output in a programmer-facing, debugging context.
type Debug[T any] interface {
	Debug() string
}

func (self ISize) Debug() String {
	return String(fmt.Sprintf("ISize(%s)", self.ToString()))
}

func (self I8) Debug() String {
	return String(fmt.Sprintf("I8(%s)", self.ToString()))
}

func (self I16) Debug() String {
	return String(fmt.Sprintf("I16(%s)", self.ToString()))
}

func (self I32) Debug() String {
	return String(fmt.Sprintf("I32(%s)", self.ToString()))
}

func (self I64) Debug() String {
	return String(fmt.Sprintf("I64(%s)", self.ToString()))
}

func (self USize) Debug() String {
	return String(fmt.Sprintf("USize(%s)", self.ToString()))
}

func (self U8) Debug() String {
	return String(fmt.Sprintf("U8(%s)", self.ToString()))
}

func (self U16) Debug() String {
	return String(fmt.Sprintf("U16(%s)", self.ToString()))
}

func (self U32) Debug() String {
	return String(fmt.Sprintf("U32(%s)", self.ToString()))
}

func (self U64) Debug() String {
	return String(fmt.Sprintf("U64(%s)", self.ToString()))
}

func (self F32) Debug() String {
	return String(fmt.Sprintf("F32(%s)", self.ToString()))
}

func (self F64) Debug() String {
	return String(fmt.Sprintf("F64(%s)", self.ToString()))
}

func (self Byte) Debug() String {
	return String(fmt.Sprintf("Byte(%s)", self.ToString()))
}

func (self Char) Debug() String {
	return String(fmt.Sprintf("Char(%s)", self.ToString()))
}

func (self String) Debug() String {
	return String(fmt.Sprintf("String(%s)", self))
}

func (self Bool) Debug() String {
	return String(fmt.Sprintf("Bool(%s)", self.ToString()))
}

func (self Complex64) Debug() String {
	return String(fmt.Sprintf("Complex64(%s)", self.ToString()))
}

func (self Complex128) Debug() String {
	return String(fmt.Sprintf("Complex128(%s)", self.ToString()))
}
