package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "csv",
		Usage: "output in CSV format",
	},
	cli.IntFlag{
		Name:  "n",
		Value: 10,
		Usage: "num of entry",
	},
}

var Commands = []cli.Command{}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
