package main

import (
	"flag"
	"fmt"
	"os"
)

// Read reads the command line
func readCommandLineArguments() {
	startCommand = getStartCommand()
	flag.IntVar(&settings.TimeLimit, "time_limit", settings.TimeLimit, "mob session time")
	flag.BoolVar(&isTimerOnly, "timer", isTimerOnly, "set if you only want to use the app as a timer")
	flag.BoolVar(&settings.Debug, "debug", settings.Debug, "display output of commands")
	flag.StringVar(&settings.IDE, "ide", settings.IDE, "ide used to open files from commit messages")
	flag.StringVar(&settings.BranchName, "branch", settings.BranchName, "current branch you are working on")
	flag.StringVar(&settings.RemoteName, "remote", settings.RemoteName, "the name of the remote")
	flag.StringVar(&settings.BaseBranchName, "base_branch", settings.BaseBranchName, "current base branch you are working on")
	flag.StringVar(&settings.BaseRemoteName, "base_remote", settings.BaseRemoteName, "the name of the base remote")
	flag.StringVar(&settings.CommitMessage, "commit_message", settings.CommitMessage, "commit message appended to start of each commit")

	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println("No command line options were found for: ", flag.Args())
	}

	saveSettings()
}

func getStartCommand() string {
	args := os.Args
	if len(args) >= 1 {
		os.Args = os.Args[1:]
		return args[1]
	}
	return "start"
}
