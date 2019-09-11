package main

import "fmt"

func main() {
	x := 10
	b := 5
	switch x {
	case 1:
		fmt.Println(x)
	case 2:
		fmt.Println(x + 1)
	}
	fmt.Println("ok")
	switch {
	case x == 1:
		fmt.Println(x)
	case x == 3:
		fmt.Println(x + 1)
	case x == 10 && b == 5:
		fmt.Println(x * x)
	}
	fmt.Println("ok")

	// multi value
	y := 1
	switch y {
	case 1, 2:
		fmt.Println("1,2")
	case 10, 11:
		fmt.Println("10")
	default:
		fmt.Println("not in values")
	}
}
