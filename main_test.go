package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func test_saySomething(t *testing.T) {
	stOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go saySomething(1, "test", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stOut

	if !strings.Contains(output, "saySomething(1, test1)") {
		t.Errorf("saySomething() did not print the expected output: %s", output)
	}
}
