package main

import "time"

func main() {
	go forever()
	start()
}

func forever() {
	time.Sleep(time.Second)
}