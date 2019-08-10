package main

import "fmt"

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
