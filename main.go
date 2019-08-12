package main

func main() {
	go forever()
	start()
}

func forever() {
	wg.Add(1)
}