package main

import "fmt"

type A struct {
	a int
	b int
}

func main() {
	s := &A{
		a: 1,
		b: 2,
	}
	ss := *s
	ss.a = 3

	// now sss is value copy of s, of same type.
	sss := &ss
	fmt.Println(s)
	fmt.Println(ss)
	fmt.Println(sss)
}
