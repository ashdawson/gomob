package main

import (
	"fmt"
	"strings"
)

var startCommand string
var mobRotation []string

func getPossibleTeam() []string {
	var possibleTeam []string
	output := git("shortlog", getBranch(), "-sne", "--since=7.days")
	if output != "" {
		lines := strings.Split(output,"\n")
		for i := 0; i < len(lines); i++ {
			member := strings.Split(lines[i],"\t")
			if !strings.Contains(member[1], getGitUserEmail()) {
				emailIndex := strings.Index(member[1], "<")
				possibleTeam = append(possibleTeam, member[1][0:emailIndex])
			}
		}
	}

	return possibleTeam
}

func getNextDriver() string {
	committers := fmt.Sprintf("--author=%s", getMobMembers())
	output := git("--no-pager", "log", committers, "--pretty=format:%cn", "--all", "--since=1.days")
	membersSlice := make(map[string]bool)
	nextIndex := 0

	if output != "" {
		lines := strings.Split(strings.TrimSpace(output),"\n")
		for _, member := range lines {
			if _, ok := membersSlice[member]; !ok {
				membersSlice[member] = true
				nextIndex = IndexOf(mobRotation, member)
			}
		}
	}

	return mobRotation[nextIndex]
}

func getMobMembers() string {
	members := strings.Split(settings.Mob, ",")
	mobRotation = append(members, getGitUserName())

	var output string

	for _, member := range mobRotation {
		output += fmt.Sprintf("\\(%s\\)\\|", member)
	}

	return strings.TrimSuffix(output, "\\|")
}

func IndexOf(haystack []string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle && len(haystack) != i {
			return i
		}
	}

	return 0
}