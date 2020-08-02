package main

import (
	"fmt"
)

type aa struct {
	a int64
}

func main() {
	a := make([]aa, 10)
	a[0] = aa{
		a: 100,
	}
	fmt.Println(a)

	for i, _ := range a {
		a[i].a = 101
	}
	fmt.Println(a)
}
