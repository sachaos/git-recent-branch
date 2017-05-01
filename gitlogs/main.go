package gitlogs

import (
	"regexp"
	"strconv"
	"time"
)

type Log struct {
	BeforeHash string
	AfterHash  string
	Author     string
	Email      string
	Message    interface{}
	CreatedAt  time.Time
	Command    string
}

type Logs []Log

func (a Logs) Len() int           { return len(a) }
func (a Logs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Logs) Less(i, j int) bool { return a[i].CreatedAt.Unix() > a[j].CreatedAt.Unix() }

var logRegex = regexp.MustCompile(`^([[:alnum:]]{40})\s([[:alnum:]]{40})\s(.*)\s\<(.*)\>\s(\d{10})\s([\+\-]\d{4})\s(.*):\s(.*)$`)
var checkoutRegex = regexp.MustCompile(`^moving\sfrom\s(.*)\sto\s(.*)$`)

type CheckoutLog struct {
	BeforeCommit string
	AfterCommit  string
}

func NewCheckoutLog(message string) CheckoutLog {
	parsedLog := checkoutRegex.FindAllStringSubmatch(message, -1)[0]
	return CheckoutLog{
		BeforeCommit: parsedLog[1],
		AfterCommit:  parsedLog[2],
	}
}

func NewLog(logString string) Log {
	parsedLog := logRegex.FindAllStringSubmatch(logString, -1)[0]
	unix, err := strconv.ParseInt(parsedLog[5], 10, 64)

	if err != nil {
		panic("Failed parse Log")
	}
	var message interface{}
	switch parsedLog[7] {
	case "checkout":
		message = NewCheckoutLog(parsedLog[8])
	default:
		message = parsedLog[8]
	}
	return Log{
		BeforeHash: parsedLog[1],
		AfterHash:  parsedLog[2],
		Author:     parsedLog[3],
		Email:      parsedLog[4],
		Command:    parsedLog[7],
		Message:    message,
		CreatedAt:  time.Unix(unix, 0),
	}
}
