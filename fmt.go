package gost

import (
	"fmt"
	"reflect"
)

type Display[T any] interface {
	Display() String
}

func (self ISize) Display() String {
	return self.ToString()
}

func (self ISize8) Display() String {
	return self.ToString()
}

func (self ISize16) Display() String {
	return self.ToString()
}

func (self ISize32) Display() String {
	return self.ToString()
}

func (self ISize64) Display() String {
	return self.ToString()
}

func (self USize) Display() String {
	return self.ToString()
}

func (self USize8) Display() String {
	return self.ToString()
}

func (self USize16) Display() String {
	return self.ToString()
}

func (self USize32) Display() String {
	return self.ToString()
}

func (self USize64) Display() String {
	return self.ToString()
}

func (self Float32) Display() String {
	return self.ToString()
}

func (self Float64) Display() String {
	return self.ToString()
}

func (self Byte) Display() String {
	return self.ToString()
}

func (self Rune) Display() String {
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

type Debug[T any] interface {
	Debug() string
}

func (self ISize) Debug() String {
	return String(fmt.Sprintf("ISize(%s)", self.ToString()))
}

func (self ISize8) Debug() String {
	return String(fmt.Sprintf("ISize8(%s)", self.ToString()))
}

func (self ISize16) Debug() String {
	return String(fmt.Sprintf("ISize16(%s)", self.ToString()))
}

func (self ISize32) Debug() String {
	return String(fmt.Sprintf("ISize32(%s)", self.ToString()))
}

func (self ISize64) Debug() String {
	return String(fmt.Sprintf("ISize64(%s)", self.ToString()))
}

func (self USize) Debug() String {
	return String(fmt.Sprintf("USize(%s)", self.ToString()))
}

func (self USize8) Debug() String {
	return String(fmt.Sprintf("USize8(%s)", self.ToString()))
}

func (self USize16) Debug() String {
	return String(fmt.Sprintf("USize16(%s)", self.ToString()))
}

func (self USize32) Debug() String {
	return String(fmt.Sprintf("USize32(%s)", self.ToString()))
}

func (self USize64) Debug() String {
	return String(fmt.Sprintf("USize64(%s)", self.ToString()))
}

func (self Float32) Debug() String {
	return String(fmt.Sprintf("Float32(%s)", self.ToString()))
}

func (self Float64) Debug() String {
	return String(fmt.Sprintf("Float64(%s)", self.ToString()))
}

func (self Byte) Debug() String {
	return String(fmt.Sprintf("Byte(%s)", self.ToString()))
}

func (self Rune) Debug() String {
	return String(fmt.Sprintf("Rune(%s)", self.ToString()))
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
