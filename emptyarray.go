package main

import (
	"fmt"
)

type A struct {
	Aa int64
}

func main() {
	empty := make([]A, 0)
	fmt.Println(empty)
}
