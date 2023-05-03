package main

import (
	"fmt"
	"sync"
)

func saySomething(i int, str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("saySomething(%d, %s)\n", i, str)
}

func main() {
	var wg sync.WaitGroup

	units := []string{
		"kg",
		"lb",
		"oz",
		"mg",
		"g",
		"t",
		"ml",
		"m",
		"cm",
	}

	wg.Add(len(units))

	for i, v := range units {
		go saySomething(i, v, &wg)
	}

	wg.Wait()
	fmt.Println("Done!")
}
