package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Network struct {
	Connections []*Connection `json:"links"`
}

type Connection struct {
	From string `json:"source"`
	To string `json:"target"`
}

func main() {
	contents, err := os.ReadFile("tfl_graph.json")
	if err != nil {
		fmt.Println(err)
	}
	var network Network
	err = json.Unmarshal(contents, &network)
	if err != nil {
		fmt.Println(err)
	}
	for _, connection := range network.Connections {
		fmt.Println(connection)
	}
	breadthFirst(crawl, "Pimlico", network)
}

func breadthFirst(f func(string, Network) []string, current string, network Network) {
	seen := make(map[string]bool)
	worklist := []string{current}
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, network)...)
			}
		}
	}
}

func crawl(current string, network Network) []string {
	fmt.Println(current)
	var result []string
	for _, connection := range network.Connections {
		if connection.From == current {
			result = append(result, connection.To)
		}
	}
	return result
}
