package main

import "fmt"

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
