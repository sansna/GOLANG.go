package main

import (
	"fmt"
)

func main() {
	v1 := 1
	v2 := 1
	// You cannot do this
	if v1 == v2 == 0 {
		fmt.Println("ok")
	}
}
