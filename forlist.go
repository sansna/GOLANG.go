package main

import (
	"fmt"
)

func main() {
	a := []bool{true, false}
	for idx, i := range a {
		fmt.Println(idx, i)
	}
}
