package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("recovered in cleanup:", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("oh dear, we hit a 2!")
		}
	}
}

func main() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")
	wg.Wait()

	wg.Add(1)
	go say("hi")
	wg.Wait()
}
