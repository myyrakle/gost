package gost

import (
	"math/big"
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

func (self I128) ToString() String {
	isNegative := self.high < 0

	if isNegative {
		self = self.Neg()
	}

	high := self.high
	low := self.low

	if high == 0 {
		if isNegative {
			return "-" + low.ToString()
		} else {
			return low.ToString()
		}
	} else {
		binaryString := ""

		for i := 0; i < 64; i++ {
			binaryString = string((low & 1).ToString()[0]) + binaryString
			low = low >> 1
		}

		for i := 0; i < 64; i++ {
			binaryString = string((high & 1).ToString()[0]) + binaryString
			high = high >> 1
		}

		return _ConvertBinaryStringToDecimalString(String(binaryString), Bool(isNegative))
	}
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

func (self U128) ToString() String {
	high := self.high
	low := self.low

	if high == 0 {
		return low.ToString()
	} else {
		binaryString := ""

		for i := 0; i < 64; i++ {
			binaryString = string((low & 1).ToString()[0]) + binaryString
			low = low >> 1
		}

		for i := 0; i < 64; i++ {
			binaryString = string((high & 1).ToString()[0]) + binaryString
			high = high >> 1
		}

		return _ConvertBinaryStringToDecimalString(String(binaryString), false)
	}
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

func _ConvertBinaryStringToDecimalString(binaryString String, isNegative Bool) String {
	decimal := big.NewInt(0)

	i := len(binaryString) - 1
	digit := big.NewInt(1)
	for i >= 0 {
		if binaryString[i] == '1' {
			decimal.Add(decimal, digit)
		}

		digit = digit.Mul(digit, big.NewInt(2))
		i--
	}

	if isNegative {
		decimal.Neg(decimal)
	}

	return String(decimal.String())
}
