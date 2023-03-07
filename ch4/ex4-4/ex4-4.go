package main

import "fmt"

func rotate(s []int, n int) []int {
	result := make([]int, len(s), cap(s))
	left := s[:n]
	right := s[n:]
	for i := 0; i < len(s); i++ {
		if i < n {
			result[i] = right[i]
		} else {
			result[i] = left[i-n]
		}
	}
	return result
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 3)
	fmt.Println(s)
}
