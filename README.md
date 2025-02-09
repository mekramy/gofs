# FlexibleFS

FlexibleFS is a Go package that provides a flexible file system interface for various file operations, including embedded and local file systems.

## Installation

To install the package, use the following command:

```sh
go get github.com/mekramy/gofs
```

## Usage

### Creating a FlexibleFS Instance

You can create a `FlexibleFS` instance using either the local file system or an embedded file system.

**NOTE:** FlexibleFS use `/` as path separator.

#### Using Local File System

```go
package main

import (
    "fmt"
    "log"
    "github.com/mekramy/gofs"
)

func main() {
    fs := gofs.NewDir("/path/to/directory")

    exists, err := fs.Exists("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File exists:", exists)
}
```

#### Using Embedded File System

```go
package main

import (
    "embed"
    "fmt"
    "log"
    "github.com/mekramy/gofs"
)

//go:embed files/*
var embeddedFiles embed.FS

func main() {
    fs := gofs.NewEmbed(embeddedFiles)

    content, err := fs.ReadFile("files/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File content:", string(content))
}
```

## API

### `Exists(path string) (bool, error)`

Checks if a file with the given name exists in the file system.

### `Open(path string) (fs.File, error)`

Opens a file with the given name and returns an `fs.File` interface for reading the file.

### `ReadFile(path string) ([]byte, error)`

Reads the entire content of the file with the given name and returns the content as a byte slice.

### `Search(dir, phrase, ignore, ext string) (*string, error)`

Searches for a phrase in files within the specified directory, optionally ignoring certain files and filtering by extension.

### `Find(dir, pattern string) (*string, error)`

Searches for a file matching the given pattern in the specified path.

### `Lookup(dir, pattern string) ([]string, error)`

Searches for files matching the given pattern in the specified path.

### `FS() fs.FS`

Returns the underlying `fs.FS` interface of the file system.

### `Http() http.FileSystem`

Returns the `http.FileSystem` instance of the file system.

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
