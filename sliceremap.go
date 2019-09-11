package main

import (
	"fmt"
)

// it works!
func main() {
	a := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := len(a)
	a = a[:4]
	a = append(a, a[0 : length-1][7:]...)
	fmt.Println(a)
	fmt.Println(len(a), cap(a), a[0:length])
}
