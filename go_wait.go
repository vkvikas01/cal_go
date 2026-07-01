package main

import (
	"fmt"
	"sync"
)

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup

	 arr:= []string{
		"Hello",
		"Welcome",
		"Learning Go",
		"Goroutines are awesome",
		"Goodbye",
	}

	wg.Add(len(arr))

	for _, msg:= range arr{
		printMessage(msg, &wg)
	}

	wg.Wait()
	fmt.Println("All messages printed.")


}