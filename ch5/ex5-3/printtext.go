package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}
	if n.Type == html.TextNode {
		text := n.Data
		text = strings.TrimLeft(text, " \n")
		text = strings.TrimRight(text, " \n")
		if len(text) > 0 {
			texts = append(texts, text)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit(texts, c)
	}
	return texts
}
