package main

import "sync"

var wg sync.WaitGroup
var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Xin chao!"

	wg.Add(3)
	go updateMessage("Hello!")
	go updateMessage("Hola!")
	go updateMessage("Ni hao!")
	wg.Wait()

	println(msg)

}
