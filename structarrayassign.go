package main

import "fmt"

type A struct {
	a int
	b int
}

type B struct {
	z []*A
}

func main() {
	a := make([]*A, 0)
	a = append(a, &A{
		a: 1,
		b: 2,
	})
	a = append(a, &A{
		a: 3,
		b: 7,
	})
	fmt.Println(a[0].b)
	b := &B{
		z: a,
	}
	fmt.Println(b.z[1].b)
}
