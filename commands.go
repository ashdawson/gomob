package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/clock"
	"github.com/ashdawson/gomob/notif"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var debug = true

func startSession() {
	if !isTimerOnly {
		openFiles()

		git("fetch", "--prune")
		git("pull")
	}

	sayInfo("session started")
	sessionStartTime = clock.New().Now()
	startTimer(settings.TimeLimit)
}

func next() {
	join()
	if !isTimerOnly {
		if !isMobbing() {
			sayError("you aren't mobbing")
			return
		}

		if !hasCommits() {
			sayInfo("nothing was done, so nothing to commit")
		} else {
			commit()
			git("push")
			sayInfo("changes pushed to " + getBranch())
		}
	}

	if getGitUserName() == showNext() {
		notif.Notify(showNext() + " is next.")
	}
}

func commitMessage() string {
	return settings.CommitMessage + getModifiedFiles()
}

func commit() {
	message := commitMessage()
	git("add", "--all")
	git("commit", "--message", message)
}

func showNext() string {
	changes := strings.TrimSpace(git("--no-pager", "log", getBranch(), "--pretty=format:%an", "--abbrev-commit"))
	lines := strings.Split(strings.Replace(changes, "\r\n", "\n", -1), "\n")
	numberOfLines := len(lines)
	gitUserName := getGitUserName()
	if numberOfLines < 1 {
		return ""
	}
	var history = ""
	for i := 0; i < len(lines); i++ {
		if lines[i] == gitUserName && i > 0 {
			return lines[i-1]
		}
		if history != "" {
			history = ", " + history
		}
		history = lines[i] + history
	}
	return ""
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
	say("\tmob [n]ext \t# hand over to next typist")
	say("\tmob [d]one \t# finish mob session")
	say("\tmob status \t# show status of mob session")
	say("\tmob --help \t# prints this help")
	say("\tmob --version \t# prints the version")
}

func openFiles() {
	// Currently only supports PHPStorm or VSCode
	supportedIDE := map[string]bool{
		"phpstorm": true,
		"vscode":   true,
	}
	app := strings.ToLower(settings.IDE)
	if len(getLastCommitMessage()) > 0 && supportedIDE[app] {
		if runtime.GOOS == "windows" && app == "phpstorm" {
			app = app + ".exe"
		}

		var command *exec.Cmd
		if app == "vscode" {
			app = "code"
		}
		command = exec.Command(app, getLastCommitMessage()...)
		err := command.Run()
		if err != nil {
			sayError(command.Args)
			sayError(err)
			os.Exit(1)
		}
	}
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	check(err)
	return dir
}
