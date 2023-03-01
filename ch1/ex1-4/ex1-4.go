package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1-4: %v\n", err)
		}
		if (duplicateLines(f)) {
			fmt.Println(arg)
		}
		f.Close()
	}
}

func duplicateLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] != 1 {
			return true
		}
	}
	return false
}
