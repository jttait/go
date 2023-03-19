package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	cancels := make(chan struct{})
	responses := make(chan http.Response, 1)
	for _, url := range os.Args[1:] {
		request, err := http.NewRequest("GET", url, nil)
		request.Cancel = cancels
		if err != nil {
			fmt.Println(err)
		}
		go fetch(request, responses)
	}
	resp := <- responses
	close(cancels)
	fmt.Printf("%s: %s\n", resp.Request.URL, resp.Status)
}

func fetch(req *http.Request, responses chan http.Response) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	responses <- *resp
}
