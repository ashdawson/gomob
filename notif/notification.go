package notif

import (
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"strings"
)

var appName = "GoMob"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Alert(message string) {
	err := beeep.Alert(appName, message, "assets/warning.png")
	check(err)
}

func Notify(message string) {
	err := beeep.Notify(appName, message, "assets/warning.png")
	check(err)
}

func Question(message string) bool {
	ok, err := dlgs.Question("Question", message, true)
	check(err)
	return ok
}

func List(message string, list []string) (string, bool) {
	selection, ok, err := dlgs.List("List", message, list)
	check(err)
	return selection, ok
}

func MultiList(message string, list []string) (string, bool) {
	selection, ok, err := dlgs.ListMulti("List", message, list)
	check(err)
	return strings.Join(selection, ","), ok
}

func Entry(message string, defaultText string) string {
	response, _, err := dlgs.Entry("Entry", message, defaultText)
	check(err)
	return response
}

func CanUse() bool {
	_, _, err := dlgs.Entry("Entry", "", "")
	if err != nil {
		return false
	}
	return true
}
