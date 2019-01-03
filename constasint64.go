package main

import "fmt"

const (
	AAAA = 1
)

// non-declared as false
var m = map[int64]bool{
//	1: true,
	2: true,
}

func main() {
	a := int32(1)
	if AAAA == a {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
	if m[AAAA] {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
