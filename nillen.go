package main

import (
	"fmt"
)

func main() {
	var a *int
	var b *int
	a = nil
	b = nil
	if a == b {
		fmt.Println("equals")
	} else {
		fmt.Println("not equal")
	}
	// This does not work.
	//fmt.Println(len(nil))
	// However this typed len works.
	var test_map map[int64]*int64
	test_map = nil
	fmt.Println(len(test_map))
}
