package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax+1]int
	digits := make(map[rune]int)
	upper := make(map[rune]int)
	lower := make(map[rune]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			digits[r]++
		}
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			upper[r]++
		}
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			lower[r]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ndigit\tcount\n")
	for i, n := range digits {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}
	fmt.Print("\nupper\tcount\n")
	for i, n := range upper {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}
	fmt.Print("\nlower\tcount\n")
	for i, n := range lower {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
