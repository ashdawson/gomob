package main

import (
	"encoding/json"
	"github.com/ashdawson/gomob/notif"
	"strings"
	"time"
)

var debug = false

func join() {
	if !isLastChangeSecondsAgo() {
		sayInfo("Actively waiting for new remote commit...")
	}
	for !isLastChangeSecondsAgo() {
		time.Sleep(time.Second)
		git("pull")
	}
}

func startSession() {
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
		git("branch", "--set-upstream-to="+settings.RemoteName+"/"+settings.BranchName, settings.BranchName)
	} else if !hasMobbingBranch() && !hasMobbingBranchOrigin() {
		sayInfo("create " + settings.BranchName + " from " + settings.BaseBranchName)
		git("checkout", settings.BaseBranchName)
		git("merge", settings.RemoteName+"/"+settings.BaseBranchName, "--ff-only")
		git("branch", settings.BranchName)
		git("checkout", settings.BranchName)
		git("push", "--set-upstream", settings.RemoteName, settings.BranchName)
	} else if !hasMobbingBranch() && hasMobbingBranchOrigin() {
		sayInfo("joining mob session")
		git("checkout", settings.BranchName)
		git("branch", "--set-upstream-to="+settings.RemoteName+"/"+settings.BranchName, settings.BranchName)
	} else {
		sayInfo("purging local branch and start new " + settings.BranchName + " branch from " + settings.BaseBranchName)
		git("branch", "-D", settings.BranchName)
		git("checkout", settings.BaseBranchName)
		git("merge", settings.RemoteName+"/"+settings.BaseBranchName, "--ff-only")
		git("branch", settings.BranchName)
		git("checkout", settings.BranchName)
		git("push", "--set-upstream", settings.RemoteName, settings.BranchName)
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
		git("commit", "--message", "\""+settings.CommitMessage+"\"")
		git("push")
	}
	showNext()
}

func getChangesOfLastCommit() string {
	return strings.TrimSpace(git("diff", "HEAD^1", "--stat"))
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
			git("commit", "--message", "\""+settings.CommitMessage+"\"")
		}
		git("push", settings.RemoteName, settings.BranchName)

		git("checkout", settings.BranchName)
		git("merge", settings.RemoteName+"/"+settings.BranchName, "--ff-only")
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

func status() {
	if isMobbing() {
		sayInfo("mobbing in progress")
	} else {
		sayInfo("you aren't mobbing right now")
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
	return strings.Contains(output, "  "+settings.RemoteName+"/"+settings.BranchName)
}

func isLastChangeSecondsAgo() bool {
	changes := git("--no-pager", "log", settings.RemoteName+".."+settings.BranchName, "--pretty=format:%cr", "--abbrev-commit")
	lines := strings.Split(strings.Replace(changes, "\r\n", "\n", -1), "\n")
	numberOfLines := len(lines)
	if numberOfLines < 1 {
		return true
	}

	return strings.Contains(lines[0], "seconds ago") || strings.Contains(lines[0], "second ago")
}

func showNext() {
	changes := strings.TrimSpace(git("--no-pager", "log", settings.BranchName, "--pretty=format:%an", "--abbrev-commit"))
	lines := strings.Split(strings.Replace(changes, "\r\n", "\n", -1), "\n")
	numberOfLines := len(lines)
	gitUserName := getGitUserName()
	if numberOfLines < 1 {
		return
	}
	var history = ""
	for i := 0; i < len(lines); i++ {
		if lines[i] == gitUserName && i > 0 {
			notif.Notify(lines[i-1] + " is (probably) next.")
			return
		}
		if history != "" {
			history = ", " + history
		}
		history = lines[i] + history
	}
}

func config() {
	say("config")
	s, _ := json.MarshalIndent(settings, "", "\t")
	say(string(s))
}

func help() {
	say("usage")
	say("\tmob [s]tart \t# start mobbing as typist")
	say("\tmob [j]oin \t# like start but waits for recent commit")
	say("\tmob [n]ext \t# hand over to next typist")
	say("\tmob [d]one \t# finish mob session")
	say("\tmob [r]eset \t# resets any unfinished mob session")
	say("\tmob status \t# show status of mob session")
	say("\tmob --help \t# prints this help")
	say("\tmob --version \t# prints the version")
}