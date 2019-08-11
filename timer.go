package main

import (
	"strconv"
	"time"
)

func startTimer(reminderTime int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(reminderTime) * time.Minute)
		Notify("Your mob time has ended")
		var swap = Question("Your mob time has ended. Are you ready to swap?")
		if swap {
			changes := getChangesOfLastCommit()
			say(changes)
			//next()
		} else {
			var selected, _ = Reminder("Remind me again in:", []string{"5", "10", "15"})
			reminderTime, err := strconv.Atoi(selected)
			check(err)
			startTimer(reminderTime)
		}
	}()
}
