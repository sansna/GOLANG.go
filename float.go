package main

import (
	"fmt"
)

func main() {
	a := make(map[int64]float64)
	if a[1] > -1 {
		fmt.Println("ok")
	} else {
		fmt.Println("not")
	}
}
