package main

import (
	"fmt"
)

type A struct {
	AA []*int
}

func ret(a *A) *A {
	m := 10
	n := 20
	b := make([]*int, 0)
	b = append(b, &m)
	b = append(b, &n)
	a.AA = b
	return a
}

func main() {
	m := 30
	n := 40
	z := make([]*int, 0)
	z = append(z, &m)
	z = append(z, &n)

	c := A{
		AA: z,
	}
	ret(&c)
	fmt.Printf("%v, %v", *c.AA[0], *c.AA[1])
}
