package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string {
	if len(s) < 3 {
		return s
	}

	var buf bytes.Buffer

	n := len(s) % 3
	
	if n != 0 {
		buf.WriteString(s[:n])
		if len(s) - n >= 3 {
			buf.WriteByte(',')
		}
	}

	for n < len(s) {
		if len(s) - n >= 3 {
			buf.WriteString(s[n:n+3])
		} else {
			buf.WriteString(s[n:])
		}
		n += 3
		if len(s) - n >= 3 {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}

func main() {
	s := "12345"
	fmt.Println(s + " => " + comma(s))
	s = "123456"
	fmt.Println(s + " => " + comma(s))
	s = "12345678"
	fmt.Println(s + " => " + comma(s))
}
