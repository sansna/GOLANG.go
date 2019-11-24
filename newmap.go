package main

import (
	"fmt"
)

var a map[int64]int64

func main() {
	a = make(map[int64]int64, 0)
	a[0] = 1
	fmt.Println(a)
}
