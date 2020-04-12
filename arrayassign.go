package main

import (
	"fmt"
)

func main() {
	a := []int64{1,2,34}
	b := []int64{3,5}
	fmt.Println(a,b)
	a = b
	fmt.Println(a,b)
}
