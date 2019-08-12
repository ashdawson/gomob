package main

import (
	"fmt"
)

func start() {
	setup()
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
			break
		case "join":
			join()
			break
		case "next":
			next()
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
