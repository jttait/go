package main

import (
	"fmt"
	"time"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	counter := func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
		}
	}
	go counter()

	squarer := func() {
		for {
			x := <- naturals
			squares <- x * x
		}
	}
	go squarer()

	for {
		fmt.Println(<- squares)
	}
}
