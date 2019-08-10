package main

import (
	"github.com/ashdawson/gomob/command"
	"os"
	"strings"
)

var argsMap = map[string]string{}
var settings Settings
var envKey = "MOB"

type Settings struct {
	BaseBranchName string
	BaseRemoteName string
	BranchName     string
	RemoteName     string
	CommitMessage  string
	TimeLimit      int
	Mob            string
}

var envVariables = map[string]string{
	"BRANCH":         "",
	"REMOTE":         "origin",
	"COMMIT_MESSAGE": "Mob Session COMPLETE [ci-skip]",
	"TIME_LIMIT":     "15",
	"TEAM":           "",
}

func setup() {
	parseEnvironmentVariables()
	command.Read()
	getArguments()
}

func parseEnvironmentVariables() {
	for envKey := range envVariables {
		setEnvVarIfExists(envKey)
	}
}

func prependToString(prepend string, subject string) string {
	subject = prepend + subject
	return subject
}

func setEnvVarIfExists(key string) {
	prependToString(envVariables["ENV_KEY"], key)
	if localVar, ok := os.LookupEnv(key); ok {
		envVariables[key] = localVar
	}
}

func getArguments() {
	for i := 1; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "--") {
			if _, ok := argsMap[os.Args[i]]; !ok {
				var hasValue = strings.Index(os.Args[i], "=")
				if hasValue > 0 {
					argsMap[os.Args[i][2:hasValue]] = os.Args[i][hasValue:]
				} else {
					argsMap[os.Args[i][2:]] = ""
				}
			}
		}
	}
}
