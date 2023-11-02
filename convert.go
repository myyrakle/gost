package gost

// Used to do a cheap reference-to-reference conversion.
type AsRef[T any] interface {
	AsRef() *T
}

func (self ISize) AsRef() *ISize {
	return &self
}

func (self I8) AsRef() *I8 {
	return &self
}

func (self I16) AsRef() *I16 {
	return &self
}

func (self I32) AsRef() *I32 {
	return &self
}

func (self I64) AsRef() *I64 {
	return &self
}

func (self USize) AsRef() *USize {
	return &self
}

func (self U8) AsRef() *U8 {
	return &self
}

func (self U16) AsRef() *U16 {
	return &self
}

func (self U32) AsRef() *U32 {
	return &self
}

func (self U64) AsRef() *U64 {
	return &self
}

func (self F32) AsRef() *F32 {
	return &self
}

func (self F64) AsRef() *F64 {
	return &self
}

func (self Byte) AsRef() *Byte {
	return &self
}

func (self Char) AsRef() *Char {
	return &self
}

func (self String) AsRef() *String {
	return &self
}

func (self Bool) AsRef() *Bool {
	return &self
}

func (self Complex64) AsRef() *Complex64 {
	return &self
}

func (self Complex128) AsRef() *Complex128 {
	return &self
}
