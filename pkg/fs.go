package gost

import "os"

// Creates a new, empty directory at the provided path
func CreateDir(path String) Result[any] {
	err := os.Mkdir(string(path), os.ModePerm)

	if err != nil {
		return Err[any](err)
	} else {
		return Ok[any](nil)
	}
}

// Removes an empty directory.
func RemoveDir(path String) Result[any] {
	err := os.Remove(string(path))

	if err != nil {
		return Err[any](err)
	} else {
		return Ok[any](nil)
	}
}

// Write a slice as the entire contents of a file.
func Write(path String, data []Byte) Result[any] {
	file, err := os.Create(string(path))

	if err != nil {
		return Err[any](err)
	}

	defer file.Close()

	casted := []byte{}

	for _, b := range data {
		casted = append(casted, byte(b))
	}

	_, err = file.Write([]byte(casted))

	if err != nil {
		return Err[any](err)
	} else {
		return Ok[any](nil)
	}
}
