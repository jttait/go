package main

import (
	"os"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type StringReader string

func (r StringReader) Read(p []byte) (n int, err error) {
	rBytes := []byte(r)
	for i := 0; i < len(p); i++ {
		if i >= len(rBytes) {
			return i, io.EOF
		}
		p[i] = rBytes[i]
	}
	return len(p), nil
}

func NewReader(s string) *StringReader {
	r := StringReader(s)
	return &r
}

func main() {
	s := "<html><head><title>Hello, world</title></head><body><h1>hello, world</h1></body></html>"
	r := NewReader(s)
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
