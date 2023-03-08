package main

import (
	"os"
	"fmt"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	images := ElementsByTagName(doc, "img")
	for _, image := range images {
		fmt.Printf("%v\n", image)
	}

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, heading := range headings {
		fmt.Printf("%v\n", heading)
	}
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var result []*html.Node
	result = forEachNode(doc, result, names...)
	return result
}

func forEachNode(n *html.Node, result []*html.Node, names ...string) []*html.Node {
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				result = append(result, n)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = forEachNode(c, result, names...)
	}
	return result
}
