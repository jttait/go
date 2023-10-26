package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squash(u []byte) []byte {
	i := 0
	length := len(u)
	for i < length {
		r, _ := utf8.DecodeRune(u[i:])
		if unicode.IsSpace(r) {
			u[i] = ' '
			j := i + utf8.RuneLen(r)
			rj, _ := utf8.DecodeRune(u[j:])
			for j < length && unicode.IsSpace(rj) {
				j += utf8.RuneLen(rj)
				rj, _ = utf8.DecodeRune(u[j:])
			}
			k := i + 1
			for  j < length {
				u[k] = u[j]
				j++
				k++
			}
			length -= j - k
			i++	
		} else {
			i += utf8.RuneLen(r)
		}
	}
	return u[:length]
}

func main() {
	s := []byte("hello  \u4e16")
	fmt.Println(s)
	s = squash(s)
	fmt.Println(s)
	s = []byte("hello\t\t\ttabs")
	fmt.Println(s)
	s = squash(s)
	fmt.Println(s)

}
