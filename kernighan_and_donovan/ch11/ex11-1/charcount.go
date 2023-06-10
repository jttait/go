package charcount

import (
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
	"strings"
)

func Count(input string) (map[rune]int, [utf8.UTFMax+1]int, int) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax+1]int
	invalid := 0


	in := strings.NewReader(input)

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
		counts[r]++
		utflen[n]++
	}

	return counts, utflen, invalid
}
