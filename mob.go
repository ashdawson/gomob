package main

import (
	"fmt"
	"strings"
)

var startCommand string
var mobRotation []string

type Committer struct {
	Name string
	Email string
	Count int
}

func getNextAuthor() string {
	// Used to determine new members
	committerStorage := make(map[string]Committer)
	committers := getCommitters()
	for i := 0; i < len(committers); i++ {
		details := strings.Split(committers[i], "|")
		if len(details) == 1 {
			break
		}
		email := details[0]
		name := details[1]
		committer, exists := committerStorage[email];

		if exists {
			committer.Count++
			committerStorage[email] = committer
		} else {
			committerStorage[email] = Committer{
				Name: name,
				Email: email,
				Count: 1,
			}
		}
	}

	committerStorageLen := len(committerStorage)
	currentCount := 1
	for _, v := range committerStorage {
		if currentCount == committerStorageLen {
			return v.Name
		}
		currentCount++
	}

	return ""
}

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
	output := git("--no-pager", "log", committers, "--pretty=format:%cn", "--all")
	membersSlice := make(map[string]bool)

	if output != "" {
		lines := strings.Split(strings.TrimSpace(output),"\n")
		for _, member := range lines {
			if _, ok := membersSlice[member]; !ok {
				membersSlice[member] = true
			}
		}
	}

	if len(mobRotation) > 1 {
		return mobRotation[len(mobRotation)-1]
	}

	return ""
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