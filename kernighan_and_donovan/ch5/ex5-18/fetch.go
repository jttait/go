package main

import (
	"net/http"
	"os"
	"io"
	"path"
	"fmt"
)

func main() {
	local, _, err := fetch(os.Args[1])
	fmt.Println(local)
	fmt.Println(err)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer closeFile(f, &err)()

	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return local, n, err
	}

	return local, n, err
}

func closeFile(file *os.File, err *error) func() {
	return func() {
		closeErr := file.Close()
		if err == nil {
			*err = closeErr
		}
	}
}
