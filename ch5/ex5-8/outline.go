package main

import (
	"os"
	"fmt"

	"golang.org/x/net/html"
)

var depth int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	element := ElementById(doc, "quote_slide2")
	if element != nil {
		fmt.Printf("%v\n", element)
	}
}

func ElementById(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	var found bool
	if pre != nil {
		found = pre(n, id)
	}
	if found == true {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := forEachNode(c, id, pre, post)
		if result != nil {
			return result
		}
	}

	if post != nil {
		found = post(n, id)
	}
	if found == true {
		return n
	}
	return nil
}

func startElement(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}
	return false
}
