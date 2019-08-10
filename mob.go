package main

import (
	"fmt"
	"strings"
)

func start() bool {

	output := silentgit("status", "--short")
	isMobbing := len(strings.TrimSpace(output)) == 0
	return isMobbing
}

func runCommands() {
	for argKey, _ := range argsMap {
		switch argKey {
		case "config":
			config()
			break
		case "start":
			//startSession()
			//status()
			break
		//case "join":
		//	start()
		//	join()
		//	status()
		//	break
		//case "next":
		//	next()
		//	break
		//case "done":
		//	done()
		//	break
		//case "help":
		//	help()
		//	break
		default:
			fmt.Println("OOPS")
		}
	}
}