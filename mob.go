package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/ashdawson/gomob/flags"
)

var argsMap = map[string]string{}
var envVariables = map[string]string{
	"MOB_BRANCH":         "",
	"MOB_REMOTE":         "origin",
	"MOB_COMMIT_MESSAGE": "Mob Session COMPLETE [ci-skip]",
	"MOB_TIME_LIMIT":     "15",
	"MOB_TEAM":           "",
}

func setup() {
	parseEnvironmentVariables()
	flags.Read()
	getArguments()
	runCommands()
}

func parseEnvironmentVariables() {
	for envKey, _ := range envVariables {
		setEnvVarIfExists(envKey)
	}
}

func setEnvVarIfExists(key string) {
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

func runCommands() {
	for argKey, _ := range argsMap {
		switch argKey {
		case "start":
			config()
			break
		//case "start":
		//	start()
		//	status()
		//	break
		//case "join":
		//	start()
		//	join()
		//	status()
		//	break
		//case "next":
		//	next()
		//	break
		//case "done":
		//	done()
		//	break
		//case "help":
		//	help()
		//	break
		default:
			fmt.Println("OOPS")
		}
	}
}
