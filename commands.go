package main

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
)

const packageNotFond = "package not found."

type Commands struct {
	PackagesCache *PackagesCache
	Packages      *Packages
}

func NewCommnads(cachePath string) *Commands {
	pj := NewPackagesCache(cachePath)
	pj.FetchIfNeeded()
	c := &Commands{pj, NewPackages(pj.PackagesCachePath)}
	return c
}

func (c *Commands) List(cli *cli.Context) {
	showDesc := cli.Bool("d")
	isPlist := cli.Bool("p")

	if cli.Bool("u") {
		c.Update()
	}

	ps := c.Packages
	if isPlist {
		fmt.Printf("(gocdnjs-version \"%s\" name-max-length %d packages (",
			AppVersion, ps.Len.NameMax)
		for _, p := range *ps.Packages {
			fmt.Printf("(name \"%s\" version \"%s\" ", p.Name, p.Version)
			if showDesc {
				fmt.Printf("description \"%s\"", TrimNewLine(p.Description))
			}
			fmt.Print(")")
		}
		fmt.Print("))")
	} else {
		for _, p := range *ps.Packages {
			fmt.Printf("%s%s (%s)",
				p.Name, strings.Repeat(" ", ps.Len.NameMax-len(p.Name)),
				p.Version)
			if showDesc {
				fmt.Printf(" %s", TrimNewLine(p.Description))
			}
			fmt.Println()
		}
	}
}

func (c *Commands) Search(str string) {
	ps := c.Packages
	ps = ps.Search(str)
	for _, p := range *ps.Packages {
		fmt.Printf("%s%s (%s)%s\n",
			p.Name, strings.Repeat(" ", ps.Len.NameMax-len(p.Name)),
			p.Version, strings.Repeat(" ", ps.Len.VersionMax-len(p.Version)))
	}
}

func (c *Commands) Info(cli *cli.Context) {
	name := cli.Args().First()
	isPlist := cli.Bool("p")

	if cli.Bool("u") {
		c.Update()
	}

	p := c.Packages.SearchWithName(name)
	if p == nil {
		fmt.Println(packageNotFond)
	} else {
		fmt.Println(infoString(p, isPlist))
	}
}

func (c *Commands) Url(cli *cli.Context) {
	name := cli.Args().First()

	if cli.Bool("u") {
		c.Update()
	}

	p := c.Packages.SearchWithName(name)
	if p == nil {
		fmt.Println(packageNotFond)
	} else {
		version := p.Version
		if v := cli.Args().Get(1); v != "" {
			version = v
		}
		link := GenereateLink(p.Name, version, p.FileName)
		fmt.Println(link)
	}
}

func (c *Commands) Update() {
	pj := c.PackagesCache
	pj.Purge()
	pj.Fetch()
}

func (c *Commands) Download(cli *cli.Context) {
	name := cli.Args().First()

	if p := c.Packages.SearchWithName(name); p == nil {
		fmt.Println(packageNotFond)
	} else {
		p.Download(cli.String("d"), cli.String("v"))
	}
}

func (c *Commands) CahcePath() {
	pj := c.PackagesCache
	fmt.Println(pj.PackagesCachePath)
}

func infoString(p *Package, isPlist bool) string {
	assets := (p.Assets)

	str := ""
	if isPlist {
		str += fmt.Sprintf("(package \"%s\" version \"%s\" ", p.Name, assets[0].Version)
		if p.Homepage != "" {
			str += fmt.Sprintf("homepage \"%s\" ", p.Homepage)
		}
		str += fmt.Sprintf("description \"%s\" ", p.Description)

		str += fmt.Sprint("assets (")
		for _, a := range assets {
			str += fmt.Sprintf("(\"%s\" (", a.Version)
			for _, f := range a.Files {
				link := GenereateLink(p.Name, a.Version, f)
				str += fmt.Sprintf("\"%s\" ", link)
			}
			str += fmt.Sprint("))")
		}
		str += fmt.Sprint("))")
	} else {
		str += fmt.Sprintf("%s\n\n", p.Name)
		if p.Homepage != "" {
			str += fmt.Sprintf("HOMEPAGE:\n   %s\n\n", p.Homepage)
		}
		str += fmt.Sprintf("DESCRIPTION:\n   %s\n\n", p.Description)
		str += fmt.Sprintln("ASSETS:")
		for _, a := range assets {
			for _, f := range a.Files {
				link := GenereateLink(p.Name, a.Version, f)
				str += fmt.Sprintf("   %s\n", link)
			}
			str += fmt.Sprintln()
		}
	}

	return str
}
