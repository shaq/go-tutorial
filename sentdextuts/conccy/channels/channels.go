package main

import (
	"fmt"
	"sync"
)

func foo(c chan int, someValue int) {

	defer wg.Done()
	c <- someValue * 5
}

var wg sync.WaitGroup

func main() {
	// Create an int channel with a buffer of 10 to
	// stop blocking.
	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		// Each time we create a goroutine, we'll add it to
		// our WaitGroup for synchronisation.
		wg.Add(1)
		go foo(fooVal, i)
	}

	// Wait for all of the goroutines to load onto the channel
	// before close it!
	wg.Wait()
	close(fooVal)

	// Read from the channel and print to stdout.
	for val := range fooVal {
		fmt.Println(val)
	}
}
