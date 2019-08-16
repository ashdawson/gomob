package main

import (
	"github.com/ashdawson/gomob/notif"
	"strconv"
	"time"
)

var isTimerOnly = false
var sessionStartTime time.Time

func startTimer(reminderTime int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		settings.updateBranch()
		time.Sleep(time.Duration(reminderTime) * time.Minute)
		notif.Notify("Your mob time has ended")
		var swap = notif.Question("Your mob time has ended. Are you ready to swap?")
		if swap {
			next()
		} else {
			var selected, _ = notif.List("Remind me again in:", []string{"5", "10", "15"})
			reminderTime, err := strconv.Atoi(selected)
			check(err)
			startTimer(reminderTime)
		}
	}()
}

func join() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		settings.updateBranch()
		time.Sleep(1 * time.Second)
		if isLastChangeSecondsAgo() {
			git("pull")

			if getGitUserName() == getNextDriver() {
				notif.Notify("It is your turn to start")
				startSession()
				return
			}
		}
		join()
	}()
}
