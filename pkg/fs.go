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
