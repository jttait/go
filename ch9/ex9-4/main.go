package main

import "fmt"

func main() {
	channel := make(chan int)
	stage(channel, 0)
}

func stage(channel chan int, depth int) {
	fmt.Println(depth)
	channel2 := make(chan int)
	go func() {
		for {
			x := <- channel
			channel2 <- x + 1
		}
	}()
	stage(channel2, depth+1)
}
