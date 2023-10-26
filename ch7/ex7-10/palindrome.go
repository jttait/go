package main

import (
	"sort"
	"fmt"
)

func main() {
	ss := sort.StringSlice([]string{"k", "a", "y", "a", "k"})
	fmt.Printf("Is %s a palindrome? %t\n", ss, IsPalindrome(ss))
	ss = sort.StringSlice([]string{"k", "i", "p", "p", "e", "r"})
	fmt.Printf("Is %s a palindrome? %t\n", ss, IsPalindrome(ss))
}

func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for i < j {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
		i++; j--
	}
	return true
}
