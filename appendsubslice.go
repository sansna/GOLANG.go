package main

import "fmt"

func main() {
	a := []int64{1, 2, 3}
	var b int64
	b = 6
	var c []int64
	c = append(c, b)
	c = append(c, a[1:]...)
	fmt.Println(c)
}
