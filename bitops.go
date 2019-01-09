package main

import "fmt"
import "math"

func main() {
	fmt.Printf("%064b\n", uint(math.Pow(2, 40)-1))
	fmt.Printf("%b\n", ^uint(math.Pow(2, 40)-1))
	var a int64
	a = 10
	if a<<40>>40 != a {
		fmt.Printf("%d\n", a)
	} else {

		fmt.Printf("err%d\n", a)
	}
	// float
	fmt.Println(math.Inf(1))
	fmt.Println(math.Inf(-1))

	var b int64
	b = 1
	b = ^b
	fmt.Printf("%b\n", ^uint64(b))
	b = b << 40
	fmt.Printf("%b\n", uint64(b))
	b = b >> 40
	b = ^b
	fmt.Printf("%d\n", b)

	// you cannot >>/<< of type float64
	var c float64
	var i int
	for i = 0; i < 66; i++ {
		c = math.Pow(2, float64(i))
		if c != float64(int64(c)) {
			break
		}
	}
	fmt.Println(i)
	g := 10
	g ^= 1
	fmt.Println("g", g)
	g ^= 1
	fmt.Println("g", g)
	g = 3
	if 1 & g == 1 {
		fmt.Println("ok")
	}
}
