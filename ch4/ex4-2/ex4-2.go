package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
)

func main() {
	sha := flag.String("sha", "SHA256", "SHA256, SHA384, or SHA512")
	flag.Parse()

	args := flag.Args()
	s := args[0]

	if *sha == "SHA384" {
		fmt.Printf("SHA384 hash of '%s': %x\n", s, sha512.Sum384([]byte(s)))
	} else if *sha == "SHA512" {
		fmt.Printf("SHA512 hash of '%s': %x\n", s, sha512.Sum512([]byte(s)))
	} else if *sha == "SHA256" {
		fmt.Printf("SHA256 hash of '%s': %x\n", s, sha256.Sum256([]byte(s)))
	} else {
		fmt.Println("SHA not provided or not recognised. Defaulting to SHA256.")
		fmt.Printf("SHA256 hash of '%s': %x\n", s, sha256.Sum256([]byte(s)))
	}
}
