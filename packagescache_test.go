package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewPackagesCache(t *testing.T) {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "/test/")
	pj := NewPackagesCache(path)
	expextd := filepath.Join(path, "package.json")
	if pj.PackagesCachePath != expextd {
		t.Errorf("%s, want %s", pj.PackagesCachePath, expextd)
	}
}
