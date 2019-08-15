package main

import (
	"strings"
)

var startCommand string

type Committer struct {
	Name string
	Email string
	Count int
}

func showNext() string {
	committerStorage := make(map[string]Committer)
	committers := getCommitters()

	for i := 0; i < len(committers); i++ {
		details := strings.Split(committers[i], "|")
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