package main

import "fmt"

func eliminateAdjacentDuplicates(s []int) []int {
	length := len(s)
	i := 0
	for i < length {
		if s[i] == s[i+1] {
			j := i+1
			for j < length && s[j] == s[i] {
				j++
			}
			k := i+1
			for j < length {
				s[k] = s[j]
				k++
				j++
			}
			length -= j - k
		}
		i++
	}
	return s[:length]
}

func main() {
	s := []int{5, 5, 6, 7, 8}
	s = eliminateAdjacentDuplicates(s)
	fmt.Println(s)
	s = []int{5, 6, 7, 8, 8}
	s = eliminateAdjacentDuplicates(s)
	fmt.Println(s)
	s = []int{5, 6, 7, 7, 7, 8}
	s = eliminateAdjacentDuplicates(s)
	fmt.Println(s)
}
