package main

import "fmt"
//import "reflect"

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

	// changes a also changes value in b.z
	a[0].b = 11
	fmt.Println(b.z[0].b)

	fmt.Println(&a, &(b.z))
}
