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
	fmt.Println(86301 / 300 * 300)

	c := float64(1)
	// true
	fmt.Println(c, c == 1)
	var d interface{}
	d = 1
	fmt.Println(d.(int), d == 1)

	z := float64(13.1)
	fmt.Printf("%f", z)
}
