package main

import (
	"fmt"
)

func main() {
	a := make(map[int64][]string)
	a[1] = append(a[1], "asdf")
	fmt.Println(a)
}
