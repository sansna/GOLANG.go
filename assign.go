package main

import (
	"fmt"
)

type Member struct {
	a *int64
	b *int64
}

func main() {
	aa := int64(1)
	bb := int64(1)
	cc := int64(2)
	a := Member{
		a: &aa,
		b: &bb,
	}
	// ptr assign
	//a := &Member{
	//	a: &aa,
	//	b: &bb,
	//}
	b := a
	a.a = &cc
	fmt.Println(*b.a, *b.b)
	fmt.Println(*a.a, *a.b)
}
