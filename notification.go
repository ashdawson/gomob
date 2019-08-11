package main

import (
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
)

func Alert(message string) {
	err := beeep.Alert("GoMob", message, "assets/warning.png")
	check(err)
}

func Notify(message string) {
	err := beeep.Notify("GoMob", message, "assets/warning.png")
	check(err)
}

func Question(message string) bool {
	ok, err := dlgs.Question("Question", message, true)
	check(err)
	return ok
}

func Reminder(message string, list []string) (string, bool) {
	selection, ok, err := dlgs.List("List", message, list)
	check(err)
	return selection, ok
}

