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
	//solve race condition problem
	//solve race condition problem
	wg.Add(3)
	go updateMessage("Hello!")
	go updateMessage("Hola!")
	go updateMessage("Ni hao!")
	wg.Wait()

	println(msg)

}

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "Xin chao!"
// 	var mutex sync.Mutex
// 	//solve race condition problem
// 	//solve race condition problem
// 	wg.Add(3)
// 	go updateMessage("Hello!", &mutex)
// 	go updateMessage("Hola!", &mutex)
// 	go updateMessage("Ni hao!", &mutex)
// 	wg.Wait()

// 	println(msg)

// }
