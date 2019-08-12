package main

import (
	"github.com/ashdawson/gomob/notif"
	"strconv"
	"time"
)

func startTimer(reminderTime int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(reminderTime) * time.Minute)
		notif.Notify("Your mob time has ended")
		var swap = notif.Question("Your mob time has ended. Are you ready to swap?")
		if swap {
			changes := getChangesOfLastCommit()
			say(changes)
			next()
		} else {
			var selected, _ = notif.Reminder("Remind me again in:", []string{"5", "10", "15"})
			reminderTime, err := strconv.Atoi(selected)
			check(err)
			startTimer(reminderTime)
		}
	}()
}
