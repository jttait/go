package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan *int)
	channel2 := make(chan *int)

	count := 0
	start := time.Now()

	go func() {
		channel1 <- &count
	}()

	go func() {
		for {
			x := <- channel2
			*x += 1
			channel1 <- x
		}
	}()
	go func() {
		for {
			x := <- channel1
			*x += 1
			channel2 <- x
		}
	}()

	fmt.Scanln()

	fmt.Printf("%.0f messages per second.\n", float64(count) / time.Since(start).Seconds())
}
