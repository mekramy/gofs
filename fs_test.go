package gofs_test

import (
	"embed"
	"testing"

	"github.com/mekramy/gofs"
)

//go:embed *.go
var embeds embed.FS

func TestDir(t *testing.T) {
	fs := gofs.NewDir(".")
	invalid := "missing.go"
	valid := "fs.go"

	t.Run("Exists", func(t *testing.T) {
		if ok, err := fs.Exists(invalid); err != nil {
			t.Fatal(err)
		} else if ok {
			t.Fatalf("%s should not exists!", invalid)
		}

		if ok, err := fs.Exists(valid); err != nil {
			t.Fatal(err)
		} else if !ok {
			t.Fatalf("%s should exists!", valid)
		}
	})

	t.Run("Open", func(t *testing.T) {
		if _, err := fs.Open(invalid); err == nil {
			t.Fatalf("%s should not be opened!", invalid)
		}

		if file, err := fs.Open(valid); err != nil {
			t.Fatal(err)
		} else {
			file.Close()
		}
	})

	t.Run("ReadFile", func(t *testing.T) {
		if _, err := fs.ReadFile(invalid); err == nil {
			t.Fatalf("%s should not be read!", invalid)
		}

		if content, err := fs.ReadFile(valid); err != nil {
			t.Fatal(err)
		} else if len(content) == 0 {
			t.Fatalf("%s should not be empty!", valid)
		}
	})

	t.Run("Search", func(t *testing.T) {
		dir := "."
		phrase := `fs\.driver`
		ignore := ""
		ext := "go"

		if result, err := fs.Search(dir, phrase, ignore, ext); err != nil {
			t.Fatal(err)
		} else if result == nil {
			t.Fatalf("Search should find a result!")
		}
	})

	t.Run("Find", func(t *testing.T) {
		dir := "."
		pattern := ".*driver.*"

		if result, err := fs.Find(dir, pattern); err != nil {
			t.Fatal(err)
		} else if result == nil {
			t.Fatalf("Find should find a result!")
		}
	})

	t.Run("Lookup", func(t *testing.T) {
		dir := "."
		pattern := `.*\.go`

		if results, err := fs.Lookup(dir, pattern); err != nil {
			t.Fatal(err)
		} else if len(results) == 0 {
			t.Fatalf("Lookup should find results!")
		}
	})

	t.Run("FS", func(t *testing.T) {
		if fs.FS() == nil {
			t.Fatalf("FS should return a valid fs.FS instance!")
		}
	})

	t.Run("Http", func(t *testing.T) {
		if fs.Http() == nil {
			t.Fatalf("Http should return a valid http.FileSystem instance!")
		}
	})
}

func TestEmbed(t *testing.T) {
	fs := gofs.NewEmbed(embeds)
	invalid := "missing.go"
	valid := "fs.go"

	t.Run("Exists", func(t *testing.T) {
		if ok, err := fs.Exists(invalid); err != nil {
			t.Fatal(err)
		} else if ok {
			t.Fatalf("%s should not exists!", invalid)
		}

		if ok, err := fs.Exists(valid); err != nil {
			t.Fatal(err)
		} else if !ok {
			t.Fatalf("%s should exists!", valid)
		}
	})

	t.Run("Open", func(t *testing.T) {
		if _, err := fs.Open(invalid); err == nil {
			t.Fatalf("%s should not be opened!", invalid)
		}

		if file, err := fs.Open(valid); err != nil {
			t.Fatal(err)
		} else {
			file.Close()
		}
	})

	t.Run("ReadFile", func(t *testing.T) {
		if _, err := fs.ReadFile(invalid); err == nil {
			t.Fatalf("%s should not be read!", invalid)
		}

		if content, err := fs.ReadFile(valid); err != nil {
			t.Fatal(err)
		} else if len(content) == 0 {
			t.Fatalf("%s should not be empty!", valid)
		}
	})

	t.Run("Search", func(t *testing.T) {
		dir := "."
		phrase := `fs\.driver`
		ignore := ""
		ext := "go"

		if result, err := fs.Search(dir, phrase, ignore, ext); err != nil {
			t.Fatal(err)
		} else if result == nil {
			t.Fatalf("Search should find a result!")
		}
	})

	t.Run("Find", func(t *testing.T) {
		dir := "."
		pattern := ".*driver.*"

		if result, err := fs.Find(dir, pattern); err != nil {
			t.Fatal(err)
		} else if result == nil {
			t.Fatalf("Find should find a result!")
		}
	})

	t.Run("Lookup", func(t *testing.T) {
		dir := "."
		pattern := `.*\.go`

		if results, err := fs.Lookup(dir, pattern); err != nil {
			t.Fatal(err)
		} else if len(results) == 0 {
			t.Fatalf("Lookup should find results!")
		}
	})

	t.Run("FS", func(t *testing.T) {
		if fs.FS() == nil {
			t.Fatalf("FS should return a valid fs.FS instance!")
		}
	})

	t.Run("Http", func(t *testing.T) {
		if fs.Http() == nil {
			t.Fatalf("Http should return a valid http.FileSystem instance!")
		}
	})
}
