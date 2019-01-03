// defer acts FILO
package main

import "fmt"

func test() (i int) {
	// after exit, sets return value 'i' to i++
	defer func() {i++}()
	return 1
}

func main() {
	defer fmt.Println("life")
	defer fmt.Println("is")
	fmt.Println("1")
	defer fmt.Println("hh")
	fmt.Println("i++:", test())
	sss()
	return
}

func sss() {
	a := 10
	var r interface{}
	r = a
	defer func() {
		fmt.Println("lif",r.(int))
	}()
	b := 20
	r = b
	return
}
