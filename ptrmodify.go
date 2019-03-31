package main

import "fmt"

type a struct {
	a int
	b int
}

func ptrsss(s []*int) {
	j := 2
	s[0] = &j
}
func sss(s *a) {
	s.a = 10
	s.b = 100
}
func main() {
	c := &a{}
	sss(c)
	fmt.Println(c)

	d := make([]*int,10)
	z := 1
	d[0]=&z
	ptrsss(d)
	fmt.Println(*d[0])
}
