package main

import (
	"os"

	"encoding/csv"
	"github.com/sachaos/git-recent-branch/gitlogs"
	"github.com/sachaos/git-recent-branch/utils"
	"github.com/urfave/cli"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

var writer utils.Writer

func gitRecentBranch(c *cli.Context) {
	defer writer.Flush()

	logsBuf, err := ioutil.ReadFile(".git/logs/HEAD")
	if err != nil {
		panic("failed to open logs")
	}
	logStrings := strings.Split(strings.Trim(string(logsBuf), "\n"), "\n")

	var logs gitlogs.Logs

	for _, logString := range logStrings {
		if log := gitlogs.NewLog(logString); log.Command == "checkout" {
			logs = append(logs, log)
		}
	}

	sort.Sort(logs)

	for _, log := range logs {
		writer.Write([]string{
			log.Message.(gitlogs.CheckoutLog).BeforeCommit,
			"(" + time.Since(log.CreatedAt).String() + ")",
			log.CreatedAt.Format(time.UnixDate),
		})
	}
}

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "sachaos"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("csv") {
			writer = csv.NewWriter(os.Stdout)
		} else {
			writer = utils.NewTSVWriter(os.Stdout)
		}
		return nil
	}

	app.Action = gitRecentBranch

	app.Run(os.Args)
}
