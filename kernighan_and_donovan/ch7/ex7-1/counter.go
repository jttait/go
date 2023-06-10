package main

import (
	"fmt"
	"bufio"
	"bytes"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	bytesWritten := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
		bytesWritten += len(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return bytesWritten, err
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	bytesWritten := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*c++
		bytesWritten += len(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return bytesWritten, err
	}
	return len(p), nil
}

func main() {
	var w WordCounter
	var l LineCounter
	w.Write([]byte("hello world"))
	fmt.Println(w)
	w = 0
	w.Write([] byte("this is four words"))
	fmt.Println(w)
	l.Write([]byte("hello world"))
	fmt.Println(l)
	l = 0
	l.Write([]byte("hello\nworld"))
	fmt.Println(l)
}
