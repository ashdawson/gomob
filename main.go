package main

import (
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	say("GoMob - An automated mobbing tool")
	setup()
	runCommands()
}
