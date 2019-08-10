package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"time"
)

func startTimer() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Minute)
		var swap = dlgSwap()
		if swap {
			fmt.Print("Pushing to git")
		} else {
			var time, _ = dlgRemindMe()
			fmt.Println(time)
		}
	}()
}

func dlgSwap() bool {
	ok, err := dlgs.Question("Question", "Your mob time has ended. Are you ready to swap?", true)
	if err != nil {
		panic(err)
	}
	return ok
}

func dlgRemindMe() (string, bool) {
	selection, ok, err := dlgs.List("List", "Remind me again in:", []string{"5", "10", "15"})
	if err != nil {
		panic(err)
	}
	return selection, ok
}
