package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channels block
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// this code blocks until the value is taken off
			c <- i
		}
	}()

	go func() {
		for {
			// print the value
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
