package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"os"

	"github.com/jttait/gopl.io/ch9/memo1"
	"github.com/jttait/gopl.io/ch9/memo2"
	"github.com/jttait/gopl.io/ch9/memo3"
	"github.com/jttait/gopl.io/ch9/memo4"
	"github.com/jttait/gopl.io/ch9/memo5"
	"github.com/jttait/gopl.io/ch9/memo6"
)

func main() {
	fmt.Println()
	fmt.Println("== memo1 ==")
	fmt.Println()
	useMemo1()
	fmt.Println()
	fmt.Println("== memo2 ==")
	fmt.Println()
	useMemo2()
	fmt.Println()
	fmt.Println("== memo3 ==")
	fmt.Println()
	useMemo3()
	fmt.Println()
	fmt.Println("== memo4 ==")
	fmt.Println()
	useMemo4()
	fmt.Println()
	fmt.Println("== memo5 ==")
	fmt.Println()
	useMemo5()
	fmt.Println()
	fmt.Println("== ex9-3 ==")
	fmt.Println()
	useMemo6()
}

func incomingUrls() []string {
	return []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://gopl.io",
	}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func httpGetBodyWithCancel(url string, done chan struct{}) (interface{}, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func useMemo1() {
	m := memo1.New(httpGetBody)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func useMemo2() {
	m := memo2.New(httpGetBody)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func useMemo3() {
	m := memo3.New(httpGetBody)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func useMemo4() {
	m := memo4.New(httpGetBody)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func useMemo5() {
	m := memo5.New(httpGetBody)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func useMemo6() {
	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	m := memo6.New(httpGetBodyWithCancel)
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}
