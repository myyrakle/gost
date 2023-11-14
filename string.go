package gost

import (
	"strconv"
	"strings"
	"unicode/utf16"
)

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

// Returns an iterator over the chars of a string slice.
func (self String) Chars() Vec[Char] {
	chars := VecNew[Char]()

	for _, c := range self {
		chars.Push(Char(c))
	}

	return chars
}

// Returns true if the given pattern matches a sub-slice of this string slice.
// Returns false if it does not.
// The pattern can be a &str, char, a slice of chars, or a function or closure that determines if a character matches.
func (self String) Contains(str String) bool {
	return strings.Contains(string(self), string(str))
}

// Returns the byte index of the first character of this string slice that matches the pattern.
// Returns None if the pattern doesn’t match.
// The pattern can be a &str, char, a slice of chars, or a function or closure that determines if a character matches.
func (self String) Find(str String) Option[USize] {
	index := strings.Index(string(self), string(str))

	if index == -1 {
		return None[USize]()
	}

	return Some[USize](USize(index))
}

// An iterator over substrings of this string slice, separated by characters matched by a string
func (self String) Split(str String) Vec[String] {
	split := VecNew[String]()

	for _, s := range strings.Split(string(self), string(str)) {
		split.Push(String(s))
	}

	return split
}

// Returns a string slice with leading and trailing whitespace removed.
// ‘Whitespace’ is defined according to the terms of the Unicode Derived Core Property White_Space, which includes newlines.
func (self String) Trim() String {
	return String(strings.TrimSpace(string(self)))
}

// Parses this string slice into another type. (I64)
func (self String) ParseI64() Result[I64] {
	parsed, err := strconv.Atoi(string(self))

	if err != nil {
		return Err[I64](err)
	} else {
		return Ok[I64](I64(parsed))
	}
}

// Parses this string slice into another type. (U64)
func (self String) ParseU64() Result[U64] {
	parsed, err := strconv.Atoi(string(self))

	if err != nil {
		return Err[U64](err)
	} else {
		return Ok[U64](U64(parsed))
	}
}

// Parses this string slice into another type. (F64)
func (self String) ParseF64() Result[F64] {
	parsed, err := strconv.ParseFloat(string(self), 64)

	if err != nil {
		return Err[F64](err)
	} else {
		return Ok[F64](F64(parsed))
	}
}

// Replaces all matches of a pattern with another string.
// replace creates a new String, and copies the data from this string slice into it. While doing so, it attempts to find matches of a pattern. If it finds any, it replaces them with the replacement string slice.
func (self String) Replace(old String, new String) String {
	return String(strings.Replace(string(self), string(old), string(new), -1))
}

// Returns the lowercase equivalent of this string slice, as a new String.
func (self String) ToLowercase() String {
	return String(strings.ToLower(string(self)))
}

// Returns the uppercase equivalent of this string slice, as a new String.
func (self String) ToUppercase() String {
	return String(strings.ToUpper(string(self)))
}

// Creates a new String by repeating a string n times.
func (self String) Repeat(n USize) String {
	return String(strings.Repeat(string(self), int(n)))
}
