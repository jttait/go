package main

import (
	"fmt"
	"bytes"
	"strings"
)

func comma(s string) string {
	if len(s) < 3 {
		return s
	}

	var buf bytes.Buffer

	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	dot := strings.LastIndex(s, ".")
	decimal := ""
	if dot != -1 {
		decimal = s[dot:]
		s = s[:dot]
	}

	n := len(s) % 3
	
	if n != 0 {
		buf.WriteString(s[:n])
		if len(s) - n >= 3 {
			buf.WriteByte(',')
		}
	}

	for n < len(s) {
		if len(s) - n > 3 {
			buf.WriteString(s[n:n+3])
		} else {
			buf.WriteString(s[n:])
		}
		n += 3
		if len(s) - n >= 3 {
			buf.WriteByte(',')
		}
	}

	buf.WriteString(decimal)

	return buf.String()
}

func main() {
	s := "12345"
	fmt.Println(s + " => " + comma(s))
	s = "123456"
	fmt.Println(s + " => " + comma(s))
	s = "-123456"
	fmt.Println(s + " => " + comma(s))
	s = "+123456"
	fmt.Println(s + " => " + comma(s))
	s = "123456.00"
	fmt.Println(s + " => " + comma(s))
	s = "12345678.05"
	fmt.Println(s + " => " + comma(s))
}
