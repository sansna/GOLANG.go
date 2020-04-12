package main

import "fmt"

func main() {
	var a bool
	a = true
	a = a || false
	if a {
		fmt.Println("ok")
	}
	//var a int
	//a = 1
	//if a {
	//	fmt.Println("ok")
	//}
	fmt.Println(true && true)
	fmt.Println(true && false)
}
