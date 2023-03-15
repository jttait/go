package main

import (
	"io"
	"os"
	"fmt"
)

type WriterWithCount struct {
	writer io.Writer
	count *int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	return WriterWithCount{w, &count}, &count
}

func (cw WriterWithCount) Write(p []byte) (i int, err error) {
	cw.writer.Write(p)
	*cw.count += int64(len(p))
	return len(p), nil
}

func main() {
	cw, i := CountingWriter(os.Stdout)
	cw.Write([]byte("hello world\n"))
	fmt.Println(*i)
}
