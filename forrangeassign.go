package main

import (
	"fmt"
	"strconv"
)

type A struct {
	S []string
	C int64
}
type Z struct {
	B []A
}

func main() {
	a := make([]*A, 0)
	a = append(a, &A{
		C: 10,
	})
	a = append(a, &A{
		C: 11,
	})
	a = append(a, &A{
		C: 12,
	})

	// This can be assigned into a. Because a is array of ptrs, and x is a ptr.
	for _, x := range a {
		x.S = append(x.S, strconv.FormatInt(x.C, 10))
	}

	for _, x := range a {
		fmt.Println(x)
	}
}
