package main

import (
	"fmt"
)

type A struct {
	B int
}

func main() {
	var c []*A
	c = nil
	fmt.Println(c)
	// You can append to a nil ptr array
	c = append(c, &A{
		B: 10,
	})
	fmt.Println(*c[0])
}
