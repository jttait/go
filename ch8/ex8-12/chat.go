package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	channel chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.channel <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			msg := "Current clients are "
			for cli := range clients {
				msg += cli.name + ", "
			}
			for cli := range clients {
				cli.channel <- msg[:len(msg)-2]
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}

	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	
	who := conn.RemoteAddr().String()
	client := client{ ch, who }
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
