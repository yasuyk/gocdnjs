gocdnjs
============

[![Build Status](https://travis-ci.org/yasuyk/gocdnjs.png?branch=master)](https://travis-ci.org/yasuyk/gocdnjs)
[![codecov](https://codecov.io/gh/yasuyk/gocdnjs/branch/master/graph/badge.svg)](https://codecov.io/gh/yasuyk/gocdnjs)

Command line interface for [cdnjs.com][cdnjs].

## Installation

    go get github.com/yasuyk/gocdnjs

## Usage

```
NAME:
   gocdnjs - command line interface for cdnjs.com

USAGE:
   gocdnjs [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   list, l      List available packages
   info, i      Show package information
   url, u       Show CDN URL of the package
   update       Update the package cache
   download, d  Download assets of the package
   cachefile, c Show cache file path
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --version, -v        print the version
   --help, -h           show help
```

[cdnjs]:http://cdnjs.com

## License

See [LICENSE][license].

[license]: https://github.com/yasuyk/gocdnjs/blob/master/LICENSE
