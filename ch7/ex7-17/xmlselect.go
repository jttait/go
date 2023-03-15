package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"flag"
)

func main() {
	var structureFlag = flag.String("structure", "", "Location in XML tree structure")
	var idFlag = flag.String("id", "", "ID of element")
	var classFlag = flag.String("class", "", "Class of element")
	flag.Parse()

	structure := strings.Split(*structureFlag, " ")

	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	id := ""
	class := ""
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			for _, attr := range tok.Attr {
				if attr.Name.Local == "id" {
					id = attr.Value
				} else {
					id = ""
				}
				if attr.Name.Local == "class" {
					class = attr.Value
				} else {
					class = ""
				}
			}	
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			structureMatch := *structureFlag == "" || containsAll(stack, structure)
			idMatch := *idFlag == "" || id == *idFlag
			classMatch := *classFlag == "" || class == *classFlag
			if structureMatch && idMatch && classMatch {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
