package main

import "fmt"

func main() {
	a := []int{1,2,4}
	b := []int{1,2,4}
	fmt.Println(append(a, b...))
	return
}
