package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func start() {
	runCommands()
}

func runCommands() {
	for argKey := range argsMap {
		switch argKey {
		case "config":
			config()
			break
		case "start":
			startSession()
			status()
			break
		case "join":
			startSession()
			join()
			status()
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
	startTimer(settings.TimeLimit)
	wg.Wait()
}
