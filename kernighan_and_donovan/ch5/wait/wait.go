package main

import (
	"time"
	"net/http"
	"fmt"
	"log"
)

func main() {
	err := waitForServer("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	err = waitForServer("https://www.abcdjklfaasdyfuasdfjkljewanlfds.com")
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	if err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
