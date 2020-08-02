package main

import (
	"fmt"
)

func main() {
	var a map[int64][]int64
	fmt.Println(a)
	a = make(map[int64][]int64)
	fmt.Println(a)
	//
	a[0] = append(a[0], 1)
	fmt.Println(a)
}
