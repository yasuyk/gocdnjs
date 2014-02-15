package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Packages struct {
	Packages *[]Package
	Len      Len
}

type Len struct {
	NameMax     int
	VersionMax  int
	KeywordsMax int
}

func (ps *Packages) Search(word string) *Packages {

	ret := []Package{}
	l := Len{0, 0, 0}

	for _, p := range *ps.Packages {
		if p.search(word) {
			ret = append(ret, p)
			l.replace(len(p.Name), len(p.Version), len(p.joinendKeywords()))
		}
	}
	return &Packages{&ret, l}
}

func (ps *Packages) SearchWithName(name string) *Package {

	for _, p := range *ps.Packages {
		if find(p.Name, name) {
			return &p
		}
	}
	return nil
}

func find(name, search string) bool {
	r, err := regexp.Compile("\\A" + search)
	if err == nil {
		if r.MatchString(name) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

type PackageSlice []Package

func (p PackageSlice) Len() int {
	return len(p)
}

func (p PackageSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PackageSlice) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

type Package struct {
	Name         string
	FileName     string
	Version      string
	Description  string
	Homepage     string
	Keywords     []string
	Licence      string
	Repositories []Repositories
	Assets       []Assets
}

func (p *Package) joinendKeywords() string {
	return strings.Join(p.Keywords, ",")
}

func (p *Package) search(sbstr string) bool {
	sbstr = strings.ToLower(sbstr)

	if contains(p.Name, sbstr) {
		return true
	} else if contains(p.Description, sbstr) {
		return true
	} else if p.containsKeywords(sbstr) {
		return true
	}
	return false
}

func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), substr)
}

func (p *Package) containsKeywords(substr string) bool {
	for _, k := range p.Keywords {
		if contains(k, substr) {
			return true
		}
	}
	return false
}

func (p *Package) Download(dstDir, version string) {

	a := AssetsSlice(p.Assets).SelectAssets(version)
	if a == nil {
		Exit(fmt.Sprintf("version %s is not found", version))
		return
	}

	if dstDir == "" {
		dstDir = "."
	}

	libdir := filepath.Join(dstDir, p.Name, a.Version)
	if err := os.MkdirAll(libdir, 0700); err != nil {
		log.Fatal(err)
	}

	for _, f := range a.Files {
		output := HttpGetPackage(GenereateLink(p.Name, a.Version, f))
		path := filepath.Join(libdir, f)

		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(path, output, 0600); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s downloaded\n", filepath.Join(libdir, f))
	}
}

type Repositories struct {
	Type string
	Url  string
}

type Assets struct {
	Version string
	Files   []string
}

type AssetsSlice []Assets

func (as AssetsSlice) SelectAssets(ver string) *Assets {

	if ver != "" && ver != Latest {
		for _, a := range as {
			if a.Version == ver {
				return &a
			}
		}
		return nil
	}
	return &(as[0])
}

type jsonRoot struct {
	Packages []interface{}
}

func NewPackages(filePath string) *Packages {
	content, e := ioutil.ReadFile(filePath)
	if e != nil {
		log.Fatal(e)
	}
	var d jsonRoot
	if e := json.Unmarshal(content, &d); e != nil {
		log.Fatal(e)
	}

	ps := []Package{}
	l := Len{0, 0, 0}

	for _, e := range d.Packages {
		p := newPackage(e)
		ps = append(ps, *p)
		l.replace(len(p.Name), len(p.Version), len(p.joinendKeywords()))
	}
	sort.Sort(PackageSlice(ps))
	packages := Packages{&ps, l}

	return &packages
}

func (l *Len) replace(num, ver, keys int) {
	if l.NameMax < num {
		l.NameMax = num
	}
	if l.VersionMax < ver {
		l.VersionMax = ver
	}

	if l.KeywordsMax < keys {
		l.KeywordsMax = keys
	}
}

func newPackage(json interface{}) *Package {
	switch data := json.(type) {
	case map[string]interface{}:
		pack := new(Package)
		if name := maybeString(data["name"]); name != "" {
			pack.Name = name
		}
		if filename := maybeString(data["filename"]); filename != "" {
			pack.FileName = filename
		}
		if version := maybeString(data["version"]); version != "" {
			pack.Version = version
		}
		if description := maybeString(data["description"]); description != "" {
			pack.Description = description
		}
		if homepage := maybeString(data["homepage"]); homepage != "" {
			pack.Homepage = homepage
		}
		if keywords := data["keywords"]; keywords != nil && isInterfaceSlice(keywords) {
			keys := []string{}
			for _, v := range keywords.([]interface{}) {
				keys = append(keys, maybeString(v))
			}
			pack.Keywords = keys
		}
		setRepositories(pack, data)

		if assets := data["assets"]; assets != nil {
			if a := newAssetsSlice(assets); a != nil {
				pack.Assets = a
			}
		}
		return pack
	default:
		return nil
	}
}

func maybeString(json interface{}) string {
	switch data := json.(type) {
	case string:
		return data
	}
	return ""
}

func setRepositories(p *Package, data map[string]interface{}) {
	if repos := data["repositories"]; repos != nil {
		if r := newRepository(repos); r != nil {
			p.Repositories = []Repositories{*r}
		}
	} else if repos := data["repository"]; repos != nil {
		if r := newRepository(repos); r != nil {
			p.Repositories = []Repositories{*r}
		}
	}
}

func newRepository(json interface{}) *Repositories {
	switch d := json.(type) {
	case map[string]interface{}:
		repos := new(Repositories)
		if t := maybeString(d["type"]); t != "" {
			repos.Type = t
		}
		if url := maybeString(d["url"]); url != "" {
			repos.Url = url
		}
		return repos
	case []interface{}:
		if len(d) != 0 {
			return newRepository(d[0])
		} else {
			return nil
		}
	default:
		return nil
	}
}

func newAssetsSlice(json interface{}) []Assets {
	switch data := json.(type) {
	case []interface{}:
		assetsSlice := make([]Assets, 0)
		for _, v := range data {
			assetsSlice = append(assetsSlice, *newAssets(v))
		}
		return assetsSlice
	default:
		return nil
	}
}

func isInterfaceSlice(json interface{}) bool {
	switch json.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func newAssets(json interface{}) *Assets {
	switch d := json.(type) {
	case map[string]interface{}:
		assets := new(Assets)
		if v := maybeString(d["version"]); v != "" {
			assets.Version = v
		}

		if files := d["files"]; files != "" && isInterfaceSlice(files) {
			assets.Files = []string{}
			for _, v := range files.([]interface{}) {
				assets.Files = append(assets.Files, maybeString(v))
			}
		}
		return assets
	default:
		return nil
	}
}
