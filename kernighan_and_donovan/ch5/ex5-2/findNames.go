package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	names := make(map[string]int)
	visit(&names, doc)
	
	for key, value := range names {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func visit(names *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		(*names)[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(names, c)
	}
}
