package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func git(args ...string) string {
	command := exec.Command("git", args...)
	if debug {
		fmt.Println(command.Args)
	}
	outputBinary, err := command.CombinedOutput()
	output := string(outputBinary)
	if debug {
		fmt.Println(output)
	}
	if err != nil {
		sayError(command.Args)
		sayError(err)
		os.Exit(1)
	}
	return output
}


func getGitUserName() string {
	return strings.TrimSpace(git("config", "--get", "user.name"))
}

func getBranchDetails() (string, string) {
	branchDetails := strings.Split(git("rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}"), "/")
	return strings.Trim(branchDetails[0], "\n"), strings.Trim(branchDetails[1], "\n")
}

func getBranch() string {
	return settings.RemoteName + "/" + settings.BranchName
}

func getLastCommitMessage() string {
	message := git("log", "-1", "--pretty=%B")
	message = strings.Replace(message, settings.CommitMessage, "", -1)
	message = strings.ReplaceAll(message, "\n", "")
	return message
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
