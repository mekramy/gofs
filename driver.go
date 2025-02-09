package gofs

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type flexible struct {
	files fs.FS
}

func (fxs flexible) Exists(path string) (bool, error) {
	path = normalizePath(path)
	info, err := fs.Stat(fxs.files, path)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return !info.IsDir(), nil
	}
}

func (fxs flexible) Open(path string) (fs.File, error) {
	path = normalizePath(path)
	return fxs.files.Open(path)
}

func (fxs flexible) ReadFile(path string) ([]byte, error) {
	path = normalizePath(path)
	return fs.ReadFile(fxs.files, path)
}

func (fxs flexible) Find(dir, pattern string) (*string, error) {
	var result string

	// Normalize
	dir = normalizePath(dir)

	// Create regex
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("invalid regexp pattern")
	}

	// Search for file
	err = fs.WalkDir(fxs.files, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !entry.IsDir() && rx.MatchString(entry.Name()) {
			result = normalizePath(path)
			return fs.SkipAll
		}

		return nil
	})

	// Handle result
	if err != nil {
		return nil, err
	} else if result == "" {
		return nil, nil
	}

	return &result, nil
}

func (fxs flexible) Search(dir, phrase, ignore, ext string) (*string, error) {
	var result string
	var err error
	var rxFind *regexp.Regexp
	var rxSkip *regexp.Regexp = nil

	// Normalize
	dir = normalizePath(dir)
	ext = strings.TrimLeft(ext, ".")

	// Create find regex
	if ext == "" {
		rxFind, err = regexp.Compile(phrase + ".*")
	} else {
		rxFind, err = regexp.Compile(phrase + `.*\.` + ext)
	}
	if err != nil {
		return nil, errors.New("invalid search pattern")
	}

	// Create skip regex
	if ignore != "" {
		rxSkip, err = regexp.Compile(".*" + ignore + ".*")
		if err != nil {
			return nil, errors.New("invalid ignore pattern")
		}
	}

	// Search for file
	err = fs.WalkDir(fxs.files, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !entry.IsDir() &&
			rxFind.MatchString(entry.Name()) &&
			(rxSkip == nil || !rxSkip.MatchString(entry.Name())) {
			result = normalizePath(path)
			return fs.SkipAll
		}

		return nil
	})

	// Handle result
	if err != nil {
		return nil, err
	} else if result == "" {
		return nil, nil
	}

	return &result, nil
}

func (fxs flexible) Lookup(dir, pattern string) ([]string, error) {
	var result []string

	// Normalize
	dir = normalizePath(dir)

	// Create regex
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, errors.New("invalid regexp pattern")
	}

	// Search for file
	err = fs.WalkDir(fxs.files, dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !entry.IsDir() && rx.MatchString(entry.Name()) {
			result = append(result, normalizePath(path))
		}

		return nil
	})

	// Handle result
	if err != nil {
		return nil, err
	} else if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}

func (fxs flexible) FS() fs.FS {
	return fxs.files
}

func (fxs flexible) Http() http.FileSystem {
	return http.FS(fxs.files)
}
