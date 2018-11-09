# git-recent-branch

## Description

git subcommand.
List recent visited branches with visit time.

[![asciicast](https://asciinema.org/a/2ivqfsaz672d36l596tdop51l.png)](https://asciinema.org/a/2ivqfsaz672d36l596tdop51l)

## Usage

```
$ git recent-branch -h
NAME:
   git-recent-branch

USAGE:
   git-recent-branch [global options] command [command options] [arguments...]

VERSION:
   0.2.0

AUTHOR:
   sachaos

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --csv          output in CSV format
   --no-unique    show not unique logs
   -n value       num of entry (default: 10)
   --help, -h     show help
   --version, -v  print the version
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/sachaos/git-recent-branch
```

### Use with peco

**RECOMMENDED**

install *peco* and load `git-recent-branch_functions.sh` on your `.zshrc`, like below.

```
$ source "$GOPATH/src/github.com/sachaos/git-recent-branch/todoist_functions.sh"
```

#### keybind

```
<C-g> <C-r>: select branch by peco, and insert command line buffer.
```

## Contribution

1. Fork ([https://github.com/sachaos/git-recent-branch/fork](https://github.com/sachaos/git-recent-branch/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[sachaos](https://github.com/sachaos)
