package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		fmt.Printf("%d ", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for char := 'a'; char <= 'z'; char++ {
		time.Sleep(time.Millisecond * 700)
		fmt.Printf("%c ", char)
	}
}

func main() {
	var wg sync.WaitGroup

	// Start goroutines
	wg.Add(1)
	go printLetters(&wg)

	// Wait for goroutines to finish
	wg.Wait()
}
