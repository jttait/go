package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	occurrences := make(map[int]int)

	for r := range s1 {
		occurrences[r]++
	}
	for r := range s2 {
		occurrences[r]--
	}

	for _, v := range occurrences {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "restful"
	s2 := "fluster"
	fmt.Printf("%s, %s => %t\n", s1, s2, isAnagram(s1, s2))
	s1 = "bird"
	s2 = "rabbit"
	fmt.Printf("%s, %s => %t\n", s1, s2, isAnagram(s1, s2))
}
