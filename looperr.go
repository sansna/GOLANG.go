package main

import (
	"fmt"
)

func main() {
	var out [][]int
	for _, i := range [][2]int{{1, 2}, {2, 3}, {3, 4}} {
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
}
