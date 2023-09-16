package gost

import (
	"fmt"
	"reflect"
)

type Display[T any] interface {
	Display() String
}

func (self Int) Display() String {
	return self.ToString()
}

func (self Int8) Display() String {
	return self.ToString()
}

func (self Int16) Display() String {
	return self.ToString()
}

func (self Int32) Display() String {
	return self.ToString()
}

func (self Int64) Display() String {
	return self.ToString()
}

func (self Uint) Display() String {
	return self.ToString()
}

func (self Uint8) Display() String {
	return self.ToString()
}

func (self Uint16) Display() String {
	return self.ToString()
}

func (self Uint32) Display() String {
	return self.ToString()
}

func (self Uint64) Display() String {
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

func (self Int) Debug() String {
	return String(fmt.Sprintf("Int(%s)", self.ToString()))
}

func (self Int8) Debug() String {
	return String(fmt.Sprintf("Int8(%s)", self.ToString()))
}

func (self Int16) Debug() String {
	return String(fmt.Sprintf("Int16(%s)", self.ToString()))
}

func (self Int32) Debug() String {
	return String(fmt.Sprintf("Int32(%s)", self.ToString()))
}

func (self Int64) Debug() String {
	return String(fmt.Sprintf("Int64(%s)", self.ToString()))
}

func (self Uint) Debug() String {
	return String(fmt.Sprintf("Uint(%s)", self.ToString()))
}

func (self Uint8) Debug() String {
	return String(fmt.Sprintf("Uint8(%s)", self.ToString()))
}

func (self Uint16) Debug() String {
	return String(fmt.Sprintf("Uint16(%s)", self.ToString()))
}

func (self Uint32) Debug() String {
	return String(fmt.Sprintf("Uint32(%s)", self.ToString()))
}

func (self Uint64) Debug() String {
	return String(fmt.Sprintf("Uint64(%s)", self.ToString()))
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
