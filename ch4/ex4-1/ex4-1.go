package main

import (
	"fmt"
	"crypto/sha256"
)

func countDifferentBits (hash1 [32]byte, hash2 [32]byte) int {
	result := 0
	for i := 0; i < 32; i++ {
		x := hash1[i]
		y := hash2[i]
		for j := 0; j < 8; j++ {
			if x != y {
				result++
			}
			x >>= 1
			y >>= 1
		}
	}
	return result
}

func main() {
	fmt.Printf("%d\n", countDifferentBits(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("x"))))
	fmt.Printf("%d\n", countDifferentBits(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X"))))
}
