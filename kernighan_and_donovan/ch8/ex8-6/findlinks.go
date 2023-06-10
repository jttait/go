package main

import (
	"fmt"
	"log"
	"flag"

	"github.com/jttait/gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<- tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	var depthFlag = flag.Int("depth", 3, "URLs reachable by at most this number of links")
	flag.Parse()
	worklist := make(chan []string)
	var n int
	d := 0
	n++
	go func() { worklist <- flag.Args() }()
	seen := make(map[string]bool)
	for ; n > 0 && d <= *depthFlag; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
		d++
	}
}
