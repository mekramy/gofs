package gofs

import "path/filepath"

// normalizePath join and normalize file path.
func normalizePath(path ...string) string {
	return filepath.ToSlash(filepath.Clean(filepath.Join(path...)))
}
