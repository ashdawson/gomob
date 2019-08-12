package main

import (
	"encoding/json"
	"fmt"
	"github.com/ashdawson/gomob/notif"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var debug = false

func startSession() {
	openFiles()

	git("fetch", "--prune")
	git("pull")

	sayInfo("session started")
	startTimer(settings.TimeLimit)
}

func next() {
	join()
	if !isMobbing() {
		sayError("you aren't mobbing")
		return
	}

	if isNothingToCommit() {
		sayInfo("nothing was done, so nothing to commit")
	} else {
		commit()
		git("push")
		sayInfo("changes pushed to " + getBranch())
	}

	if getGitUserName() == showNext() {
		notif.Notify(showNext() + " is next.")
	}
}

func commitMessage() string {
	return settings.CommitMessage + getModifiedFiles()
}

func getCachedChanges() string {
	return strings.TrimSpace(git("diff", "--cached", "--stat"))
}

func commit() {
	message := commitMessage()
	git("add", "--all")
	git("commit", "--message", message)
}

func isNothingToCommit() bool {
	output := git("status", "--short")
	isMobbing := len(strings.TrimSpace(output)) == 0
	return isMobbing
}

func isMobbing() bool {
	output := git("branch")
	return strings.Contains(output, "* "+settings.BranchName)
}

func hasMobbingBranch() bool {
	output := git("branch")
	return strings.Contains(output, "  "+settings.BranchName) || strings.Contains(output, "* "+settings.BranchName)
}

func hasMobbingBranchOrigin() bool {
	output := git("branch", "--remotes")
	return strings.Contains(output, "  "+getBranch())
}

func isLastChangeSecondsAgo() bool {
	changes := git("--no-pager", "log", getBranch(), "--pretty=format:%cr", "--abbrev-commit")
	lines := strings.Split(strings.Replace(changes, "\r\n", "\n", -1), "\n")
	numberOfLines := len(lines)
	if numberOfLines < 1 {
		return true
	}

	return strings.Contains(lines[0], "seconds ago") || strings.Contains(lines[0], "second ago")
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
	app := strings.ToLower(settings.IDE)

	if runtime.GOOS == "windows" {
		app = app + ".exe"
	}

	if app == "vscode" {
		app = "code -g"
	}

	fmt.Println(app + " " + getLastCommitMessage())

	exec.Command(app + " " + getLastCommitMessage())
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	check(err)
	return dir
}