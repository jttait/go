package main

import (
	"fmt"
	"strings"
)

func join(sep string, elems ...string) string {
	result := ""
	for _, elem := range elems {
		result += elem + sep
	}
	return result[:len(result)-len(sep)]

}

func main() {
	fmt.Println(join(" ", "hello", "world"))
	fmt.Println(strings.Join([]string{"hello", "world"}, " "))
	fmt.Println(join("-", "hello", "world"))
	fmt.Println(strings.Join([]string{"hello", "world"}, "-"))
}
