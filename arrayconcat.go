package main

import "fmt"

func main() {
	a := []int64{1, 23, 4, 5}
	b := []int64{5, 0, 8}
	c := make([]int64, 0)
	c = append(c, a...)
	c = append(c, b...)
	fmt.Println(c)
}
