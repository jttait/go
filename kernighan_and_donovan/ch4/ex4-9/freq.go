package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Error: Must provide filename as argument.\n")
		os.Exit(1)
	}
	
	arg := os.Args[1]
	f, err := os.Open(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "freq: %v\n", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Print("word\toccurrences\n")
	for word, n := range counts {
		fmt.Printf("%s\t%d\n", word, n)
	}
}
