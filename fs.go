package gofs

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
)

// FlexibleFS represents a flexible file system interface that provides
// various methods for embed and local file operations.
type FlexibleFS interface {
	// Exists checks if a file with the given name exists in the file system.
	// It returns a boolean indicating existence and an error if any occurs.
	Exists(path string) (bool, error)

	// Open opens a file with the given name and returns an fs.File interface
	// for reading the file. It returns an error if the file cannot be opened.
	Open(path string) (fs.File, error)

	// ReadFile reads the entire content of the file with the given name and
	// returns the content as a byte slice. It returns an error if the file
	// cannot be read.
	ReadFile(path string) ([]byte, error)

	// Search searches for a phrase in files within the specified directory,
	// optionally ignoring certain files and filtering by extension.
	Search(dir, phrase, ignore, ext string) (*string, error)

	// Find searches for a file matching the given pattern in the specified
	// path. It returns the path of the found file and an error if any occurs.
	Find(dir, pattern string) (*string, error)

	// Lookup searches for files matching the given pattern in the specified
	// path. It returns a slice of paths of the found files and an error if any occurs.
	Lookup(dir, pattern string) ([]string, error)

	// FS returns the underlying fs.FS interface of the file system.
	FS() fs.FS

	// Http returns the http.FileSystem instance of the file system.
	Http() http.FileSystem
}

// NewDir creates a new FlexibleFS instance using the local file system
// at the specified path.
func NewDir(path string) FlexibleFS {
	driver := new(flexible)
	driver.files = os.DirFS(path)
	return driver
}

// NewEmbed creates a new FlexibleFS instance using the provided embedded
// file system.
func NewEmbed(fs embed.FS) FlexibleFS {
	driver := new(flexible)
	driver.files = fs
	return driver
}
