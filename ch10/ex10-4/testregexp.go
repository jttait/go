package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "hello   world"
	re := regexp.MustCompile(`\s\s+`)
	s = re.ReplaceAllString(s, " ")
	fmt.Println(s)
}
