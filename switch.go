package main

import "fmt"

func main() {
	x := 10
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
	case x == 4:
		fmt.Println(x * x)
	}
	fmt.Println("ok")
}
