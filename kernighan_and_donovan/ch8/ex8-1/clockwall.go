package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"bufio"
	"fmt"
)

func main() {
	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, "=")
		go display(split[0], split[1])
	}
	for {
	}
}

func display(name string, url string) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "%s\t%s\n", name, scanner.Text())
	}
}

func mustCopy(dst io.Writer, src io.Reader, name string) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
