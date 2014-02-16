package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const (
	// AppName is application name
	AppName = "gocdnjs"
	// Latest indicates latest version
	Latest                       = "latest"
	packageArg                   = "[package]"
	downloadPrefixflagDescripton = `Set directory prefix to prefix.
	The directory prefix is the directory where all other files
	and subdirectories will be saved to, i.e. the top of the retrieval tree.
	The default is . (the current directory).`
)

var plistOption, updateOption cli.BoolFlag

func init() {
	plistOption = cli.BoolFlag{Name: "plist, p", Usage: "print as a Property List of Lisp"}
	updateOption = cli.BoolFlag{Name: "update, u ", Usage: "with updating the package cache"}
	cli.CommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   command {{.Name}} [command options] {{.Description}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
}

func main() {
	lazycmd := lazyCmd()

	app := cli.NewApp()
	app.Name = AppName
	app.Usage = "command line interface for cdnjs.com"
	app.Version = AppVersion
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Commands = *commands(lazycmd)

	app.Run(os.Args)
}

func commands(cmd lazyLoadCmd) *[]cli.Command {
	return &[]cli.Command{
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List available packages",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "description, d", Usage: "Show description"},
				plistOption, updateOption,
			},
			Action: func(c *cli.Context) { cmd().List(c) },
		},
		{
			Name:        "info",
			ShortName:   "i",
			Usage:       "Show package information",
			Description: packageArg,
			Flags:       []cli.Flag{plistOption, updateOption},
			Action:      func(c *cli.Context) { cmd().Info(c) },
		},
		{
			Name:        "url",
			ShortName:   "u",
			Usage:       "Show CDN URL of the package",
			Description: packageArg,
			Flags:       []cli.Flag{updateOption},
			Action:      func(c *cli.Context) { cmd().Url(c) },
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "Update the package cache",
			Action:    func(c *cli.Context) { cmd().Update() },
		},
		{
			Name:        "download",
			ShortName:   "d",
			Usage:       "Download assets of the package",
			Description: packageArg,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "version, v", Value: Latest, Usage: "specify version"},
				cli.StringFlag{Name: "directory-prefix, d", Value: ".", Usage: downloadPrefixflagDescripton},
				updateOption},
			Action: func(c *cli.Context) { cmd().Download(c) },
		},
		{
			Name:      "cachefile",
			ShortName: "c",
			Usage:     "Show cache file path",
			Action:    func(c *cli.Context) { cmd().CahcePath() },
		},
	}
}

type lazyLoadCmd func() *Commands

func lazyCmd() lazyLoadCmd {
	var cmd *Commands
	return func() *Commands {
		if cmd == nil {
			cmd = NewCommnads(CachePath())
		}
		return cmd
	}
}
