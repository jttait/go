package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages("https://golang.org")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(words)
	fmt.Println(images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s\n", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	words = 0
	images = 0
	visit(n, &words, &images)
	return
}

func visit(n *html.Node, words *int, images *int) {
	if n.Type == html.TextNode {
		text := n.Data
		text = strings.TrimLeft(text, " \n")
		text = strings.TrimRight(text, " \n")
		texts := strings.Split(text, " ")
		*words += len(texts)
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*images++
	} else if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, images)
	}
}
