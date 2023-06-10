package main

import (
	"fmt"
)

var prereqs = map[string][]string {
	"algorithms": { "data structures" },
	"calculus": { "linear algebra" },
	"compilers": { 
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": { "discrete math" },
	"databases": { "data structures" },
	"discrete math": { "intro to programming " },
	"formal languages": { "discrete math" },
	"networks": { "operating systems" },
	"operating systems": {
		"data structures",
		"computer organization",
	},
	"programming languages": {
		"data structures", 
		"computer organization",
	},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string

	seen := make(map[string]bool)
	var visitAll func(key string)

	visitAll = func(key string) {
		for _, v := range m[key] {
			if !seen[v] {
				seen[v] = true
				visitAll(v)
				order = append(order, v)
			}
		}
		if !seen[key] {
			seen[key] = true
			order = append(order, key)
		}
	}

	for k, _ := range m {
		visitAll(k)
	}
	return order
}
