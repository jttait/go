package main

import (
	"github.com/jttait/gopl.io/ch12/ex12-2/display"
)

type cycle struct {
	value int
	tail *cycle
}

func main() {
	var c cycle
	c = cycle{42, &c}
	display.Display("c", c)
}
