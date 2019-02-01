package main

import "fmt"

func main() {
	a := []int{1,2,4}
	b := []int{1,2,4}
	fmt.Println(append(a, b...))
	fmt.Println(append(a[:2],a[3:]...))
	return
}
