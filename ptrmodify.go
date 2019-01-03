package main

import "fmt"

type a struct {
	a int
	b int
}

func sss(s *a) {
	s.a = 10
	s.b = 100
}
func main() {
	c := &a{}
	sss(c)
	fmt.Println(c)
}
