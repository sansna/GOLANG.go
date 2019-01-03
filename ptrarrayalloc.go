package main

//import "math"
import "fmt"
type aaa struct {
	a int
	b int
}

func main() {
	var a float64
	a = 0.0
	fmt.Println(a)
	p := []*aaa{}
	p = make([]*aaa, 1, 10)
	p[0] = &aaa{}
	p[0].a = 10
	fmt.Println(p[0].a)
	d := p
	fmt.Println(d[0].a)
}
