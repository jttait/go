package main

import (
	"fmt"
	"github.com/jttait/gopl.io/ch9/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(uint64(5)))
	fmt.Println(popcount.PopCount(uint64(220)))
}
