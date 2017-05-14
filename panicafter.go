package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const defaultSeconds = 10
const errorMsg = `error: invalid seconds`
const helpMsg = `usage: %s [seconds=10]

Panic After - throw exceptions after some time\n`

// PanicAfter panics after a number of seconds
func PanicAfter(s int) {
	time.Sleep(time.Second * time.Duration(s))
	panic("Oops")
}

func parseArgs() (int, error) {
	if len(os.Args) > 1 {
		s, err := strconv.Atoi(os.Args[1])
		return s, err
	}
	return defaultSeconds, nil
}

func parseHelp() bool {
	if len(os.Args) > 1 {
		if os.Args[1] == "?" || os.Args[1] == "help" {
			return true
		}
	}
	return false
}

func printHelp() {
	fmt.Printf(helpMsg, os.Args[0])
}

func printError() {
	fmt.Println(errorMsg)
}

func main() {
	if parseHelp() {
		printHelp()
		os.Exit(0)
	}

	seconds, err := parseArgs()
	if err != nil {
		printError()
		printHelp()
		os.Exit(0)
	}
	PanicAfter(seconds)
}
