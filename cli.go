package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AskInput(s string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.TrimSpace(response)
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

		if response == "n" || response == "no" {
			return false
		} else {
			return true
		}
	}
}