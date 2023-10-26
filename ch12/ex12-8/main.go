package main

import (
	"fmt"
	"io"
	//"log"
	"bytes"

	"github.com/jttait/gopl.io/ch12/ex12-8/sexpr"
)

type Movie struct {
	Title string
	Year int
}


func main() {
	strangelove := []byte("((Title \"Dr. Strangelove\") (Year 1964))")

	dec := sexpr.NewDecoder(bytes.NewReader(strangelove))
	for {
		var m Movie
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			//log.Fatal(err)

		}
		fmt.Printf("Title: %s; Year: %d\n", m.Title, m.Year)
	}
}
