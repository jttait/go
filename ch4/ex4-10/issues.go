package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	month, _ := time.ParseDuration("720h")
	year, _ := time.ParseDuration("8760h")

	var lessThanMonth []*github.Issue
	var lessThanYear []*github.Issue
	var moreThanYear []*github.Issue

	for _, item := range result.Items {
		age := time.Since(item.CreatedAt)
		
		if age < month {
			lessThanMonth = append(lessThanMonth, item)
		} else if age < year {
			lessThanYear = append(lessThanYear, item)
		} else {
			moreThanYear = append(moreThanYear, item)
		}
	}

	fmt.Println("Age less than a month:")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Age less than a year:")
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Age more than a year:")
	for _, item := range moreThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
