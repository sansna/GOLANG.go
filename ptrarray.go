package main

import "fmt"
//import "reflect"

type Elem struct {
	a1 int
	a2 int
}

func main() {
	// a is a thing that stands for slice of ptrs
	a := []*Elem{}
	// b is slice of ptrs that is allocated with space, but each ptr is nil
	b := make([]*Elem, 100)
	// a now stands for the slice of ptrs of b
	a = b
	// c is a ptr of elem
	c := &Elem{
		a1: 1,
	}
	// now b[10] points to the same elem as c
	b[10] = c
	// so now a[10] is b[10]
	fmt.Println(a[10].a1)
}
