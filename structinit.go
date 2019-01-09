package main

import "fmt"

type bbb struct {
	a int
	b int
}

type aaa struct {
	c int
	d int
	bbb
}

func main() {
	a := aaa{
		c: 5,
		d: 5,
	}
	// prints 0, a.e exists.
	//fmt.Println(a.e.a)
	fmt.Println(a)

	b := bbb{
		a: 1,
		b: 2,
	}
	d := aaa{
		bbb: b,
		c:   3,
		d:   4,
	}
	fmt.Println(d)
}
