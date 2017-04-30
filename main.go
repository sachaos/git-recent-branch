package main

import (
	"os"

	"encoding/csv"
	"github.com/sachaos/git-recent-branch/utils"
	"github.com/urfave/cli"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	BeforeHash string
	AfterHash  string
	Author     string
	Email      string
	Message    string
	CreatedAt  time.Time
	Command    string
}

type Logs []Log

func (a Logs) Len() int           { return len(a) }
func (a Logs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Logs) Less(i, j int) bool { return a[i].CreatedAt.Unix() > a[j].CreatedAt.Unix() }

var writer utils.Writer
var logRegex = regexp.MustCompile(`^([[:alnum:]]{40})\s([[:alnum:]]{40})\s(.*)\s\<(.*)\>\s(\d{10})\s([\+\-]\d{4})\s(.*):\s(.*)$`)
var checkoutRegex = regexp.MustCompile(`^moving\sfrom\s(.*)\sto\s(.*)$`)

func NewLog(logString string) Log {
	parsedLog := logRegex.FindAllStringSubmatch(logString, -1)[0]
	unix, err := strconv.ParseInt(parsedLog[5], 10, 64)

	if err != nil {
		panic("Failed parse Log")
	}
	return Log{
		BeforeHash: parsedLog[1],
		AfterHash:  parsedLog[2],
		Author:     parsedLog[3],
		Email:      parsedLog[4],
		Command:    parsedLog[7],
		Message:    parsedLog[8],
		CreatedAt:  time.Unix(unix, 0),
	}
}

func gitRecentBranch(c *cli.Context) {
	defer writer.Flush()

	logsBuf, err := ioutil.ReadFile(".git/logs/HEAD")
	if err != nil {
		panic("failed to open logs")
	}
	logStrings := strings.Split(strings.Trim(string(logsBuf), "\n"), "\n")

	var logs Logs

	for _, logString := range logStrings {
		if log := NewLog(logString); log.Command == "checkout" {
			logs = append(logs, log)
		}
	}

	sort.Sort(logs)

	for _, log := range logs {
		parsedMessage := checkoutRegex.FindAllStringSubmatch(log.Message, -1)[0]
		writer.Write([]string{
			parsedMessage[1],
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
