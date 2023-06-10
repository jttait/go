package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type xml.Name
	Atrr []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)

	tok, err := dec.Token()
	if err == io.EOF {
		fmt.Fprintf(os.Stderr, "empty XML document.")
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
	tokse := tok.(xml.StartElement)
	e := Element{tokse.Name, []xml.Attr{}, []Node{}}
	root := &e
	stack := []*Element{&e}

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
			e := Element{tok.Name, []xml.Attr{}, []Node{}}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, e)
			}
			stack = append(stack, &e)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	fmt.Println()
	fmt.Println(root)
	fmt.Println()
	fmt.Println(root.Children)
	fmt.Println()
	fmt.Println(root.Children[0])
}

func printTree(n *Element, depth int) {
	if n == nil {
		return 
	}
	fmt.Println(n)
	//fmt.Printf("Element:%s\n", n.Type)
	for _, c := range n.Children {
		/*if cd, ok := c.(CharData); ok {
			fmt.Printf("CharData:%s", cd)
		} else {*/
			ce, _ := c.(Element)
			printTree(&ce, depth+1)
		//}
	}
}
