package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkSay(e error, message string) {
	if e != nil {
		sayInfo(message)
		exit()
	}
}

func exit() {
	sayInfo("exiting application")
	os.Exit(1)
}

func debug() bool {
	return settings.Debug
}
