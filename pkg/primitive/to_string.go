package primitive

import "strconv"

func (self Int) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Int8) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Int16) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Int32) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Int64) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Uint) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Uint8) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Uint16) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Uint32) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Uint64) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Float32) ToString() string {
	return strconv.FormatFloat(float64(self), 'f', -1, 32)
}

func (self Float64) ToString() string {
	return strconv.FormatFloat(float64(self), 'f', -1, 64)
}

func (self Byte) ToString() string {
	return strconv.Itoa(int(self))
}

func (self Rune) ToString() string {
	return strconv.Itoa(int(self))
}

func (self String) ToString() string {
	return string(self)
}

func (self Bool) ToString() string {
	return strconv.FormatBool(bool(self))
}

func (self Complex64) ToString() string {
	return strconv.FormatComplex(complex128(self), 'f', -1, 64)
}

func (self Complex128) ToString() string {
	return strconv.FormatComplex(complex128(self), 'f', -1, 128)
}
