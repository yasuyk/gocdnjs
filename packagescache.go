package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type PackagesCache struct {
	PackagesCachePath string
}

var PackagesCacheUrl = "http://cdnjs.com/packages.json"

func NewPackagesCache(root string) *PackagesCache {
	p := PackagesCache{filepath.Join(root, "package.json")}
	return &p
}

func (p *PackagesCache) isChached() bool {
	return Exists(p.PackagesCachePath)
}

func (p *PackagesCache) cacheFor() []byte {
	content, err := ioutil.ReadFile(p.PackagesCachePath)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func (p *PackagesCache) FetchIfNeeded() {
	if !p.isChached() {
		p.Fetch()
	}
}

func (p *PackagesCache) Fetch() []byte {

	if p.isChached() {
		return p.cacheFor()
	}

	output := HttpGetPackage(PackagesCacheUrl)

	dir, _ := filepath.Split(p.PackagesCachePath)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(p.PackagesCachePath, output, 0600)

	return output
}

func (p *PackagesCache) Purge() {
	path := p.PackagesCachePath
	if Exists(path) {
		if err := os.Remove(path); err != nil {
			log.Fatal(err)
		}
	}
}
