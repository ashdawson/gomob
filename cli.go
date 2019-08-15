package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AskInput(s string, choices []string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", s)

		for choice := range choices {
			fmt.Println(choice)
		}

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))
		return response
	}
}

func AskConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}