package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/notif"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var debug = false

func startSession() {
	openFiles()
	sayInfo("session started")
	startTimer(settings.TimeLimit)

	if !isNothingToCommit() {
		notif.Notify("You have uncommitted changes")
		return
	}

	git("fetch", "--prune")
	git("pull")

	if hasMobbingBranch() && hasMobbingBranchOrigin() {
		sayInfo("rejoining mob session")
		git("branch", "-D", settings.BranchName)
		git("checkout", settings.BranchName)
		git("branch", "--set-upstream-to="+getBranch(), settings.BranchName)
	} else {
		sayInfo("joining mob session")
		git("checkout", settings.BranchName)
		git("branch", "--set-upstream-to="+getBranch(), settings.BranchName)
	}
}

func next() {
	if !isMobbing() {
		sayError("you aren't mobbing")
		return
	}

	if isNothingToCommit() {
		sayInfo("nothing was done, so nothing to commit")
	} else {
		git("add", "--all")
		git("commit", "--message", commitMessage())
		git("push")
		sayInfo("changes pushed to " + getBranch())
	}

	if getGitUserName() == showNext() {
		notif.Notify(showNext() + " is next.")
	}
}

func commitMessage() string {
	commitMessage := settings.CommitMessage + getChangedFiles(false)
	say(commitMessage)
	return commitMessage
}

func getCachedChanges() string {
	return strings.TrimSpace(git("diff", "--cached", "--stat"))
}

func done() {
	if !isMobbing() {
		sayError("you aren't mobbing")
		return
	}

	git("fetch", "--prune")

	if hasMobbingBranchOrigin() {
		if !isNothingToCommit() {
			git("add", "--all")
			git("commit", "--message", commitMessage())
		}
		git("push", settings.RemoteName, settings.BranchName)

		git("checkout", settings.BranchName)
		git("merge", getBranch(), "--ff-only")
		git("merge", "--squash", settings.BranchName)

		git("branch", "-D", settings.BranchName)
		git("push", settings.RemoteName, "--delete", settings.BranchName)
		say(getCachedChanges())
		sayTodo("git commit -m 'describe the changes'")
	} else {
		git("checkout", settings.BranchName)
		git("branch", "-D", settings.BranchName)
		sayInfo("someone else already ended your mob session")
	}
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
	app := strings.ToLower(settings.IDE)
	if runtime.GOOS == "windows" {
		app = app + ".exe"
	}
	exec.Command(app + getChangedFiles(true))
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	check(err)
	return dir
}