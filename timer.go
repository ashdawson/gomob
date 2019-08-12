package main

import (
	"fmt"
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
			next()
			join()
		} else {
			var selected, _ = notif.Reminder("Remind me again in:", []string{"5", "10", "15"})
			reminderTime, err := strconv.Atoi(selected)
			check(err)
			startTimer(reminderTime)
		}
	}()
}

func checkStatus() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		if getGitUserName() == showNext() && !isMobbing() && isLastChangeSecondsAgo() {
			fmt.Println("It is your turn to start")
		}
		checkStatus()
	}()
}
