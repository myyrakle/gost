package gost

import "unicode/utf16"

// Converts a vector of bytes to a String.
func StringFromUTF8(bytes Vec[U8]) String {
	return String(bytes.data)
}

// Decode a UTF-16–encoded vector v into a String, returning Err if v contains any invalid data.
func StringFromUTF16(bytes Vec[U16]) Result[String] {
	data := bytes.data
	convertedData := []uint16{}

	for i := 0; i < len(data); i += 2 {
		convertedData = append(convertedData, uint16(data[i]))
	}

	decoded := utf16.Decode(convertedData)

	return Ok[String](String(decoded))
}

// Converts a String into a byte vector.
// This consumes the String, so we do not need to copy its contents.
func (self String) IntoBytes() Vec[U8] {
	bytes := VecNew[U8]()

	for _, b := range []byte(self) {
		bytes.Push(U8(b))
	}

	return bytes
}

// Returns the length of this String, in bytes, not chars or graphemes.
// In other words, it might not be what a human considers the length of the string.
func (self String) Len() USize {
	return USize(len(self))
}

// Returns true if this String has a length of zero, and false otherwise.
func (self String) IsEmpty() bool {
	return len(self) == 0
}

// Appends the given char to the end of this String.
func (self *String) Push(c Char) String {
	*self += String(c)
	return *self
}

// Removes the last character from the string buffer and returns it.
// Returns None if this String is empty.
func (self *String) Pop() Option[Char] {
	if self.IsEmpty() {
		return None[Char]()
	}

	runes := []rune(*self)

	last := runes[len(runes)-1]
	*self = String(runes[:len(runes)-1])

	return Some[Char](Char(last))
}
