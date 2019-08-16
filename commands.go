package main

import (
	"encoding/json"
	"fmt"
	"github.com/ashdawson/gomob/clock"
	"github.com/ashdawson/gomob/notif"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func runCommands() {
	switch startCommand {
	case "config":
		config()
		break
	case "start":
		startSession()
		break
	case "join":
		joinSession()
		break
	case "next":
		next()
		break
	case "driver":
		notif.Notify(getNextDriver() + " is next.")
		break
	case "help":
		help()
		break
	default:
		fmt.Println("OOPS")
	}
	wg.Wait()
}

func startSession() {
	if !isTimerOnly {
		openFiles()

		git("fetch", "--prune")
		git("pull")
	}

	sayInfo(fmt.Sprintf("Session started (%d minutes)", settings.TimeLimit))
	sessionStartTime = clock.New().Now()
	startTimer(settings.TimeLimit)
}

func joinSession() {
	sayInfo("Tracking changes to: " + getBranch())
	join()
}

func next() {
	if !isTimerOnly {
		if hasCommits() {
			commit()
		} else {
			commitWithEmpty()
		}
	}

	git("push")
	sayInfo("Changes pushed to " + getBranch())
	notif.Notify(getNextDriver() + " is next.")

	join()
}

func config() {
	say("config")
	s, _ := json.MarshalIndent(settings, "", "\t")
	say(string(s))
}

func help() {
	say("usage")
	say("\tmob [s]tart \t# start mobbing")
	say("\tmob [j]oin \t# like start but waits for recent commit")
	say("\tmob [d]river \t# display the next driver")
	say("\tmob [c]onfig \t# display the config")
	say("\tmob [r]eset \t# clear your settings")
	say("\tmob --help \t# prints this help")
}

func openFiles() {
	// Currently only supports PHPStorm or VSCode
	supportedIDE := map[string]bool{
		"phpstorm": true,
		"vscode":   true,
	}
	app := strings.ToLower(settings.IDE)
	lastMessage := getLastCommitMessage()
	committedFiles := getLastCommittedFiles(lastMessage)
	if len(committedFiles) > 0 && supportedIDE[app] && !strings.Contains(lastMessage, "Empty commit") {
		if runtime.GOOS == "windows" && app == "phpstorm" {
			const systemType = 32 << (^uint(0) >> 32 & 1)
			app = app + strconv.Itoa(systemType)
		}

		if app == "vscode" {
			app = "code"
		}

		var command *exec.Cmd
		command = exec.Command(app, committedFiles...)
		err := command.Run()
		if err != nil {
			sayError(command.Args)
			sayError(err)
		}
	}
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	check(err)
	return dir
}
