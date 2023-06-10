package main

import (
	"strings"
	"io"
	"os"
	"log"
	"fmt"
)

type MyLimitedReader struct {
	R io.Reader
	N int64
}

func (lr MyLimitedReader) Read(p []byte) (n int, err error) {
	buf := make([]byte, lr.N)
	lr.R.Read(buf)
	for i := 0; i < int(lr.N); i++ {
		p[i] = buf[i]
	}
	return int(lr.N), io.EOF
}

func MyLimitReader(r io.Reader, n int64) io.Reader {
	var lr MyLimitedReader
	lr.R = r
	lr.N = n
	return lr
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := MyLimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
