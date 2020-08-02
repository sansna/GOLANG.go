package main

import (
	"fmt"
)

func insert(p *[]int64, pid int64) {
	*p = append(*p, pid)
	return
}

func main() {
	slice := []int64{1,2}
	insert(&slice, 100)
	fmt.Println(slice)
	return
}
