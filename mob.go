package main

import "strings"

var startCommand string

type Committer struct {
	Name string
	Email string
	Count int
}

func showNext() string {
	committersStorage := make(map[string]Committer)
	committers := getCommitters()

	for i := 0; i < len(committers); i++ {
		details := strings.Split(committers[i], "|")
		email := details[0]
		name := details[1]
		committer, exists := committersStorage[email];

		if exists {
			committer.Count++
			committersStorage[email] = committer
		} else {
			committersStorage[email] = Committer{
				Name: name,
				Email: email,
				Count: 1,
			}
		}
	}

	return ""
}