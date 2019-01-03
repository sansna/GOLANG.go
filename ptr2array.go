package main

import "fmt"

func sss(a *[]int64) {
	*a = make([]int64, 10)
	(*a)[1] = 1
	(*a)[2] = 10
	return
}

func main() {
	var disps []int64
	sss(&disps)
	fmt.Println(disps)
}
