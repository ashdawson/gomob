package main

import (
	"github.com/ashdawson/gomob/notif"
	"time"
	"fmt"
)

func startTimer(reminderTime int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		config.updateBranch()
		time.Sleep(time.Duration(reminderTime) * time.Minute)
		swap := AskConfirmation("Your mob time has ended. Are you ready to swap?")
		if swap {
			next()
		} else {
			sayInfo(fmt.Sprintf("Will remind you again in (%d minutes)", 5))
			time.Sleep(time.Duration(5) * time.Minute)
		}
	}()
}

func joinTimer() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		config.updateBranch()
		time.Sleep(1 * time.Second)
		if isLastChangeSecondsAgo() {
			if getGitUserName() == getNextDriver() {
				git("pull")

				notif.Notify("It is your turn to start")
				startSession()
				return
			}
		}
		joinTimer()
	}()
}
