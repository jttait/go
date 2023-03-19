package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/jttait/gopl.io/ch5/links"
)

func crawl(req *http.Request) []string {
	fmt.Println(req.URL)
	list, err := links.Extract(req)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	cancels := make(chan struct{})
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	go func() { worklist <- os.Args[1:] }()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancels)
		os.Exit(0)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				req, _ := http.NewRequest("GET", link, nil)
				req.Cancel = cancels
				foundLinks := crawl(req)
				go func() { worklist<- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
