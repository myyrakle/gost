package gokio

import (
	"github.com/myyrakle/gost"
)

// Creates a new, empty directory at the provided path
func CreateDir(path gost.String) gost.Future[gost.Result[any]] {
	return Spawn(func() gost.Result[any] {
		return gost.CreateDir(path)
	})
}

// Removes an empty directory.
func RemoveDir(path gost.String) gost.Future[gost.Result[any]] {
	return Spawn(func() gost.Result[any] {
		return gost.RemoveDir(path)
	})
}

// Write a slice as the entire contents of a file.
func Write(path gost.String, data []gost.Byte) gost.Future[gost.Result[any]] {
	return Spawn(func() gost.Result[any] {
		return gost.Write(path, data)
	})
}

// Removes a file from the filesystem.
func RemoveFile(path gost.String) gost.Future[gost.Result[any]] {
	return Spawn(func() gost.Result[any] {
		return gost.RemoveFile(path)
	})
}

// Rename a file or directory to a new name, replacing the original file if to already exists.
// This will not work if the new name is on a different mount point.
func Rename(from gost.String, to gost.String) gost.Future[gost.Result[any]] {
	return Spawn(func() gost.Result[any] {
		return gost.Rename(from, to)
	})
}

// Read the entire contents of a file into a bytes vector.
func Read(path gost.String) gost.Future[gost.Result[[]gost.Byte]] {
	return Spawn(func() gost.Result[[]gost.Byte] {
		return gost.Read(path)
	})
}

// Returns an iterator over the entries within a directory.
func ReadDir(path gost.String) gost.Future[gost.Result[gost.Vec[gost.DirEntry]]] {
	return Spawn(func() gost.Result[gost.Vec[gost.DirEntry]] {
		return gost.ReadDir(path)
	})
}

// Copies the contents of one file to another. This function will also copy the permission bits of the original file to the destination file.
// This function will overwrite the contents of to.
func Copy(from gost.String, to gost.String) gost.Future[gost.Result[any]] {
	return Spawn(func() gost.Result[any] {
		return gost.Copy(from, to)
	})
}

// Read the entire contents of a file into a string.
func ReadToString(path gost.String) gost.Future[gost.Result[gost.String]] {
	return Spawn(func() gost.Result[gost.String] {
		return gost.ReadToString(path)
	})
}
