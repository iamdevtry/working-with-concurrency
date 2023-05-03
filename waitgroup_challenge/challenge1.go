package main

import (
	"fmt"
	"sync"
)

var msg string

func updateString(s string) {
	msg = s
}

func printMessage() {
	println(msg)
}

func main() {
	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	var wg sync.WaitGroup
	msg = "Hello, world!"

	wg.Add(1)

	go func() {
		updateString("Hello, universe!")
		printMessage()

		defer wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		updateString("Hello, Cosmo!")
		printMessage()

		defer wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		updateString("Hello, world!")
		printMessage()
		defer wg.Done()
	}()
	wg.Wait()

	fmt.Println("done")
}
