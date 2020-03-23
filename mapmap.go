package main

import (
	"fmt"
)

func main() {
	a := make(map[int64]map[int64]int64)
	if a[1] == nil {
		a[1] = make(map[int64]int64, 0)
	}
	a[1][2] = 3
	fmt.Printf("%v", a)
}
