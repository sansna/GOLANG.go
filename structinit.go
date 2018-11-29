package main

import "fmt"

type bbb struct {
	a int
	b int
}

type aaa struct {
	c int
	d int
	e bbb
}

func main() {
	a := aaa{
		c: 5,
		d: 5,
	}
	// prints 0, a.e exists.
	fmt.Println(a.e.a)
}
