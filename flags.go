package main

import (
	"flag"
	"fmt"
	"os"
)

// Read reads the command line
func readCommandLineArguments() {
	startCommand = getStartCommand()
	flag.IntVar(&config.TimeLimit, "time_limit", config.TimeLimit, "mob session time")
	flag.StringVar(&config.IDE, "ide", config.IDE, "ide used to open files from commit messages")
	flag.StringVar(&config.BranchName, "branch", config.BranchName, "current branch you are working on")
	flag.StringVar(&config.RemoteName, "remote", config.RemoteName, "the name of the remote")
	flag.StringVar(&config.BaseBranchName, "base_branch", config.BaseBranchName, "current base branch you are working on")
	flag.StringVar(&config.BaseRemoteName, "base_remote", config.BaseRemoteName, "the name of the base remote")
	flag.StringVar(&config.CommitMessage, "commit_message", config.CommitMessage, "commit message appended to start of each commit")

	flag.BoolVar(&settings.TimerOnly, "timer", false, "set if you only want to use the app as a timer")
	flag.BoolVar(&settings.Debug, "debug", false, "display output of commands")

	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println("No command line options were found for: ", flag.Args())
	}

	config.save()
}

func getStartCommand() string {
	args := os.Args
	if len(args) >= 1 {
		os.Args = os.Args[1:]
		return args[1]
	}
	return "start"
}
