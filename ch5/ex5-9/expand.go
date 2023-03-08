package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, my name is $name and I am $age years old."
	s = expand(s, populateField)
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {
	result := ""
	for _, word := range strings.Split(s, " ") {
		result += f(word) + " "
	}
	return result
}

func populateField(word string) string {
	if word == "$name" {
		return "Bob"
	} else if word == "$age" {
		return "99"
	}
	return word
}


