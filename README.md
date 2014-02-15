gocdnjs [![Build Status](https://travis-ci.org/yasuyk/gocdnjs.png)](https://travis-ci.org/yasuyk/gocdnjs)
============

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
   list, l      List available libraries
   info, i      Show library information
   tag, t       Generate HTML tag
   update, u    Update the package cache
   download, d  Download assets
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
