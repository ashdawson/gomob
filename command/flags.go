package command

import (
	"flag"
	"fmt"
	"github.com/ashdawson/gomob/clock"
	"log"
)

var (
	StartDate clock.Date
)

// Read reads the command line
func Read() {
	startString := flag.String("start", clock.CurrentDate(), "a date string DD-MM-YYYY")
	flag.Parse()

	log.Printf("Using start date %v\n", *startString)
	StartDate = clock.CreateDate(startString)
	if len(flag.Args()) > 0 {
		fmt.Println("No command line options were found for: ", flag.Args())
	}
}
