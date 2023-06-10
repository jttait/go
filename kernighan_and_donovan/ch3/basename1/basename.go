package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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
