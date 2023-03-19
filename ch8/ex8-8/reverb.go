package main

import (
	"net"
	"fmt"
	"log"
	"bufio"
	"time"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	countdown := 10
	timeout := make(chan bool)
	inputs := make(chan string)
	scanner := bufio.NewScanner(c)

	go func() {
		for countdown > 0 {
			countdown--
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Connection timed out after 10 seconds.")
		timeout <- true
	}()

	go func() {
		for scanner.Scan() {
			inputs <- scanner.Text()
		}
	}()

	for {
		select {
		case x := <-inputs:
			go echo(c, x, 1 * time.Second)
			countdown = 10
		case <- timeout:
			c.Close()
			return
		default:
		} }
}
