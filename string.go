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
