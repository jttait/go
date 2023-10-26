package main

import "fmt"

func main() {
	fmt.Println(f(1))
}

func f(x int) (s string) {
	defer func() {
		recover()
		s = "I was returned without a return statement."
	}()
	f(1/x-1)
	panic("panic!")
}
