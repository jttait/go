package main

import (
	"fmt"
	"log"
	"os"
	"net/url"
	"io"
	"net/http"
	"strings"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) [] string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(rawURL string) []string {
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	if u.Hostname() == "go.dev" {
		resp, err := http.Get(rawURL)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			fmt.Println(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			fmt.Println(err)
		}

		pathComponents := strings.Split(u.Path, "/")
		path := u.Hostname() + "/"
		for i := 0; i < len(pathComponents)-1; i++ {
			fmt.Println(path)
			path += pathComponents[i] + "/"
			os.Mkdir(path, 0750)
		}
		
		err = os.WriteFile(path + pathComponents[len(pathComponents)-1] + ".html", body, 0666)
		if err != nil {
			fmt.Println(err)
		}
	}

	list, err := links.Extract(rawURL)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
