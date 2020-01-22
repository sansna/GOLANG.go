package main

import (
	"fmt"
)

func main() {
	var b []int64
	a := []int64{1, 2, 3, 4}
	// slice is passed by ptr
	b = a
	b[3] = 8
	fmt.Println(a)
	fmt.Println(b)
}
