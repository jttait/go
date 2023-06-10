package main

import (
	"github.com/jttait/gopl.io/ch12/ex12-1/display"
)

type mapkey struct {
	num int
	name string
}

func main() {
	display.Display("x", 1)
	display.Display("x", map[int]string{1: "a", 2: "b", 3:"c"})
	display.Display("x", mapkey{1, "a"})
	m := map[mapkey]string{
		mapkey{1, "a"}: "hello",
		mapkey{2, "b"}: "world",
	}
	display.Display("x", m)
	a1 := [3]int{1,2,3}
	display.Display("x", a1)
	a2 := [3]int{4,5,6}
	m2 := map[[3]int]string{
		a1: "hello",
		a2: "world",
	}
	display.Display("x", m2)
}
