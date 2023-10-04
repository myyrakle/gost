package gost

type Clone[T any] interface {
	Clone() T
}

func (self ISize) Clone() ISize {
	return self
}

func (self ISize8) Clone() ISize8 {
	return self
}

func (self ISize16) Clone() ISize16 {
	return self
}

func (self ISize32) Clone() ISize32 {
	return self
}

func (self ISize64) Clone() ISize64 {
	return self
}

func (self USize) Clone() USize {
	return self
}

func (self USize8) Clone() USize8 {
	return self
}

func (self USize16) Clone() USize16 {
	return self
}

func (self USize32) Clone() USize32 {
	return self
}

func (self USize64) Clone() USize64 {
	return self
}

func (self Float32) Clone() Float32 {
	return self
}

func (self Float64) Clone() Float64 {
	return self
}

func (self Byte) Clone() Byte {
	return self
}

func (self Rune) Clone() Rune {
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
