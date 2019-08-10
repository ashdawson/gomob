package main

import (
	"sync"
)

var wg = &sync.WaitGroup{}

func version() {

}

func main() {
	setup()
	wg.Wait()
}
