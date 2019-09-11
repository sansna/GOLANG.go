package main

import (
	"fmt"
)

func main() {
	//fmt.Println(len(nil))
	var a *int
	var b *int
	a = nil
	b = nil
	if a == b {
		fmt.Println("equals")
	} else {
		fmt.Println("not equal")
	}
}
