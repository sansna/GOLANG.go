package main

import (
	"fmt"
)

func main() {
	a := make([]int, 1000)
	fmt.Println(len(a), cap(a))
	for i := 0; i < 1000; i++ {
		a[i] = i
	}

	// empty a slice quickly
	a = a[:0]
	fmt.Println(len(a), cap(a))

	// recover to original slice
	//a=a[0:1000]
	//fmt.Println(a[300])

	// now a contains only 1 item
	a = append(a, 321)
	fmt.Println(len(a), cap(a))
	fmt.Println("start")
	for _, i := range a {
		fmt.Println(i)
	}
	// only first item changed value.
	a = a[0:1000]
	fmt.Println("recovered")
	for _, i := range a {
		fmt.Println(i)
	}
}
