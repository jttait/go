package main

import (
	"fmt"
	"time"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	counter := func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
		}
		close(naturals)
	}
	go counter()

	squarer := func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}
	go squarer()

	for x := range squares{
		fmt.Println(x)
	}
}
