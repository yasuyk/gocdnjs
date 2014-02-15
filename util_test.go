package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCachePath(t *testing.T) {
	dir := "package"
	file := "package.json"
	path := []string{os.Getenv("HOME"), ".gocdnjs", "cache", dir, file}

	e := strings.Join(path, string(filepath.Separator))
	if r := CachePath(dir, file); e != r {
		t.Errorf("%s, want %s", r, e)
	}
}

func TestGenereateLink(t *testing.T) {
	link := GenereateLink("lib", "2.0", "lib.js")
	e := "http://cdnjs.cloudflare.com/ajax/libs/lib/2.0/lib.js"
	if link != e {
		t.Errorf("%s, want %s", link, e)
	}
}
