package main

import (
	"log"
	"net"
	"bufio"
	"fmt"
	"strings"
	//"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Fprintln(c, "220")
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Fprintln(c, "error!")
			continue
		}
		t := scanner.Text()
		fmt.Println(t)
		if strings.HasPrefix(t, "USER") {
			fmt.Fprintln(c, "230")
		} else if t == "close" {
			fmt.Fprintln(c, "Closing connection.")
			c.Close()
		} else if t == "ls" {
			fmt.Fprintln(c, "hello, ls")
		}
		fmt.Fprintln(c, t)
	}
}
