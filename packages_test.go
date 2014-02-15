package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var packages *Packages

func init() {
	dir, _ := os.Getwd()
	filepath := filepath.Join(dir, "/test/package.json")
	packages = NewPackages(filepath)
}

func TestNewPackages(t *testing.T) {
	a := (*packages.Packages)[1]

	if e := *FixturePackage(); !reflect.DeepEqual(a, e) {
		t.Fatalf("Expected %v, but %v:", e, a)
	}
}

func TestPackageSearch(t *testing.T) {
	a := Package{"Chart", "", "0", "", "", []string{"chart", "Graph"}, "", nil, nil}

	if e := a.search("foo"); e != false {
		t.Errorf("Expected false, but %t", e)
	}
	if e := a.search("chart"); e != true {
		t.Errorf("Expected true, but %t", e)
	}
	if e := a.search("Char"); e != true {
		t.Errorf("Expected true, but %t", e)
	}
	if e := a.search("ph"); e != true {
		t.Errorf("Failed searching keywoards: Expected true, but %t", e)
	}
}

func TestPackagesSearch(t *testing.T) {
	a := Package{"Chart", "", "0", "", "", []string{"chart", "Graph"}, "", nil, nil}
	ps := &Packages{&[]Package{a}, Len{0, 0, 0}}

	e := &[]Package{}

	if a := ps.Search("foo").Packages; !reflect.DeepEqual(e, a) {
		t.Errorf("Expected %v, but (%v)", e, a)
	}

	a2 := (*ps.Search("chart").Packages)[0]

	if e2 := (*ps.Packages)[0]; !reflect.DeepEqual(e2, a2) {
		t.Fatalf("Expected %v, but %v:", e, a2)
	}
}

func TestPackagesSearchWithName(t *testing.T) {
	a := Package{"Chart", "", "0", "", "", []string{"chart", "Graph"}, "", nil, nil}
	ps := &Packages{&[]Package{a}, Len{0, 0, 0}}

	if e := ps.SearchWithName("Ch"); !reflect.DeepEqual(e, &a) {
		t.Errorf("Expected%s, want %s", &a, e)
	}

	if e := ps.SearchWithName("ha"); reflect.DeepEqual(e, &a) {
		t.Errorf("Expected that %s is not %s", &a, e)
	}

	if e := ps.SearchWithName("ha"); reflect.DeepEqual(e, &a) {
		t.Errorf("Expected that %s is not %s", &a, e)
	}
}
