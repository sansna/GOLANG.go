package main

import "fmt"

func main() {
	var a *[]*int64
	c := make([]*int64, 0)
	a = &c
	b := int64(10)
	*a = append(*a, &b)
	*a = append(*a, &b)
	*a = append(*a, &b)
	*a = append(*a, &b)
	fmt.Println(*(*a)[0])
	fmt.Println(*(*a)[1])
	fmt.Println(*(*a)[2])
	fmt.Println(*(*a)[3])
	// nil
	//fmt.Println(*(*a)[4])

	var z *[]*int64
	z = a
	fmt.Println("z", *(*z)[3])
}
