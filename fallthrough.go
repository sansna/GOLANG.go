package main

import (
	"fmt"
)

func main() {
	a := 2
	b := make(map[string]int64)
	switch a {
	case 3:
		b["3"] = 100
		fallthrough
	case 2:
		b["2"] = 100
		fallthrough
	case 1:
		b["1"] = 100
		fallthrough
	case 100:
		b["100"] = 100
		fallthrough
	default:
	}
	fmt.Println(b)
}
