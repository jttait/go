package main

import "fmt"

func reverse(p *[4]int) {
	for i, j := 0, 3; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}	
}

func main() {
	integers := []int{1,2,3,4}
	pointer := (*[4]int)(integers)
	reverse(pointer)
	fmt.Println(integers)
}
