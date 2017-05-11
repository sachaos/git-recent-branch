package main

import (
	"os"

	"encoding/csv"
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"github.com/sachaos/git-recent-branch/gitlogs"
	"github.com/sachaos/git-recent-branch/utils"
	"github.com/urfave/cli"
	"io/ioutil"
	"os/exec"
	"sort"
	"strings"
	"time"
)

var writer utils.Writer

func contains(s []string, x string) bool {
	for _, a := range s {
		if a == x {
			return true
		}
	}
	return false
}

func gitRoot() string {
	buf, err := exec.Command("git-rev-parse", "--show-cdup").Output()
	if err != nil {
		fmt.Printf("%s", err)
		panic(err)
	}

	return strings.TrimSpace(string(buf))
}

func gitRecentBranch(c *cli.Context) {
	defer writer.Flush()

	logsBuf, err := ioutil.ReadFile(gitRoot() + ".git/logs/HEAD")
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

	branches := []string{}
	out := [][]string{}

	for _, log := range logs {
		beforeBranche := log.Message.(gitlogs.CheckoutLog).BeforeCommit

		if !c.Bool("no-unique") && contains(branches, beforeBranche) {
			continue
		}

		out = append(out, []string{
			beforeBranche,
			"(" + humanize.Time(log.CreatedAt) + ")",
			log.CreatedAt.Format(time.UnixDate),
		})

		branches = append(branches, beforeBranche)
	}

	num := c.Int("n")
	lnum := len(out)

	for i := 0; i < num && i < lnum; i++ {
		writer.Write(out[i])
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
