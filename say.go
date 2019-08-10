package main

import "fmt"

func say(s string) {
	fmt.Print(s)
	fmt.Print("\n")
}

func sayError(s string) {
	fmt.Print(" ⚡ ")
	say(s)
}

func sayOkay(s string) {
	fmt.Print(" ✓ ")
	say(s)
}

func sayNote(s string) {
	fmt.Print(" ❗ ")
	say(s)
}

func sayTodo(s string) {
	fmt.Print(" ☐ ")
	say(s)
}

func sayInfo(s string) {
	fmt.Print(" > ")
	say(s)
}