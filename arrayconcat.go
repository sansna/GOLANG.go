package main

import "fmt"

func main() {
	a := []int64{1, 23, 4, 5}
	b := []int64{5, 0, 8}
	c := make([]int64, 0)

	// nil ptr can be appended
	d := []int64{}
	d = nil

	fmt.Printf("c's address: %p\n", &c)
	c = append(c, a...)
	c = append(c, b...)
	c = append(c, d...)
	fmt.Println(c)
	// address do not change after copy
	fmt.Printf("c's address:%p\n", &c)
}
