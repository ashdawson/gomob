package main

import (
	"os"
)

func setup() {
	runChecks()
	checkSettings()
	readCommandLineArguments()
}

func runChecks() {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		checkSay(err, "git has not been added to this directory")
	}
}
