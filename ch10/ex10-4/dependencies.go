package main

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
	"regexp"
)

func main() {
	output, err := exec.Command("go", "list", "-f", "{{join .Deps \" \"}}", os.Args[1]).Output()
	if err != nil {
		fmt.Println(err)
	}
	s := fmt.Sprintf("%s", output)
	re := regexp.MustCompile(`\r?\n\s\s+`)
	s = re.ReplaceAllString(s, " ")
	items := strings.Split(s, " ")

	result := make(map[string]struct{})

	for _, item := range items {
		output, err := exec.Command("go", "list", "-f", "{{join .Deps \" \"}}", item).Output()
		if err != nil {
			fmt.Println(err)
		}
		s := fmt.Sprintf("%s", output)
		re := regexp.MustCompile(`\r?\n`)
		s = re.ReplaceAllString(s, " ")
		re = regexp.MustCompile(`\s\s+`)
		s = re.ReplaceAllString(s, " ")
		deps := strings.Split(s, " ")

		for _, dep := range deps {
			if len(dep) > 0 {
				result[dep] = struct{}{}
			}
		}
	}

	for i := range result {
		fmt.Println(i)
	}
}
