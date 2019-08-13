package main

import (
	"fmt"
)

var startCommand string

func start() {
	setup()
	runCommands()
}

func runCommands() {
	sayInfo("running command: " + startCommand)
	switch startCommand {
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
	wg.Wait()
}
