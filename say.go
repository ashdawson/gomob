package main

import (
	"fmt"
	"github.com/ashdawson/gomob/notif"
)

func say(s interface{}) {
	fmt.Print(s)
	fmt.Print("\n")
}

func sayError(s interface{}) {
	fmt.Print(" ⚡ ")
	say(s)
}

func sayOkay(s interface{}) {
	fmt.Print(" ✓ ")
	say(s)
}

func sayNote(s interface{}) {
	fmt.Print(" ❗ ")
	say(s)
}

func sayTodo(s interface{}) {
	fmt.Print(" ☐ ")
	say(s)
}

func sayInfo(s interface{}) {
	fmt.Print(" > ")
	say(s)
}

func sayNotify(s interface{}) {
	fmt.Print(" > ")
	say(s)
	notif.Notify(fmt.Sprintf("%v", s))
}
