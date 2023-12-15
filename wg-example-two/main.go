package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {

	msg = "Hello, world!"

	data := []string{
		"universe!",
		"cosmos!",
		"world!",
	}

	for _, x := range data {
		wg.Add(1)
		go updateMessage(fmt.Sprintf("Hello, %s", x))
		wg.Wait()
		printMessage()
	}
}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}
