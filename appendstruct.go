package main

import (
	"fmt"
)
type A struct {
	AA int64
}
type B struct {
	BB []*A
}

func main() {
	b := &B{}
	b.BB = append(b.BB, &A{
		AA: 10,
	})
	fmt.Println(b, b.BB[0].AA)

	c := make([]*B, 10)
	// range cannot assign list of structs, but loop can.
	for _, i := range c {
		i = &B{}
		i.BB = make([]*A, 0)
	}
	fmt.Println("range")
	for _, i := range c {
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		c[i] = &B{}
	}
	fmt.Println("loop")
	for _, i := range c {
		fmt.Println(i)
	}
}
