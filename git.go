package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var lastBranch string
var changeList []string

func hasGit() {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		checkSay(err, "git has not been added to this directory")
	}
}

func git(args ...string) string {
	command := exec.Command("git", args...)
	if debug() {
		fmt.Println(command.Args)
	}
	outputBinary, err := command.CombinedOutput()
	output := string(outputBinary)
	if debug() {
		fmt.Println(output)
	}
	if err != nil {
		sayError(command.Args)
		sayError(err)
		exit()
	}
	return strings.TrimSpace(output)
}

func getGitUserName() string {
	return strings.TrimSpace(git("config", "--get", "user.name"))
}

func getGitUserEmail() string {
	return strings.TrimSpace(git("config", "--get", "user.email"))
}

func getBranchDetails() (string, string) {
	branchDetails := strings.Split(git("rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}"), "/")
	return strings.Trim(branchDetails[0], "\n"), strings.Trim(branchDetails[1], "\n")
}

func getBranch() string {
	return config.RemoteName + "/" + config.BranchName
}

func getLastCommitMessage() string {
	return git("log", "-1", "--pretty=%B")
}

func getLastCommittedFiles(message string) []string {
	var files []string
	if strings.Contains(message, config.CommitMessage) {
		message = strings.Replace(message, config.CommitMessage, "", -1)
		message = strings.Replace(message, "\n", "", -1)
		files = strings.Split(message, " ")
		for i := range files {
			files[i] = getCurrentDir() + `\` + files[i]
		}
	}

	return files
}

func getModifiedFiles() string {
	fileNames := strings.Split(git("diff", "--name-only"), "\n")
	fileString := ""
	for _, file := range fileNames {
		if len(file) > 0 {
			fileString = fileString + file + " "
		}
	}
	fileString = strings.TrimSuffix(fileString, " ")
	return fileString
}

func getCommitters() []string {
	commits := strings.TrimSpace(git("--no-pager", "log", "-n", "10", getBranch(), "--since=1.days", "--pretty=format:%ae|%an"))
	return strings.Split(strings.Replace(commits, "\r\n", "\n", -1), "\n")
}

func hasCommits() bool {
	output := git("status", "--short")
	isMobbing := len(strings.TrimSpace(output)) > 0
	return isMobbing
}

func isMobbing() bool {
	output := git("branch")
	return strings.Contains(output, "* "+config.BranchName)
}

func isLastChangeSecondsAgo() bool {
	_ = git("fetch")
	recentlyUpdated := git("--no-pager", "log", "HEAD", "-1", "--pretty=format:%cr", "--abbrev-commit")
	return strings.Contains(recentlyUpdated, "seconds ago") || strings.Contains(recentlyUpdated, "second ago")
}

func (settings *Config) updateBranch() {
	remoteName, branchName := getBranchDetails()
	if settings.RemoteName != remoteName || settings.BranchName != branchName {
		settings.RemoteName = remoteName
		settings.BranchName = branchName

		sayInfo(fmt.Sprintf("Now tracking changes to: %s/%s", settings.RemoteName, settings.BranchName))
		settings.save()
	}
}

func getLastFileChanges(filenames []string) {
	for i := range filenames {
		for minute := 1; minute < config.TimeLimit; minute++ {
			blame := git("blame","--since=" + strconv.Itoa(minute) + ".seconds",filenames[i],"|", doesNotInclude("^\\^"))

			if len(blame) > 0 {
				fmt.Println(blame)
				break
			}
		}
	}
}

func doesNotInclude(content string) string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("findstr /V '%s'", content)
	}
	return fmt.Sprintf("grep -v '%s'", content)
}

func commitMessage() string {
	modifiedFiles := getModifiedFiles()
	if len(modifiedFiles) == 0 {
		modifiedFiles = "Empty commit to be removed during rebase"
	}
	return config.CommitMessage + modifiedFiles
}

func commit(input string) {
	message := input
	if len(input) <= 0 {
		message = commitMessage()
	}

	git("add", "--all")
	git("commit", "--message", message)
	//git("commit", "--allow-empty", "--message", message)
}
