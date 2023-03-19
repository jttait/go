package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	"os"
	//"fmt"
)

func main() {
	portFlag := flag.String("port", "8000", "Port that server runs on")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:" + *portFlag)
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
	tz := os.Getenv("TZ")
	location, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}
	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
