package main

import (
	"fmt"
	"os"
	"strings"
)

var argsMap = map[string]string{}
var envVariables = map[string]string{
	"MOB_BRANCH":         "",
	"MOB_REMOTE":         "origin",
	"MOB_COMMIT_MESSAGE": "Mob Session COMPLETE [ci-skip]",
	"MOB_TIME_LIMIT":     "15",
	"MOB_TEAM":           "",
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

func main() {
	parseEnvironmentVariables()
	getArguments()
	runCommand()
}

func runCommand() {
	for argKey, _ := range argsMap {
		switch argKey {
		case "start":
			config()
			break;
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

//
//func join() {
//	if !isLastChangeSecondsAgo() {
//		sayInfo("Actively waiting for new remote commit...")
//	}
//	for !isLastChangeSecondsAgo() {
//		time.Sleep(time.Second)
//		git("pull")
//	}
//}
//
//func startTimer(timerInMinutes string) {
//	timeoutInMinutes, _ := strconv.Atoi(timerInMinutes)
//	timeoutInSeconds := timeoutInMinutes * 60
//	timerInSeconds := strconv.Itoa(timeoutInSeconds)
//
//	command := exec.Command("sh", "-c", "( sleep "+timerInSeconds+" && say \"time's up\" && (/usr/bin/osascript -e 'display notification \"time is up\"' || /usr/bin/notify-send \"time is up\")  & )")
//	if debug {
//		fmt.Println(command.Args)
//	}
//	err := command.Start()
//	if err != nil {
//		sayError("timer couldn't be started... (timer only works on OSX)")
//		sayError(err)
//	} else {
//		timeOfTimeout := time.Now().Add(time.Minute * time.Duration(timeoutInMinutes)).Format("15:04")
//		sayOkay(timerInMinutes + " minutes timer started (finishes at approx. " + timeOfTimeout + ")")
//	}
//}
//
//func start() {
//	if !isNothingToCommit() {
//		sayNote("uncommitted changes")
//		return
//	}
//
//	git("fetch", "--prune")
//	git("pull")
//
//	if hasMobbingBranch() && hasMobbingBranchOrigin() {
//		sayInfo("rejoining mob session")
//		git("branch", "-D", wipBranch)
//		git("checkout", wipBranch)
//		git("branch", "--set-upstream-to="+remoteName+"/"+wipBranch, wipBranch)
//	} else if !hasMobbingBranch() && !hasMobbingBranchOrigin() {
//		sayInfo("create " + wipBranch + " from " + baseBranch)
//		git("checkout", baseBranch)
//		git("merge", remoteName+"/"+baseBranch, "--ff-only")
//		git("branch", wipBranch)
//		git("checkout", wipBranch)
//		git("push", "--set-upstream", remoteName, wipBranch)
//	} else if !hasMobbingBranch() && hasMobbingBranchOrigin() {
//		sayInfo("joining mob session")
//		git("checkout", wipBranch)
//		git("branch", "--set-upstream-to="+remoteName+"/"+wipBranch, wipBranch)
//	} else {
//		sayInfo("purging local branch and start new " + wipBranch + " branch from " + baseBranch)
//		git("branch", "-D", wipBranch) // check if unmerged commits
//
//		git("checkout", baseBranch)
//		git("merge", remoteName+"/"+baseBranch, "--ff-only")
//		git("branch", wipBranch)
//		git("checkout", wipBranch)
//		git("push", "--set-upstream", remoteName, wipBranch)
//	}
//}
//
//func next() {
//	if !isMobbing() {
//		sayError("you aren't mobbing")
//		return
//	}
//
//	if isNothingToCommit() {
//		sayInfo("nothing was done, so nothing to commit")
//	} else {
//		git("add", "--all")
//		git("commit", "--message", "\""+wipCommitMessage+"\"")
//		changes := getChangesOfLastCommit()
//		git("push", remoteName, wipBranch)
//		say(changes)
//	}
//	showNext()
//
//	git("checkout", baseBranch)
//}
//
//func getChangesOfLastCommit() string {
//	return strings.TrimSpace(silentgit("diff", "HEAD^1", "--stat"))
//}
//
//func getCachedChanges() string {
//	return strings.TrimSpace(silentgit("diff", "--cached", "--stat"))
//}
//
//func done() {
//	if !isMobbing() {
//		sayError("you aren't mobbing")
//		return
//	}
//
//	git("fetch", "--prune")
//
//	if hasMobbingBranchOrigin() {
//		if !isNothingToCommit() {
//			git("add", "--all")
//			git("commit", "--message", "\""+wipCommitMessage+"\"")
//		}
//		git("push", remoteName, wipBranch)
//
//		git("checkout", baseBranch)
//		git("merge", remoteName+"/"+baseBranch, "--ff-only")
//		git("merge", "--squash", wipBranch)
//
//		git("branch", "-D", wipBranch)
//		git("push", remoteName, "--delete", wipBranch)
//		say(getCachedChanges())
//		sayTodo("git commit -m 'describe the changes'")
//	} else {
//		git("checkout", baseBranch)
//		git("branch", "-D", wipBranch)
//		sayInfo("someone else already ended your mob session")
//	}
//}
//
//func status() {
//	if isMobbing() {
//		sayInfo("mobbing in progress")
//
//		output := silentgit("--no-pager", "log", baseBranch+".."+wipBranch, "--pretty=format:%h %cr <%an>", "--abbrev-commit")
//		say(output)
//	} else {
//		sayInfo("you aren't mobbing right now")
//	}
//
//	if !hasSay() {
//		sayNote("text-to-speech disabled because 'say' not found")
//	}
//}
//
//func isNothingToCommit() bool {
//	output := silentgit("status", "--short")
//	isMobbing := len(strings.TrimSpace(output)) == 0
//	return isMobbing
//}
//
//func isMobbing() bool {
//	output := silentgit("branch")
//	return strings.Contains(output, "* "+wipBranch)
//}
//
//func hasMobbingBranch() bool {
//	output := silentgit("branch")
//	return strings.Contains(output, "  "+wipBranch) || strings.Contains(output, "* "+wipBranch)
//}
//
//func hasMobbingBranchOrigin() bool {
//	output := silentgit("branch", "--remotes")
//	return strings.Contains(output, "  "+remoteName+"/"+wipBranch)
//}
//
//func getGitUserName() string {
//	return strings.TrimSpace(silentgit("config", "--get", "user.name"))
//}
//
//func isLastChangeSecondsAgo() bool {
//	changes := silentgit("--no-pager", "log", baseBranch+".."+wipBranch, "--pretty=format:%cr", "--abbrev-commit")
//	lines := strings.Split(strings.Replace(changes, "\r\n", "\n", -1), "\n")
//	numberOfLines := len(lines)
//	if numberOfLines < 1 {
//		return true
//	}
//
//	return strings.Contains(lines[0], "seconds ago") || strings.Contains(lines[0], "second ago")
//}
//
//func displayLog() {
//
//}
//
//func showNext() {
//	changes := strings.TrimSpace(silentgit("--no-pager", "log", baseBranch+".."+wipBranch, "--pretty=format:%an", "--abbrev-commit"))
//	lines := strings.Split(strings.Replace(changes, "\r\n", "\n", -1), "\n")
//	numberOfLines := len(lines)
//	if debug {
//		say("there have been " + strconv.Itoa(numberOfLines) + " changes")
//	}
//	gitUserName := getGitUserName()
//	if debug {
//		say("current git user.name is '" + gitUserName + "'")
//	}
//	if numberOfLines < 1 {
//		return
//	}
//	var history = ""
//	for i := 0; i < len(lines); i++ {
//		if lines[i] == gitUserName && i > 0 {
//			sayInfo("Committers after your last commit: " + history)
//			sayInfo("***" + lines[i-1] + "*** is (probably) next.")
//			return
//		}
//		if history != "" {
//			history = ", " + history
//		}
//		history = lines[i] + history
//	}
//}

func config() {
	say("config")
	for envKey, envValue := range envVariables {
		say(fmt.Sprintf("\t[%s] = %s", envKey, envValue))
	}
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

//
//func silentgit(args ...string) string {
//	command := exec.Command("git", args...)
//	if debug {
//		fmt.Println(command.Args)
//	}
//	outputBinary, err := command.CombinedOutput()
//	output := string(outputBinary)
//	if debug {
//		fmt.Println(output)
//	}
//	if err != nil {
//		fmt.Println(output)
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	return output
//}
//
//func hasSay() bool {
//	command := exec.Command("which", "say")
//	if debug {
//		fmt.Println(command.Args)
//	}
//	outputBinary, err := command.CombinedOutput()
//	output := string(outputBinary)
//	if debug {
//		fmt.Println(output)
//	}
//	return err == nil
//}
//
//func git(args ...string) string {
//	command := exec.Command("git", args...)
//	if debug {
//		fmt.Println(command.Args)
//	}
//	outputBinary, err := command.CombinedOutput()
//	output := string(outputBinary)
//	if debug {
//		fmt.Println(output)
//	}
//	if err != nil {
//		sayError(command.Args)
//		sayError(err)
//		os.Exit(1)
//	}
//	sayOkay(command.Args)
//
//	return output
//}

func say(s string) {
	fmt.Print(s)
	fmt.Print("\n")
}

func sayError(s string) {
	fmt.Print(" ⚡ ")
	say(s)
}

func sayOkay(s string) {
	fmt.Print(" ✓ ")
	say(s)
}

func sayNote(s string) {
	fmt.Print(" ❗ ")
	say(s)
}

func sayTodo(s string) {
	fmt.Print(" ☐ ")
	say(s)
}

func sayInfo(s string) {
	fmt.Print(" > ")
	say(s)
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
