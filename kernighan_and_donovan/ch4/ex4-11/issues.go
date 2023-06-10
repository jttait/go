package main

import (
	"fmt"
	"flag"
	
	"gopl.io/ch4/github"
)

func main() {
	var operationFlag = flag.String("operation", "get", "create, get, list, close, or update")
	var issueFlag = flag.String("issue", "not set", "issue number")
	var titleFlag = flag.String("title", "not set", "issue title")
	flag.Parse()
	
	if *operationFlag == "create" {
		if *titleFlag == "not set" {
			fmt.Printf("error: must provide title")
		} else {
			result, _ := github.CreateIssue(*titleFlag)
			fmt.Printf("issue created: %d\n", result.Number)
		}
	} else if *operationFlag == "close" {
		if *issueFlag == "not set" {
			fmt.Printf("error: must provide issue")
		} else {
			err := github.CloseIssue(*issueFlag)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("issue deleted: %s\n", *issueFlag)
		}
	} else if *operationFlag == "get" {
		if *issueFlag == "not set" {
			fmt.Printf("error: must provide issue")
		} else {
			err := github.GetIssue(*issueFlag)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if *operationFlag == "update" {
		if *issueFlag == "not set" {
			fmt.Printf("error: must provide issue")
		} else if *titleFlag == "not set" {
			fmt.Printf("error: must provide title")
		} else {
			err := github.UpdateIssue(*issueFlag, github.UpdateIssueRequest{Title: *titleFlag})
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("Error invalid arguments")
	}
}
