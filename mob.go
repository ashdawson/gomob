package main

import (
	"fmt"
)

func start() {
	setup()
	runCommands()
}

func runCommands() {
	go join()
	for argKey := range argsMap {
		switch argKey {
		case "config":
			config()
			break
		case "start":
			startSession()
			status()
			startTimer(settings.TimeLimit)
			break
		case "next":
			next()
			break
		case "done":
			done()
			break
		case "help":
			help()
			break
		default:
			fmt.Println("OOPS")
		}
	}
	wg.Wait()
}
