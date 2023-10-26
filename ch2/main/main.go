package main

import (
	"fmt"
	"github.com/jttait/gopl.io/ch2/popcount"
)

func main() {
	val := uint64(5)
	fmt.Println(popcount.PopCount(val))
	fmt.Println(popcount.PopCount2(val))
	fmt.Println(popcount.PopCount3(val))
	fmt.Println(popcount.PopCount4(val))
}
