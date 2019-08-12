package command

import (
	"flag"
	"fmt"
)

// Read reads the command line
func Read() {
	startString := flag.String("ide", "", "the IDE you are currently using")
	flag.Parse()

	fmt.Println(startString)
	if len(flag.Args()) > 0 {
		fmt.Println("No command line options were found for: ", flag.Args())
	}
}
