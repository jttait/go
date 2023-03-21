package main

import (
	"fmt"
	//"archive/tar"
	"archive/zip"
	"io"
	"os"
	"log"
)

func main() {
	reader, _ := os.Open(os.Args[1])
	tr, _ := zip.NewReader(reader)
	//tr = tar.NewReader(reader)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}

func unarchive(format string, filename string) {

}
