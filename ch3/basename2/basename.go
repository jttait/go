package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	s := "a/b/c.go"
	fmt.Println(s + " => " + basename(s))
	s = "a.d.go"
	fmt.Println(s + " => " + basename(s))
	s = "abc"
	fmt.Println(s + " => " + basename(s))
}
