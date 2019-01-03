package main

import "fmt"

func main() {
	var i [2]int64
	j := []int64{1, 2}
	fmt.Println(i)
	copy(i[:], j[:2])
	fmt.Println(i)
}
