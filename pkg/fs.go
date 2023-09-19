package gost

import (
	"os"
	"path/filepath"
)

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

// Removes a file from the filesystem.
func RemoveFile(path String) Result[any] {
	err := os.Remove(string(path))

	if err != nil {
		return Err[any](err)
	} else {
		return Ok[any](nil)
	}
}

// Rename a file or directory to a new name, replacing the original file if to already exists.
// This will not work if the new name is on a different mount point.
func Rename(from String, to String) Result[any] {
	err := os.Rename(string(from), string(to))

	if err != nil {
		return Err[any](err)
	} else {
		return Ok[any](nil)
	}
}

// Read the entire contents of a file into a bytes vector.
func Read(path String) Result[[]Byte] {
	file, err := os.Open(string(path))

	if err != nil {
		return Err[[]Byte](err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return Err[[]Byte](err)
	}

	data := make([]byte, stat.Size())

	_, err = file.Read(data)

	if err != nil {
		return Err[[]Byte](err)
	} else {
		casted := []Byte{}

		for _, b := range data {
			casted = append(casted, Byte(b))
		}

		return Ok[[]Byte](casted)
	}
}

type FileType struct {
	typeInfo string
}

func (self FileType) IsDir() bool {
	return self.typeInfo == "dir"
}

func (self FileType) IsFile() bool {
	return self.typeInfo == "file"
}

func (self FileType) IsSymlink() bool {
	return self.typeInfo == "symlink"
}

type DirEntry struct {
	FileName String
	Path     String
	FileType FileType
}

// Returns an iterator over the entries within a directory.
func ReadDir(path String) Result[Vec[DirEntry]] {
	entries, err := os.ReadDir(string(path))

	if err != nil {
		return Err[Vec[DirEntry]](err)
	} else {
		vec := VecWithLen[DirEntry](Int(len(entries)))

		for _, entry := range entries {
			fileType := FileType{}

			if entry.Type().Perm().IsDir() {
				fileType.typeInfo = "dir"
			} else if entry.Type().Perm().IsRegular() {
				fileType.typeInfo = "file"
			} else if entry.Type().Perm()&os.ModeSymlink != 0 {
				fileType.typeInfo = "symlink"
			}

			entryPath := filepath.Join(string(path), entry.Name())

			vec.Push(DirEntry{
				FileName: String(entry.Name()),
				FileType: fileType,
				Path:     String(entryPath),
			})
		}

		return Ok[Vec[DirEntry]](vec)
	}
}

// Copies the contents of one file to another. This function will also copy the permission bits of the original file to the destination file.
// This function will overwrite the contents of to.
func Copy(from String, to String) Result[any] {
	err := os.Link(string(from), string(to))

	if err != nil {
		return Err[any](err)
	} else {
		return Ok[any](nil)
	}
}
